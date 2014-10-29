package online

import (
	"encoding/json"
	"testing"
)

func TestUnmarshallIP(t *testing.T) {
	b := []byte(`{
			"address": "212.47.226.88",
			"id": "b50cd740-892d-47d3-8cbf-88510ef626e7",
			"organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			"server": null
		}`)

	// d as domain
	var d IP
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := "b50cd740-892d-47d3-8cbf-88510ef626e7"
	if d.Id != exp {
		t.Fatalf("Expected %#v, got %#v", exp, d.Id)
	}

	exp = "212.47.226.88"
	if d.Address != exp {
		t.Fatalf("Expected %#v, got %#v", exp, d.Address)
	}

}

func TestGetIPsResp(t *testing.T) {
	b := []byte(`{
  "ips": [
    {
      "address": "212.47.226.88",
      "id": "b50cd740-892d-47d3-8cbf-88510ef626e7",
      "organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
      "server": null
    },
		{
			"address": "212.47.226.99",
			"id": "b50cd740-892d-47d3-9999-88510ef626e7",
			"organization": "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			"server": null
		}

  ]
}`)

	// d as domain
	var d GetIPsResp
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}

	exp := 2
	if len(d.IPs) != exp {
		t.Fatalf("Expected %#v, got %#v", exp, len(d.IPs))
	}
}
