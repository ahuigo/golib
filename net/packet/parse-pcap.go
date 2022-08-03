package main

// https://blog.apnic.net/2021/05/12/programmatically-analyse-packet-captures-with-gopacket/


import (
	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func TestParsePcap() {
	handle, err := pcap.OpenOffline("/tmp/lo.pcap")
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(
		handle, handle.LinkType()).Packets()
	for pkt := range packets {
        println(pkt)
		// Your analysis here!
	}
}
