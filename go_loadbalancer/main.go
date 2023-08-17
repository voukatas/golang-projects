package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	hostname string
	proxy    *httputil.ReverseProxy
}

func (s *simpleServer) Address() string {
	return s.hostname
}

func (s *simpleServer) IsAlive() bool { return true }

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)

}

func newSimpleServer(hostname string) *simpleServer {
	serverUrl, err := url.Parse(hostname)
	if err != nil {
		os.Exit(1)
	}

	return &simpleServer{
		hostname: hostname,
		proxy:    httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {

	return &LoadBalancer{
		port:    port,
		servers: servers,
	}

}

func (lb *LoadBalancer) getNextAvailiableServer() Server {

	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {

		lb.roundRobinCount++

		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailiableServer()

	fmt.Printf("forwarding request to %q\n", targetServer.Address())

	targetServer.Serve(rw, req)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.google.com"),
		newSimpleServer("https://www.yahoo.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("5000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at %s", lb.port)
	http.ListenAndServe(":"+lb.port, nil)

}
