package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/gitnishantjha/Api_Gateway/pkg/config"
)

type ServiceProxy struct {
	Service *config.Service
	proxy   *httputil.ReverseProxy
	counter uint64
}

func NewServiceProxy(service *config.Service) *ServiceProxy {
	sp := &ServiceProxy{
		Service: service,
	}
	sp.proxy = &httputil.ReverseProxy{

		Director: func(req *http.Request) {
			backendUrl := sp.selectBackend()
			if backendUrl == nil {
				log.Printf("No available backends for service: %s", sp.Service.Name)
				return
			}
			log.Printf("Forwarding request to backend: %s", backendUrl.String())

			req.URL.Scheme = backendUrl.Scheme
			req.URL.Host = backendUrl.Host
			req.Host = backendUrl.Host
		},
	}
	return sp
}

func (sp *ServiceProxy) selectBackend() *url.URL {
	if len(sp.Service.Servers) == 0 {
		return nil
	}

	index := atomic.AddUint64(&sp.counter, 1) % uint64(len(sp.Service.Servers))

	backend, _ := url.Parse(sp.Service.Servers[index])

	return backend
}

func (sp *ServiceProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sp.proxy.ServeHTTP(w, r)
}
