package online

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshallPublicIp(t *testing.T) {
	b := []byte(`{
		"id": "d3f55ed5-7a13-4611-92bf-c0bec332ef80",
		"dynamic": true,
		"address": "212.47.237.1"
	}`)

	var ip PublicIp
	err := json.Unmarshal(b, &ip)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("ip: %#v\n", ip)

}

func TestUnmarshallServer(t *testing.T) {
	b := []byte(`{
			"tags": [],
			"state_detail": "stopping",
			"image": {
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
			},
			"creation_date": "2014-10-28T07:00:42.739043+00:00",
			"public_ip": {
				"dynamic": true,
				"id": "1eacf9ed-55d5-4006-ad71-dc06dd6866f6",
				"address": "212.47.234.149"
			},
			"private_ip": "10.1.6.10",
			"id": "b43f78aa-c856-4d1d-9a47-73a503ea27a9",
			"modification_date": "2014-10-29T15:47:42.627496+00:00",
			"name": "my_server",
			"dynamic_public_ip": false,
			"state": "running",
			"bootscript": null,
			"volumes": {
				"0": {
					"size": 20000000000,
					"name": "NEW DOCKER snapshot",
					"modification_date": "2014-10-28T06:57:47.832880+00:00",
					"organization": "8d768494-0f72-4929-81ca-b8e157e61a49",
					"export_uri": "nbd://10.1.14.132:4352",
					"creation_date": "2014-10-28T07:00:42.739043+00:00",
					"id": "2da015cc-765b-4848-a207-427bd16a3ece",
					"volume_type": "l_ssd",
					"server": {
						"id": "b43f78aa-c856-4d1d-9a47-73a503ea27a9",
						"name": "seq-02"
					}
				},
				"1": {
					"size": 100000000000,
					"name": "volume",
					"modification_date": "2014-10-28T06:57:47.834355+00:00",
					"organization": "8d768494-0f72-4929-81ca-b8e157e61a49",
					"export_uri": "nbd://10.1.14.132:4353",
					"creation_date": "2014-10-28T07:00:42.739043+00:00",
					"id": "ef4048b4-9225-4140-bfd2-65b3405d078c",
					"volume_type": "l_hdd",
					"server": {
						"id": "b43f78aa-c856-4d1d-9a47-73a503ea27a9",
						"name": "seq-02"
					}
				}
			},
			"organization": "8d768494-0f72-4929-81ca-b8e157e61a49"
		}`)

	var server Server
	err := json.Unmarshal(b, &server)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("server.PublicIp: %#v\n", server.PublicIp)
	fmt.Printf("server.Image: %#v\n", server.Image)

	exp := "my_server"
	if server.Name != exp {
		t.Fatalf("Expected %s, got %s", exp, server.Name)
	}

	exp = "running"
	if server.State != exp {
		t.Fatalf("Expected %s, got %s", exp, server.State)
	}

	exp = "212.47.234.149"
	if server.PublicIp.Address != exp {
		t.Fatalf("Expected %s, got %s", exp, server.PublicIp.Address)
	}

}
