package inadapter

import (
	"net/http"
	"strings"

	"github.com/yuraxdrumz/golang-starter-kit/internal/app/example"

	log "github.com/sirupsen/logrus"
)

// HTTPAdapter - struct with necessary use-cases for adapter to run
type HTTPAdapter struct {
	example example.Port
	port    string
}

// NewHTTPAdapter - create a new instance of NewHTTPAdapter with passed implementations
func NewHTTPAdapter(example example.Port, port string) *HTTPAdapter {
	return &HTTPAdapter{example: example, port: port}
}

// Run - initializes http server with port
func (in *HTTPAdapter) Run() {
	sayHello := func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		// call example use case passed to in adapter
		in.example.Run()
		message = strings.TrimPrefix(message, "/")
		message = "Hello " + message

		w.Write([]byte(message))
	}
	http.HandleFunc("/", sayHello)
	log.Info("Starting server on port " + in.port)
	err := http.ListenAndServe(":"+in.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
