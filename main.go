package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	docopt "github.com/docopt/docopt-go"
)

const version = "v0.1.5"

func main() {
	usage := `Usage:
	online  images [--verbose|-v]

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

}
