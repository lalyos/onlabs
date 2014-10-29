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
      "dynamic_public_ip": false,
      "id": "741db378-6b87-46d4-a8c5-4e46a09ab1f8",
      "image": {
        "id": "85917034-46b0-4cc5-8b48-f0a2245e357e",
        "name": "ubuntu working"
      },
      "name": "my_server",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "private_ip": null,
      "public_ip": {
        "id": "d3f55ed5-7a13-4611-92bf-c0bec332ef80",
				"dynamic": true,
        "address": "212.47.237.1"
      },
      "running": false,
      "tags": [
        "test",
        "www"
      ],
      "volumes": {
        "0": {
          "export_uri": null,
          "id": "c1eb8f3a-4f0b-4b95-a71c-93223e457f5a",
          "name": "vol simple snapshot",
          "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
          "server": {
            "id": "741db378-6b87-46d4-a8c5-4e46a09ab1f8",
            "name": "my_server"
          },
          "size": 10000000000,
          "volume_type": "l_hdd"
        }
      }
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

	exp = "212.47.237.1"
	if server.PublicIp.Address != exp {
		t.Fatalf("Expected %s, got %s", exp, server.PublicIp.Address)
	}

}
