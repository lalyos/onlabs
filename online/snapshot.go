package online

type Snapshot struct {
	Id           string
	Name         string
	Organization string
	Size         int64
}

type GetSnapshotsResp struct {
	Snapshots []Snapshot
}
