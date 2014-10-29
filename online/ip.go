package online

type IP struct {
	Id           string
	Address      string
	Organization string
	Server       struct {
		Id   string
		Name string
	}
}

type GetIPsResp struct {
	IPs []IP
}
