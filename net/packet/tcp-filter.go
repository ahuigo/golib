package main

import (
        "fmt"

        "github.com/google/gopacket"
        "github.com/google/gopacket/layers"
        "github.com/google/gopacket/pcap"
)

func main() {
        if handle, err := pcap.OpenLive("any", 1600, true, pcap.BlockForever); err != nil {
                panic(err)
        } else if err := handle.SetBPFFilter("tcp"); err != nil { // optional
                panic(err)
        } else {
                packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
                for packet := range packetSource.Packets() {
                        handlePacket(packet) // Do something with a packet here.
                }
        }
}

func handlePacket(packet gopacket.Packet) {
        if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
                ip, _ := ipLayer.(*layers.IPv4)

                if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
                        // Get actual TCP data from this layer
                        tcp, _ := tcpLayer.(*layers.TCP)
                        fmt.Printf("%s:%s -> %s:%s\n", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)
                }
        }
}
