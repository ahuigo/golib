package main

import (
	"fmt"

    "flag"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	// The same default as tcpdump.
	defaultSnapLen = 262144
)

func main() {
	//handle, err := pcap.OpenLive("en0", defaultSnapLen, true,
	//handle, err := pcap.OpenLive("lo0", defaultSnapLen, true,
	handle, err := pcap.OpenLive("any", defaultSnapLen, true,
		pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	// capture 9999 only
    // nc -l 9999
    // curl localhost:9999 // for lo0
    portPtr := flag.String("p", "9999", "read port")
    flag.Parse()
	port := *portPtr
	println("capture port ", port)
	if err := handle.SetBPFFilter("port " + port); err != nil {
		panic(err)
	}
	packets := gopacket.
        NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		// Your analysis here!
		// fmt.Println("pkt:", pkt)
		fmt.Println("dump:-----------------------------\n", pkt.Dump())

	}
}
