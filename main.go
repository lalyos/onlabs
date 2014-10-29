package main

import (
	log "github.com/Sirupsen/logrus"
	docopt "github.com/docopt/docopt-go"
)

func init() {
}

func main() {
	usage := `Usage:
	onlabs  images [--verbose|-v]

Options:
  -h --help         this message
  -v --verbose      verbose mode`

	arguments, _ := docopt.Parse(usage, nil, true, "Online Labs Client: v0.1.0", false)

	if arguments["--verbose"].(bool) {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("args: %v", arguments)

	if arguments["images"].(bool) {
		cmdListImages()
	}

}
