package mock

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/bpicode/fritzctl/assert"
	"github.com/julienschmidt/httprouter"
)

// Fritz represents the mock of the FB.
type Fritz struct {
	LoginChallengeResponse string
	LoginResponse          string
	Server                 *httptest.Server
}

// New creates a new *mock.Fritz with default configuration.
func New() *Fritz {
	return &Fritz{
		LoginChallengeResponse: "../mock/login_challenge.xml",
		LoginResponse:          "../mock/login_response_success.xml",
	}
}

// Start creates and starts the server.
func (f *Fritz) Start() *Fritz {
	server := f.UnstartedServer()
	server.Start()
	f.Server = server
	return f
}

// Close closes the server.
func (f *Fritz) Close() {
	f.Server.Close()
}

// UnstartedServer sets up the routes and creates a server.
func (f *Fritz) UnstartedServer() *httptest.Server {
	router := f.fritzRoutes()
	return httptest.NewUnstartedServer(router)
}

func (f *Fritz) fritzRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/login_sid.lua", f.loginHandler)
	return router
}

func (f *Fritz) loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.URL.Query().Get("response") == "" {
		f.writeFromFs(w, f.LoginChallengeResponse)
	} else {
		f.writeFromFs(w, f.LoginResponse)
	}
}

func (f *Fritz) writeFromFs(w http.ResponseWriter, path string) {
	file, err := os.Open(path)
	assert.NoError(err)
	defer file.Close()
	io.Copy(w, file)
}
