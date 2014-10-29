package online

type Snapshot struct {
	Id           string
	Name         string
	Organization string
	Size         int
}

type GetSnapshotsResp struct {
	Snapshots []Snapshot
}
