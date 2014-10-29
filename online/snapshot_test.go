package online

import (
	"encoding/json"
	"testing"
)

func TestUnmarshallSnapshot(t *testing.T) {
	b := []byte(`{
		"name": "myname"
		}`)

	// d as domain
	var d Snapshot
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := "myname"
	if d.Name != exp {
		t.Fatalf("Expected %#v, got %#v", exp, d.Name)
	}

}

func TestGetSnapshotsResp(t *testing.T) {
	b := []byte(`{
  "snapshots": [
    {
      "base_volume": {
        "id": "09a4184c-733b-43c8-99c3-f1dde30536fe",
        "name": "vol simple snapshot"
      },
      "creation_date": "2014-05-22T12:11:06.055998+00:00",
      "id": "6f418e5f-b42d-4423-a0b5-349c74c454a4",
      "name": "snapshot-0-1",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "size": 10000000000,
      "state": "snapshotting",
      "volume_type": "l_hdd"
    },
    {
      "base_volume": {
        "id": "09a4184c-733b-43c8-99c3-f1dde30536fe",
        "name": "vol simple snapshot"
      },
      "creation_date": "2014-05-22T12:13:09.877961+00:00",
      "id": "c6ff5501-eb35-44b8-aa01-8777211a830b",
      "name": "snapshot-0-2",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "size": 10000000000,
      "state": "snapshotting",
      "volume_type": "l_hdd"
    }
  ]
}`)

	// d as domain
	var d GetSnapshotsResp
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := 2
	if len(d.Snapshots) != exp {
		t.Fatalf("Expected %#v, got %#v", exp, len(d.Snapshots))
	}
}
