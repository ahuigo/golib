package packet

/*
Keep in mind:

This example assumes you're just crafting and sending a packet for educational purposes;
The IP header checksum is calculated over the IP header only.
The TCP checksum includes a pseudo-header, the TCP header, and any data (which there isn't any in an ACK).
You may need to adjust the sequence number and acknowledgment number according to what the server expects to receive, which usually involves capturing network traffic with a tool like Wireshark.
*/

import (
	"encoding/binary"
	"log"
	"net"
	"syscall"
	"testing"
)

// converts a short (uint16) from host-to-network byte order.
func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}
func TestSendSyn(t *testing.T) {
	// 测试不成功
	log.SetFlags(log.Lshortfile)
	const (
		sourceIP        = "127.0.0.1"
		destinationIP   = "127.0.0.1"
		destinationPort = 1025  // Port where the TCP server is running
		sourcePort      = 43591 // Arbitrary source port
		synFlag         = 0x02  // SYN flag for TCP
	)

	// Parse the destination IP address
	ip := net.ParseIP(destinationIP).To4()
	if ip == nil {
		log.Fatalf("Invalid IP address: %v\n", destinationIP)
	}

	// Create a raw socket
	sockFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatalf("Could not create socket: %v\n", err)
	}
	defer syscall.Close(sockFd)

	// Enable IP_HDRINCL to tell the kernel that we are providing the IP header
	if err := syscall.SetsockoptInt(sockFd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1); err != nil {
		log.Fatalf("Could not set IP_HDRINCL: %v\n", err)
	}

	// Construct the IP header
	ipHeader := make([]byte, 20)
	ipHeader[0] = 0x45                               // IP version and header length
	ipHeader[1] = 0                                  // Type of service
	binary.BigEndian.PutUint16(ipHeader[2:4], 40)    // Total length (IP header + TCP header)
	binary.BigEndian.PutUint16(ipHeader[4:6], 54321) // Identification
	binary.BigEndian.PutUint16(ipHeader[6:8], 0)     // Flags and fragment offset
	ipHeader[8] = 64                                 // TTL
	ipHeader[9] = syscall.IPPROTO_TCP                // Protocol TCP
	// Skip checksum for now; calculated later on ipHeader[10:12]
	copy(ipHeader[12:16], net.ParseIP(sourceIP).To4()) // Source IP address
	copy(ipHeader[16:20], ip)                          // Destination IP address

	// Calculate the IP header checksum
	ipChecksum := checksum(ipHeader)
	binary.BigEndian.PutUint16(ipHeader[10:12], ipChecksum)

	// Construct the TCP header
	tcpHeader := make([]byte, 20)
	binary.BigEndian.PutUint16(tcpHeader[0:2], uint16(sourcePort))      // Source port
	binary.BigEndian.PutUint16(tcpHeader[2:4], uint16(destinationPort)) // Destination port
	binary.BigEndian.PutUint32(tcpHeader[4:8], 0)                       // Sequence number
	binary.BigEndian.PutUint32(tcpHeader[8:12], 0)                      // Acknowledgement number
	tcpHeader[12] = 80                                                  // Header length and reserved; 80 in decimal is 0x50 in hex, which means 5*4=20 bytes long header
	tcpHeader[13] = synFlag                                             // Flags with the SYN bit set
	binary.BigEndian.PutUint16(tcpHeader[14:16], 0xFFFF)                // Window size

	// Pseudo header for checksum calculation
	pseudoHeader := make([]byte, 12)
	copy(pseudoHeader[0:4], net.ParseIP(sourceIP).To4())                    // Source IP
	copy(pseudoHeader[4:8], ip)                                             // Destination IP
	pseudoHeader[8] = 0                                                     // Reserved
	pseudoHeader[9] = syscall.IPPROTO_TCP                                   // Protocol
	binary.BigEndian.PutUint16(pseudoHeader[10:12], uint16(len(tcpHeader))) // TCP header length

	// Concat pseudo-header and TCP header for checksum calculation
	chksumData := append(pseudoHeader, tcpHeader...)
	tcpChecksum := checksum(chksumData)
	binary.BigEndian.PutUint16(tcpHeader[16:18], tcpChecksum)

	// Concatenate IP header and TCP header to form the complete packet
	packet := append(ipHeader, tcpHeader...)

	// Send the packet
	dstport := htons(uint16(destinationPort))
	dstSockaddr := syscall.SockaddrInet4{Port: int(dstport)}
	// dstSockaddr = syscall.SockaddrInet4{Port: destinationPort}
	copy(dstSockaddr.Addr[:], ip)
	if err := syscall.Sendto(sockFd, packet, 0, &dstSockaddr); err != nil {
		log.Fatalf("Sendto failed: %v\n", err)
	}
	log.Println("SYN packet sent")
}

// checksum calculates the Internet checksum specified in RFC 1071.
func checksum(data []byte) uint16 {
	var sum uint32
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(data[i])<<8 + uint32(data[i+1])
	}
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}
	for sum > 0xffff {
		sum = (sum & 0xffff) + (sum >> 16)
	}
	return ^uint16(sum)
}
