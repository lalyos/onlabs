package online

type Server struct {
	Id           string
	Name         string
	Organization string
	Image        struct {
		Id   string
		Name string
	}
	PublicIp PublicIp `json:"public_ip"`
	State    string
}

type PublicIp struct {
	Id      string
	Dynamic bool
	Address string
}

type GetServersResp struct {
	Servers []Server
}
