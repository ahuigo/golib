# refer: net/net-tcp-queue.md
srv:
	go run srv.go
client:
	go run client.go

syn-flood-tcpdump:
	sudo tcpdump -tn -i any port 8877
syn-flood-srv:
	gcc tcp-recv-server.c && ./a.out

syn-flood:
	pip install scapy
	hash brew && brew install libpcap
	#sysctl -w net.ipv4.tcp_syncookies=0
	sudo python send-syn.py



