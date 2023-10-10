package packet

import (
	"net"
	"runtime"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// sudo go test -timeout 1800s -run ^TestSendRSTPacket$ t/packet -v
func TestSendRSTPacket(t *testing.T) {
	// 打开网络设备
	handle, err := pcap.OpenLive(getInterface(), 1024, false, pcap.BlockForever)
	if err != nil {
		t.Fatal(err)
	}
	defer handle.Close()

	// 构造 IP 层
	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{127, 0, 0, 1},
	}

	// 构造 TCP 层
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(12345),
		DstPort: layers.TCPPort(80),
		RST:     true,
	}

	// 构造完整的数据包
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buffer, options, ipLayer, tcpLayer)
	if err != nil {
		t.Fatal(err)
	}

	// 发送数据包
	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		t.Fatal(err)
	}
}
func getInterface() string {
	if runtime.GOOS == "darwin" {
		return "lo0" //"en0" //"lo0"
	} else {
		return "eth0"
	}
}
