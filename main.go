package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/lalyos/online/onlabs"
)

func init() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}
}

func listImages() {
	cl := onlabs.NewClient(os.Getenv("ONLINE_TOKEN"))
	images, err := cl.Images()
	if err != nil {
		log.Error(err)
		return
	}

	for _, i := range images {
		fmt.Printf(" %-40s %s\n", i.Name, i.Id)
	}
}

func main() {
	fmt.Println("Online Labs Client v0.1.0")
	listImages()
}
