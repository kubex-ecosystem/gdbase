package cli

type IWebSrvServer interface {
	GetServerProcessData() ([]IWebSrvServer, error)
	GetServerData() (IWebSrvProcessServer, error)
	GetPorts() []string
	GetPubAddress() string
	GetUptime() string
	GetProcesses() []IWebSrvServer
}
type WebSrvServer struct {
	Pid         int    `json:"pid"`
	Port        string `json:"port"`
	IpVersion   string `json:"ipVersion"`
	BindAddress string `json:"bindAddress"`
	SourcePath  string `json:"sourcePath"`
	Uptime      string `json:"uptime"`
	Processes   []IWebSrvServer
	Ports       []string
}

func NewWebSrvServer(pid int, port, ipVersion, bindAddress, sourcePath, uptime string, processes []IWebSrvServer, ports []string) IWebSrvServer {
	return &WebSrvServer{
		Pid:         pid,
		Port:        port,
		IpVersion:   ipVersion,
		BindAddress: bindAddress,
		SourcePath:  sourcePath,
		Uptime:      uptime,
		Processes:   processes,
		Ports:       ports,
	}
}

func (wss *WebSrvServer) GetServerProcessData() ([]IWebSrvServer, error) {
	return []IWebSrvServer{}, nil
}
func (wss *WebSrvServer) GetServerData() (IWebSrvProcessServer, error) {
	return WebSrvProcessServer{}, nil
}
func (wss *WebSrvServer) GetPorts() []string            { return wss.Ports }
func (wss *WebSrvServer) GetPubAddress() string         { return wss.BindAddress }
func (wss *WebSrvServer) GetUptime() string             { return wss.Uptime }
func (wss *WebSrvServer) GetProcesses() []IWebSrvServer { return wss.Processes }
