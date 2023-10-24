package netstat

import (
	"errors"
	"ginapp/utils/os/shell"
	"strings"
)

/*
*
# netstat -antup 输出保存到stdout中，格式如下：
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 10.244.64.176:54764     192.168.145.102:4318    ESTABLISHED 1/mpush-server
tcp        0      0 10.244.64.176:45146     47.93.126.196:443       ESTABLISHED 1/mpush-server
tcp        0      0 10.244.64.176:43816     10.245.218.86:80        ESTABLISHED 1/mpush-server
*/
func GetAllTcpConnections() (conns []TcpConnection, err error) {
	// lsof -iTCP -sTCP:ESTABLISHED,LISTEN -n -P
	cmd := "lsof -iTCP -n -P"
	stdout, errmsg, errno := shell.ExecCommand("sh", "-c", cmd)
	if errno != 0 {
		err = errors.New(errmsg)
		return conns, err
	}
	lines := strings.Split(stdout, "\n")
	if len(lines) <= 2 {
		return conns, nil
	}
	conns = make([]TcpConnection, 0)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		if fields[0] != "tcp" {
			continue
		}
		localAddr := fields[3]
		remoteAddr := fields[4]
		state := fields[5]
		pid, programName, _ := strings.Cut(fields[6], "/")
		conns = append(conns, TcpConnection{
			LocalAddr:   localAddr,
			ForeignAddr: remoteAddr,
			State:       state,
			Pid:         pid,
			Program:     programName,
		})
	}
	return conns, err
}
