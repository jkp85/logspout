package main

import (
	_ "github.com/rkorkosz/logspout/adapters/raw"
	_ "github.com/rkorkosz/logspout/adapters/syslog"
	_ "github.com/rkorkosz/logspout/httpstream"
	_ "github.com/rkorkosz/logspout/routesapi"
	_ "github.com/rkorkosz/logspout/transports/tcp"
	_ "github.com/rkorkosz/logspout/transports/udp"
	_ "github.com/rkorkosz/logspout/transports/tls"
)
