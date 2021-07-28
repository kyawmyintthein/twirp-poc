package injector

import (
	"fmt"
	"net/http"

	"github.com/kyawmyintthein/twirp-poc/app/config"
)

func ProvideHTTPServer(config *config.GlobalCfg) *http.Server {
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", config.App.ServerPort),
	}
	return httpServer
}
