package api

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// Router is the basic interface for all implemented routers
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter().StrictSlash(true)
)

// NewMuxRouter creates a new mux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) SERVE(port string) {
	// Setup middleware routes for swagger docu
	// handler for documentation
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	muxDispatcher.Handle("/docs", sh)
	muxDispatcher.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//serve webserver on specific port
	http.ListenAndServe(port, muxDispatcher)
	log.Fatalf("Listening on port %v", port)

}
