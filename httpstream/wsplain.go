package httpstream

import (
	"net/http"

	"github.com/rkorkosz/logspout/router"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

func init() {
	router.HttpHandlers.Register(WebSocketPlainLog, "ws-plain")
}

func WebSocketPlainLog() http.Handler {
	plainLogs := mux.NewRouter()
	plainLogs.HandleFunc("/ws-plain/{predicate:[a-zA-Z]+}:{value}", plainLogsHandler).Methods("GET")
	plainLogs.HandleFunc("/ws-plain", plainLogsHandler).Methods("GET")
	return plainLogs
}

func plainLogsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") != "websocket" {
		http.Error(w, "closing", 1000)
		return
	}
	params := mux.Vars(r)
	route := new(router.Route)
	if params["value"] != "" {
		switch params["predicate"] {
		case "id":
			route.FilterID = params["value"]
			if len(route.ID) > 12 {
				route.FilterID = route.FilterID[:12]
			}
		case "name":
			route.FilterName = params["value"]
		}
	}
	if route.FilterID != "" && !router.Routes.RoutingFrom(route.FilterID) {
		http.NotFound(w, r)
		return
	}
	logStream := make(chan *router.Message)
	defer close(logStream)
	var closer chan bool
	defer webSocketStreamer(w, r, logStream, closer)
	route.OverrideCloser(closer)
	router.Routes.Route(route, logStream)
}

func webSocketStreamer(w http.ResponseWriter, r *http.Request, logStream chan *router.Message, closer chan bool) {
	websocket.Handler(func(ws *websocket.Conn) {
		for logLine := range logStream {
			if r.URL.Query().Get("source") != "" && logLine.Source != r.URL.Query().Get("source") {
				continue
			}
			_, err := ws.Write([]byte(logLine.Data))
			if err != nil {
				closer <- true
				return
			}
		}
	}).ServeHTTP(w, r)
}
