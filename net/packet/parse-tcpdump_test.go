package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

/**
S：SYN（同步），用于建立连接。
F：FIN（结束），用于断开连接。
R：RST（重置），用于重置连接。
P：PSH（推送），提示对方尽快将数据交给应用层。
.：ACK（确认），表示确认号有效。
U：URG（紧急），表示紧急指针有效。
E：ECE（显式拥塞回执）。
W：CWR（拥塞窗口减少）。
flag "P." 代表同时设置了 PSH（P）和 ACK（.）标志，意思是“推送+确认”。

其它常见组合如：

S.：SYN+ACK
F.：FIN+ACK
R：仅 RST
.：仅 ACK
P.：PSH+ACK
*/

func TestParseTcpdump(t *testing.T) {
	// tcp seq 与ack 对不上
	logFile := "./tmp1/busy.log"
	// 统计各 TCP 标志的出现次数
	flagCounts := make(map[string]int)
	packetCount := 0
	tcpPacketCount := 0

	// 打开日志文件
	f, err := os.Open(logFile)
	if err != nil {
		t.Fatalf("无法打开日志文件: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	flagRe := regexp.MustCompile(`Flags \[([^\]]+)\]`)
	timeRe := regexp.MustCompile(`^(\d+):(\d+):(\d+\.\d+)`)
	seqRe := regexp.MustCompile(`seq (\d+):(\d+)`) // 提取 seq start:end
	ackRe := regexp.MustCompile(`ack (\d+)`)       // 提取 ack 值

	// 用于统计延迟：记录配对的 Out/In 包时间戳和序列号
	type req_resp struct {
		out_ts  float64 // Out 包的时间戳（秒）
		in_ts   float64 // In 包的时间戳（秒）
		out_seq uint32  // Out 包的 seq end 值
	}
	var pairs []req_resp
	// 使用 map 存储 Out 包和 In 包，key 是 seq end 值，value 是时间戳
	pendingOuts := make(map[uint32]float64)
	// 基准值
	var outSeqBase uint32

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Flags [") {
			tcpPacketCount++
			m := flagRe.FindStringSubmatch(line)
			if len(m) > 1 {
				flags := m[1]
				flagCounts[flags]++
			}
		}
		if len(line) > 0 && line[0] != '\t' && strings.Contains(line, ": Flags [") {
			packetCount++
		}
		// 统计 Out/In 包时间（只统计 [P.] 包）
		if len(line) > 0 && line[0] != '\t' && strings.Contains(line, "Flags [P.]") {
			tm := timeRe.FindStringSubmatch(line)
			seqMatch := seqRe.FindStringSubmatch(line)
			ackMatch := ackRe.FindStringSubmatch(line)

			if tm != nil && seqMatch != nil && ackMatch != nil {
				h, _ := strconv.Atoi(tm[1])
				m_, _ := strconv.Atoi(tm[2])
				s, _ := strconv.ParseFloat(tm[3], 64)
				ts := float64(h*3600+m_*60) + s

				seqStart, _ := strconv.ParseUint(seqMatch[1], 10, 32)
				seqEnd, _ := strconv.ParseUint(seqMatch[2], 10, 32)
				ack, _ := strconv.ParseUint(ackMatch[1], 10, 32)

				seqStartVal := uint32(seqStart)
				seqEndVal := uint32(seqEnd)
				ackValue := uint32(ack)

				if strings.Contains(line, " Out ") {
					// 处理第一个 Out 包
					if outSeqBase == 0 {
						outSeqBase = seqStartVal
						seqEndVal = seqEndVal - outSeqBase
						ackValue = 1
					} else {
						// 后续 Out 包已经是相对值，直接用 seqEnd - seqStart
					}

					// 存储 Out 包的相对 seq end
					pendingOuts[seqEndVal] = ts
				} else if strings.Contains(line, " In  ") {
					// In 包的 ack 对应 Out 包的相对 seq end
					if outTs, found := pendingOuts[ackValue]; found {
						pairs = append(pairs, req_resp{out_ts: outTs, in_ts: ts, out_seq: ackValue})
						delete(pendingOuts, ackValue)
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("读取日志文件出错: %v", err)
	}

	// 计算延迟
	var delays []float64
	for _, pair := range pairs {
		delay := pair.in_ts - pair.out_ts
		delays = append(delays, delay)
	}
	// 计算 p50, p90
	if len(delays) > 0 {
		sort.Float64s(delays)
		p50 := delays[len(delays)*50/100]
		p90 := delays[len(delays)*90/100]
		p99 := delays[len(delays)*99/100]
		t.Logf("心跳延迟 p50=%.6f 秒, p90=%.6f 秒, p99=%.6f 秒", p50, p90, p99)
	} else {
		t.Logf("未统计到心跳延迟")
	}

	t.Logf("总包数: %d, TCP包数: %d", packetCount, tcpPacketCount)
	for flag, cnt := range flagCounts {
		t.Logf("标志 %q 出现 %d 次", flag, cnt)
	}

	if len(flagCounts) == 0 {
		t.Errorf("未统计到任何TCP标志")
	}
}
