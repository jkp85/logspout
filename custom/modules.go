package main

import (
	_ "github.com/3Blades/logspout/adapters/syslog"
	_ "github.com/3Blades/logspout/transports/tcp"
	_ "github.com/3Blades/logspout/transports/tls"
	_ "github.com/3Blades/logspout/transports/udp"
	_ "github.com/looplab/logspout-logstash"
)
