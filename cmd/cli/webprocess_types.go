package cli

type IWebSrvProcessServer interface {
	GetPorts() []string
	GetPubAddress() string
	GetAuthToken() string
	GetUptime() string
	GetProcesses() []WebSrvServer
}
type WebSrvProcessServer struct {
	Ports      []string       `json:"ports"`
	PubAddress string         `json:"pubAddress"`
	AuthToken  string         `json:"authToken"`
	Uptime     string         `json:"uptime"`
	Processes  []WebSrvServer `json:"processes"`
}

func NewWebSrvProcessServer(ports []string, pubAddress, authToken, uptime string, processes []WebSrvServer) IWebSrvProcessServer {
	return WebSrvProcessServer{
		Ports:      ports,
		PubAddress: pubAddress,
		AuthToken:  authToken,
		Uptime:     uptime,
		Processes:  processes,
	}
}

func (w WebSrvProcessServer) GetPorts() []string           { return w.Ports }
func (w WebSrvProcessServer) GetPubAddress() string        { return w.PubAddress }
func (w WebSrvProcessServer) GetAuthToken() string         { return w.AuthToken }
func (w WebSrvProcessServer) GetUptime() string            { return w.Uptime }
func (w WebSrvProcessServer) GetProcesses() []WebSrvServer { return w.Processes }
