package online

type Volume struct {
	Id           string
	Name         string
	Organization string
	Size         int
}

type GetVolumesResp struct {
	Volumes []Volume
}
