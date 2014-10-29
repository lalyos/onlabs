package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/lalyos/onlabs/online"
)

func cmdListImages() {
	cl := online.NewClient(os.Getenv("ONLINE_TOKEN"))
	images, err := cl.Images()
	if err != nil {
		log.Error(err)
		return
	}

	for _, i := range images {
		fmt.Printf(" %-40s %s\n", i.Name, i.Id)
	}
}

func cmdListServers() {
	cl := online.NewClient(os.Getenv("ONLINE_TOKEN"))
	servers, err := cl.Servers()
	if err != nil {
		log.Error(err)
		return
	}

	for _, s := range servers {
		fmt.Printf(" %-15s %-15s %s %s\n", s.Name, s.PublicIp.Address, s.Id, s.Image.Name)
	}
}

func cmdListVolumes() {
	cl := online.NewClient(os.Getenv("ONLINE_TOKEN"))
	all, err := cl.Volumes()
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-20s %s %16d\n", o.Name, o.Id, o.Size)
	}
}

func cmdListSnapshots() {
	cl := online.NewClient(os.Getenv("ONLINE_TOKEN"))
	all, err := cl.Snapshots()
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-20s %s %16d\n", o.Name, o.Id, o.Size)
	}
}
