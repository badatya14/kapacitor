// +build windows

package dns

import "net"

type SessionUDP struct {
	raddr *net.UDPAddr
}

func (s *SessionUDP) RemoteAddr() net.Addr { return s.raddr }

// ReadFromSessionUDP acts just like net.UDPConn.ReadFrom(), but returns a session object instead of a
// net.UDPAddr.
func ReadFromSessionUDP(conn *net.UDPConn, b []byte) (int, *SessionUDP, error) {
	n, raddr, err := conn.ReadFrom(b)
	if err != nil {
		return n, nil, err
	}
	session := &SessionUDP{raddr.(*net.UDPAddr)}
	return n, session, err
}

// WriteToSessionUDP acts just like net.UDPConn.WritetTo(), but uses a *SessionUDP instead of a net.Addr.
func WriteToSessionUDP(conn *net.UDPConn, b []byte, session *SessionUDP) (int, error) {
	n, err := conn.WriteTo(b, session.raddr)
	return n, err
}

