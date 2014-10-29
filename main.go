package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	docopt "github.com/docopt/docopt-go"
)

const version = "v0.1.5"

func main() {
	usage := `Usage:
  onlabs  images    [--verbose|-v]
  onlabs  servers   [--verbose|-v]
  onlabs  volumes   [--verbose|-v]
  onlabs  snapshots [--verbose|-v]

Options:
  -h --help         this message
  -v --verbose      verbose mode`

	ver := fmt.Sprintf("Online Labs Client: %s", version)
	arguments, _ := docopt.Parse(usage, nil, true, ver, false)

	if arguments["--verbose"].(bool) {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("args: %v", arguments)

	if arguments["images"].(bool) {
		cmdListImages()
	}

	if arguments["servers"].(bool) {
		cmdListServers()
	}

	if arguments["volumes"].(bool) {
		cmdListVolumes()
	}

	if arguments["snapshots"].(bool) {
		cmdListSnapshots()
	}

}
