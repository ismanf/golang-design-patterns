package decorator

import (
	"os"
	"io"
	"fmt"
	"net/http"
)

// simple http server
type SimpleServer struct{
	msg string
}
func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.msg)
}

// log decorator for simple server
type ServerWithLog struct {
	handler http.Handler
	logger  io.Writer
}
func (s *ServerWithLog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(s.logger, "Got request with uri: %s \n", r.RequestURI)
	s.handler.ServeHTTP(w, r)
}

func Demo(){
	http.Handle("/", &ServerWithLog{&SimpleServer{"Simple server"}, os.Stdout})
	http.ListenAndServe(":3000", nil)
}
