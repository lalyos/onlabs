package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	docopt "github.com/docopt/docopt-go"
)

const version = "v0.1.8"

func main() {
	usage := `Usage:
  onlabs  images    [--verbose|-v]
  onlabs  servers   [--verbose|-v]
  onlabs  volumes   [--verbose|-v]
  onlabs  snapshots [--verbose|-v]
  onlabs  ips       [--verbose|-v]
  onlabs  actions   --server=SERVERID  [--verbose|-v]
  onlabs  reboot    --server=SERVERID  [--verbose|-v]
  onlabs  poweron   --server=SERVERID  [--verbose|-v]
  onlabs  poweroff  --server=SERVERID  [--verbose|-v]
  onlabs  terminate --server=SERVERID  [--verbose|-v]

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

	if arguments["ips"].(bool) {
		cmdListIPs()
	}

	if arguments["actions"].(bool) {
		cmdListActions(arguments["--server"].(string))
	}

	if arguments["reboot"].(bool) {
		cmdDoActions(arguments["--server"].(string), "reboot")
	}

	if arguments["poweron"].(bool) {
		cmdDoActions(arguments["--server"].(string), "poweron")
	}

	if arguments["poweroff"].(bool) {
		cmdDoActions(arguments["--server"].(string), "poweroff")
	}

	if arguments["terminate"].(bool) {
		cmdDoActions(arguments["--server"].(string), "terminate")
	}

}
