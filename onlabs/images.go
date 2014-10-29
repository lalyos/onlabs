package onlabs

type Image struct {
	Id           string
	Arch         string
	Name         string
	Organization string
	Public       bool
}

type GetImagesResp struct {
	Images []Image
}
