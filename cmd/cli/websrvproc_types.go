package cli

type IWebSrvServerProcessData interface {
	GetWevSrvServerProcessData() ([]WebSrvServer, error)
	GetWebSrvServerData() (WebSrvProcessServer, error)
}
type WebSrvServerProcessData struct{}

func NewWebSrvServerProcessData() IWebSrvServerProcessData { return &WebSrvServerProcessData{} }

func (wss *WebSrvServerProcessData) GetWevSrvServerProcessData() ([]WebSrvServer, error) {
	return []WebSrvServer{}, nil
}
func (wss *WebSrvServerProcessData) GetWebSrvServerData() (WebSrvProcessServer, error) {
	return WebSrvProcessServer{}, nil
}
