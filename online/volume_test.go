package online

import (
	"encoding/json"
	"testing"
)

func TestUnmarshallVolume(t *testing.T) {
	b := []byte(`{
      "export_uri": null,
      "id": "f929fe39-63f8-4be8-a80e-1e9c8ae22a76",
      "name": "volume-0-1",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "server": null,
      "size": 10000000000,
      "volume_type": "l_hdd"
    }`)

	// d as domain
	var d Volume
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := "volume-0-1"
	if d.Name != exp {
		t.Fatalf("Expected %#v, got %#v", exp, d.Name)
	}

	exp = "000a115d-2852-4b0a-9ce8-47f1134ba95a"
	if d.Organization != exp {
		t.Fatalf("Expected %#v, got %#v", exp, d.Organization)
	}

	expSize := int64(10000000000)
	if d.Size != expSize {
		t.Fatalf("Expected %#v, got %#v", expSize, d.Size)
	}
}

func TestGetVolumesResp(t *testing.T) {
	b := []byte(`{
  "volumes": [
    {
      "export_uri": null,
      "id": "f929fe39-63f8-4be8-a80e-1e9c8ae22a76",
      "name": "volume-0-1",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "server": null,
      "size": 10000000000,
      "volume_type": "l_hdd"
    },
    {
      "export_uri": null,
      "id": "0facb6b5-b117-441a-81c1-f28b1d723779",
      "name": "volume-0-2",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "server": null,
      "size": 20000000000,
      "volume_type": "l_ssd"
    }
  ]
}`)

	// d as domain
	var d GetVolumesResp
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := 2
	if len(d.Volumes) != exp {
		t.Fatalf("Expected %#v, got %#v", exp, len(d.Volumes))
	}
}
