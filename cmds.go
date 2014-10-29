package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/lalyos/online/onlabs"
)

func cmdListImages() {
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
