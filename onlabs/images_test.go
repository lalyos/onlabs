package onlabs

import (
	"encoding/json"
	"testing"
)

func TestUnmarshallImage(t *testing.T) {
	b := []byte(`{
      "default_bootscript": null,
      "creation_date": "2014-10-07T15:29:49.970088+00:00",
      "name": "Gentoo stable (2014.10.07)",
      "modification_date": "2014-10-24T15:07:19.423013+00:00",
      "organization": "a283af0b-d13e-42e1-a43f-855ffbf281ab",
      "extra_volumes": "[]",
      "arch": "arm",
      "id": "a8a3c7e8-737a-4198-ae67-775702a75c6a",
      "root_volume": {
        "id": "0042cc7e-b739-48b1-b3b9-748d8f150ea3",
        "name": "Gentoo stable on SSD 2014.10.07 snapshot"
      },
      "public": true
    }`)

	var image Image
	err := json.Unmarshal(b, &image)
	if err != nil {
		t.Fatal(err)
	}

	if image.Name != "Gentoo stable (2014.10.07)" {
		t.Fatalf("Expected image.Name=Gentoo stable (2014.10.07) got %s", image.Name)
	}

	if image.Organization != "a283af0b-d13e-42e1-a43f-855ffbf281ab" {
		t.Fatalf("Expected image.Organization=a283af0b-d13e-42e1-a43f-855ffbf281ab got %s", image.Organization)
	}

	if image.Arch != "arm" {
		t.Fatalf("Expected image.Arch=arm got %s", image.Arch)
	}

	if image.Public != true {
		t.Fatalf("Expected image.Public=true got %s", image.Public)
	}

}

func TestGetImagesResp(t *testing.T) {
	b := []byte(`{
  "images": [
    {
      "default_bootscript": null,
      "creation_date": "2014-10-24T10:11:24.393374+00:00",
      "name": "Debian Wheezy (7.7)",
      "modification_date": "2014-10-24T15:03:55.110363+00:00",
      "organization": "a283af0b-d13e-42e1-a43f-855ffbf281ab",
      "extra_volumes": "[]",
      "arch": "arm",
      "id": "05cd1a2d-9f99-4252-8791-c7b12d5137d8",
      "root_volume": {
        "id": "edaf2c29-7542-49b3-ada7-1ec0bb5252cc",
        "name": "NEW DEBIAN 5151 snapshot"
      },
      "public": true
    },
    {
      "default_bootscript": {
        "kernel": {
          "dtb": "dtb/pimouss-computing.dtb.3.17",
          "path": "kernel/pimouss-uImage-3.17-90-perf",
          "id": "2a04936b-dc53-4221-a0d7-a482af8b29d3",
          "title": "Pimouss 3.17-90-perf"
        },
        "initrd": {
          "path": "initrd/pimouss-uInitrd",
          "id": "fe70e4dc-fb87-47e8-bf61-4c75c6f5a61e",
          "title": "pimouss-uInitrd"
        },
        "bootcmdargs": {
          "id": "bf308e86-dc34-43d5-bb37-cdd3b5596ca2",
          "value": "ip=dhcp boot=local root=/dev/nbd0"
        },
        "id": "025b67f9-10a5-4478-94a7-1c13eaa16356"
      },
      "creation_date": "2014-10-23T21:47:25.553354+00:00",
      "name": "Docker (1.2) on Ubuntu Utopic (14.10)",
      "modification_date": "2014-10-24T15:08:17.246210+00:00",
      "organization": "a283af0b-d13e-42e1-a43f-855ffbf281ab",
      "extra_volumes": "[]",
      "arch": "arm",
      "id": "e6cdbc3f-5e9a-48f8-9335-494bac14befc",
      "root_volume": {
        "id": "67f36d57-d571-48dd-b08f-2a7242379aac",
        "name": "NEW DOCKER snapshot"
      },
      "public": true
    }
   ]
  }`)

	var resp GetImagesResp
	err := json.Unmarshal(b, &resp)
	if err != nil {
		t.Fatal(err)
	}

	s := len(resp.Images)
	if s != 2 {
		t.Fatalf("Expected 2: got %d", s)
	}

	expName := "Debian Wheezy (7.7)"
	i0 := resp.Images[0]
	if i0.Name != expName {
		t.Fatalf("Expected %q, got: %q", expName, i0.Name)
	}

	expName = "Docker (1.2) on Ubuntu Utopic (14.10)"
	i1 := resp.Images[1]
	if i1.Name != expName {
		t.Fatalf("Expected %q, got: %q", expName, i1.Name)
	}

}
