package netstat

import "ginapp/utils/os/lsof"

func GetAllTcpConnections() (conns []TcpConnection, err error) {
	lsofTcps, err := lsof.GetLsofTcps()
	if err != nil {
		return conns, err
	}
	conns = make([]TcpConnection, len(lsofTcps))
	for _, tcp := range lsofTcps {
		conns = append(conns, TcpConnection{
			Proto:       tcp.Node,
			LocalAddr:   tcp.LocalAddr,
			ForeignAddr: tcp.ForeignAddr,
			State:       tcp.TcpState,
		})

	}
	return conns, err
}
