package api

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
)

// used by CRUD helper tests

// Example is an example of how to use the mocking
func Example() {
	// go run . & sleep 2 && curl -k --location --request POST "http://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op": "showConfig", "path": ["firewall", "ipv4", "name", "SrvMain-WanTelia"]}'

	addr := "localhost:8080"
	srv := &http.Server{Addr: addr}

	// Input
	uri := "/retrieve"
	key := "test-key"
	ops := `{"op": "showConfig", "path": ["firewall", "ipv4", "name", "SrvMain-WanTelia"]}`

	// Output
	response := `{"success": true, "data": {"default-action": "reject", "default-log": {}, "description": "Managed by terraform", "rule": {"1": {"action": "accept", "description": "Allow established connections", "protocol": "all", "state": "established"}, "2": {"action": "accept", "description": "Allow related connections", "protocol": "all", "state": "related"}, "3": {"action": "drop", "description": "Disallow invalid packets", "log": {}, "protocol": "all", "state": "invalid"}, "1000": {"action": "accept", "description": "Allow http outgoing traffic", "destination": {"group": {"port-group": "Web"}}, "protocol": "tcp"}}}, "error": null}`

	el := NewExchangeList()
	e := el.Add()
	e.Expect(uri, key, ops)
	e.Response(200, response)
	Server(srv, el)

	/*
		... make http requests here to test that they match
	*/
}

// NewExchangeList constructs a new list used to add new exchanges
func NewExchangeList() *ExchangeList {
	return &ExchangeList{}
}

// ExchangeList holds and handles exchanges to be matched against
type ExchangeList struct {
	exchanges []*Exchange
}

// Add appends a new exchange to the list and returns a reference of the exchange
func (el *ExchangeList) Add() *Exchange {
	e := &Exchange{
		matched: false,
	}
	el.exchanges = append(el.exchanges, e)
	return e
}

func (el *ExchangeList) handle(w http.ResponseWriter, r *http.Request) (foundMatch bool) {
	for i, e := range el.Unmatched() {
		log.Printf("Checking unmatched pattern %d of %d", i+1, len(el.Unmatched()))
		if e.handle(w, r) {
			return true
		}
	}

	return false
}

// Unmatched returns all exchanges that did not trigger a match so far, useful for inspection
func (el *ExchangeList) Unmatched() []*Exchange {
	var ret []*Exchange
	for _, e := range el.exchanges {
		if !e.matched {
			ret = append(ret, e)
		}
	}

	return ret
}

// Exchange holds information used to match requests sent to the Server
type Exchange struct {
	matched  bool
	expect   expect
	response response
}

// Expect configures how the incoming request is expected to look
func (e *Exchange) Expect(uri, key string, ops string) *Exchange {
	e.expect = expect{
		uri: uri,
		key: key,
		ops: ops,
	}

	return e
}

// Response configures how the Exchange should respond when matched
func (e *Exchange) Response(code int, body string) *Exchange {
	e.response = response{
		code: code,
		body: body,
	}

	return e
}

func (e *Exchange) handle(w http.ResponseWriter, r *http.Request) (ok bool) {
	if !e.expect.matches(r) {
		return false
	}

	if !e.response.reply(w) {
		panic("unable to send reply")
	}

	log.Printf("Matched: %v", e.expect)
	e.matched = true
	return true
}

type expect struct {
	uri, key string
	ops      string
}

func (e expect) matches(r *http.Request) bool {
	if !(e.uri == r.RequestURI && e.key == r.FormValue("key")) {
		return false
	}

	formData := r.FormValue("data")
	if e.ops != formData {
		log.Printf("FormData mismatch!\t")
		log.Printf("Expected: '%#v'\t", e.ops)
		log.Printf("Got: '%#v'\n", formData)
		return false
	}

	return true

}

type response struct {
	code int
	body string
}

func (r response) reply(w http.ResponseWriter) (ok bool) {
	_, err := w.Write([]byte(r.body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	return true
}

// Server starts and maintains the http server until all exchanges are matched
func Server(srv *http.Server, el *ExchangeList) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if !el.handle(w, r) {
			log.Printf("No match found for:\n")
			log.Printf("URI: %s\n", r.RequestURI)
			log.Printf("Key: %s\n", r.FormValue("key"))
			log.Printf("Ops: %s\n", r.FormValue("data"))

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("NO MATCH"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			log.Printf("Flushed\n")
		}
		if len(el.Unmatched()) == 0 {
			// Only works if the response.reply() does a Flush() on the http.ResponseWriter,
			//   otherwise it will close the server before response is sent
			// Run as a go func as being in this handler and shutting down the server is self-blocking
			go srv.Shutdown(context.TODO())
			log.Printf("Srv Shutdown\n")
		}
	})

	srv.Handler = mux

	// Split out listen and serve functions to reduce chances of server not being ready when test starts
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatalf("error starting listner: %s\n", err)
	}

	go func() {
		err := srv.Serve(l)
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server closed\n")
		} else if err != nil {
			log.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}
