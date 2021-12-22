package core

import (
	"net/http"
	"os"
)

type (
	Core interface {
		Run()
	}

	core struct {
		server *http.Server
	}
)

func NewCore() Core {
	addr_port := os.ExpandEnv("$IP:$PORT")
	return &core{server: &http.Server{Addr: addr_port,}} 
}

func (c *core) Run() {
	c.server.Handler = c
	c.server.ListenAndServe()
}

func (c *core) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("stok"))
}
