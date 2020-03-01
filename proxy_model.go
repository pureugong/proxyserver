package proxylist

// Proxy is
type Proxy struct {
	ip        string
	port      string
	anonymity string
	speed     int
}

func newProxy(ip, port, anonymity string, speed int) Proxy {
	return Proxy{
		ip:        ip,
		port:      port,
		anonymity: anonymity,
		speed:     speed,
	}
}

// GetIP is
func (p Proxy) GetIP() string {
	return p.ip
}

// GetPort is
func (p Proxy) GetPort() string {
	return p.port
}

// GetSpeed is
func (p Proxy) GetSpeed() int {
	return p.speed
}

// GetAnonymity is
func (p Proxy) GetAnonymity() string {
	return p.anonymity
}
