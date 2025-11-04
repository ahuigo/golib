package main

// https://blog.apnic.net/2021/05/12/programmatically-analyse-packet-captures-with-gopacket/

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func TestParsePcapAndStatis(t *testing.T) {
	// --- 配置 ---
	// 将这里的路径替换为你的 pcap 文件实际路径
	// 如果在 Android 的 /sdcard/ 目录下，你需要先将它 adb pull 到电脑上
	// 例如：adb pull /sdcard/netCapture.pcap .
	// pcapFile := "./tmp/pcap/busy.pcap"
	pcapFile := "./tmp/a.pcap"

	// 检查文件是否存在
	if _, err := os.Stat(pcapFile); os.IsNotExist(err) {
		log.Fatalf("Pcap 文件不存在: %s\n请确保文件和程序在同一目录，或提供完整路径。", pcapFile)
	}

	// --- 打开 Pcap 文件 ---
	handle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// --- 初始化统计计数器 ---
	// 使用 map 来存储不同 flag 的计数
	flagCounts := make(map[string]int)

	// --- 循环读取数据包 ---
	// 使用 gopacket.NewPacketSource 创建一个数据包源
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetCount := 0
	tcpPacketCount := 0

	fmt.Println("开始分析 Pcap 文件...")

	// 遍历数据包源中的所有数据包
	for packet := range packetSource.Packets() {
		packetCount++

		// 检查数据包是否包含 TCP 层
		// packet.Layer() 会返回指定类型的第一个层
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			// 如果没有 TCP 层，则跳过这个包
			continue
		}

		tcpPacketCount++

		// 将获取的层进行类型断言，转换为 *layers.TCP 类型，以便访问其字段
		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			// 理论上这不应该发生，但作为健壮性检查
			log.Println("无法将层转换为 *layers.TCP")
			continue
		}

		// --- 检查并统计 TCP 标志 ---
		// 一个数据包可能同时拥有多个标志，例如 SYN-ACK 包会同时增加 SYN 和 ACK 的计数
		if tcp.SYN {
			flagCounts["SYN"]++
		}
		if tcp.ACK {
			flagCounts["ACK"]++
		}
		if tcp.RST {
			flagCounts["RST"]++
		}
		if tcp.FIN {
			flagCounts["FIN"]++
		}
		if tcp.PSH {
			flagCounts["PSH"]++
		}
		if tcp.URG {
			flagCounts["URG"]++
		}
		if tcp.ECE {
			flagCounts["ECE"]++
		}
		if tcp.CWR {
			flagCounts["CWR"]++
		}
	}

	// --- 打印统计结果 ---
	fmt.Println("\n--- 分析结果 ---")
	fmt.Printf("总共扫描数据包数量: %d\n", packetCount)
	fmt.Printf("其中 TCP 数据包数量: %d\n", tcpPacketCount)
	fmt.Println("\nTCP 标志 (Flags) 统计:")
	fmt.Println("----------------------")
	// 为了美观，可以按特定顺序列出
	flagsToShow := []string{"SYN", "ACK", "RST", "FIN", "PSH", "URG", "ECE", "CWR"}
	for _, flag := range flagsToShow {
		if count, found := flagCounts[flag]; found {
			fmt.Printf("%-5s : %d\n", flag, count)
		}
	}
	fmt.Println("----------------------")
	fmt.Println("\n注意: 一个数据包可能包含多个标志 (例如 SYN-ACK)，因此会为每个标志单独计数。")
}

func TestParsePcap(t *testing.T) {
	handle, err := pcap.OpenOffline("/tmp/lo.pcap")
	if err != nil {
		t.Fatal(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(
		handle, handle.LinkType()).Packets()
	for pkt := range packets {
		println(pkt)
		// Your analysis here!
	}
}
