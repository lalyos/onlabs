package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/lalyos/onlabs/online"
)

func cmdListImages() {
	cl := online.NewClient()
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
	cl := online.NewClient()
	servers, err := cl.Servers()
	if err != nil {
		log.Error(err)
		return
	}

	for _, s := range servers {
		fmt.Printf("%8s %-15s %-15s %v %s\n", s.State, s.Name, s.PublicIp.Address, s.Id, s.Image.Name)
	}
}

func sizeInGB(s int) string {
	gb := s / 1000000000
	return fmt.Sprintf("%d GB", gb)
}

func cmdListVolumes() {
	cl := online.NewClient()
	all, err := cl.Volumes()
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-20s %s %10s\n", o.Name, o.Id, sizeInGB(o.Size))
	}
}

func cmdListSnapshots() {
	cl := online.NewClient()
	all, err := cl.Snapshots()
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-20s %s %10s\n", o.Name, o.Id, sizeInGB(o.Size))
	}
}

func cmdListIPs() {
	cl := online.NewClient()
	all, err := cl.IPs()
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-20s %10s\n", o.Address, o.Id)
	}
}

func cmdListActions(serverId string) {
	cl := online.NewClient()
	all, err := cl.Actions(serverId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, o := range all {
		fmt.Printf(" %-21s\n", o)
	}
}
