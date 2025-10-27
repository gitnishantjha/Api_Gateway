package router

import (
	"log"
	"net/http"

	"github.com/gitnishantjha/Api_Gateway/pkg/config"
	"github.com/gitnishantjha/Api_Gateway/pkg/middleware"
	"github.com/gitnishantjha/Api_Gateway/pkg/proxy"
)

func NewRouter(cfg *config.Config) (http.Handler, error) {
	mux := http.NewServeMux()

	serviceMap := make(map[string]*proxy.ServiceProxy)
	for i := range cfg.Services {
		service := &cfg.Services[i]
		serviceProxy := proxy.NewServiceProxy(service)
		serviceMap[service.Name] = serviceProxy
	}

	middlewareRegistry := map[string]middleware.Middleware{
		"logging":      middleware.Logging,
		"api_key_auth": middleware.APIKeyAuth,
	}
	log.Println("configuring Routes")

	for _, endpoint := range cfg.EndPoints {

		proxyHandler, ok := serviceMap[endpoint.Service]
		if !ok {
			log.Fatalf("Service '%s' for endpoint '%s' not found", endpoint.Service, endpoint.Path)
		}

		var mwChain []middleware.Middleware

		for _, mwName := range endpoint.Middlewares {
			if mw, exists := middlewareRegistry[mwName]; exists {
				mwChain = append(mwChain, mw)
			} else {
				log.Printf("Warning: Middleware '%s' not found, skipping.", mwName)
			}
		}

		finalHandler := middleware.Chain(proxyHandler, mwChain...)

		mux.Handle(endpoint.Path, finalHandler)
		log.Printf("  Registered endpoint: %s -> %s (Middlewares: %v)", endpoint.Path, endpoint.Service, endpoint.Middlewares)

	}
	return mux, nil
}
