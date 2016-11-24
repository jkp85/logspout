package main

import (
	_ "github.com/3Blades/logspout/adapters/raw"
	_ "github.com/3Blades/logspout/adapters/syslog"
	_ "github.com/3Blades/logspout/httpstream"
	_ "github.com/3Blades/logspout/routesapi"
	_ "github.com/3Blades/logspout/transports/tcp"
	_ "github.com/3Blades/logspout/transports/udp"
	_ "github.com/3Blades/logspout/transports/tls"
)
