package online

type Volume struct {
	Id           string
	Name         string
	Organization string
	Size         int64
}

type GetVolumesResp struct {
	Volumes []Volume
}
