package injector

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/kyawmyintthein/twirp-poc/app/config"
)

type ConnectionWatcher struct {
	n int64
}

// OnStateChange records open connections in response to connection
// state changes. Set net/http Server.ConnState to this method
// as value.
func (cw *ConnectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		atomic.AddInt64(&cw.n, 1)
	case http.StateHijacked, http.StateClosed:
		atomic.AddInt64(&cw.n, -1)
	}

	// rzlog.InfoKV(context.Background(), rzlog.KV{"count": int(atomic.LoadInt64(&cw.n)), "remote_addr": conn.RemoteAddr()}, "cw watcher")
}

// Count returns the number of connections at the time
// the call.
func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))

}

func ProvideHTTPServer(config *config.GlobalCfg) *http.Server {
	var cw ConnectionWatcher
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.App.ServerPort),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
		ConnState:    cw.OnStateChange,
	}
	return httpServer
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
