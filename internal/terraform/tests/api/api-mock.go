package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	return &ExchangeList{failed: "N/A"}
}

// ExchangeList holds and handles exchanges to be matched against
type ExchangeList struct {
	exchanges []*Exchange
	failed    string
}

// Add appends a new exchange to the list and returns a reference of the exchange
func (el *ExchangeList) Add() *Exchange {
	e := &Exchange{
		matched: false,
		delay:   0,
	}
	el.exchanges = append(el.exchanges, e)
	return e
}

func (el *ExchangeList) handle(w http.ResponseWriter, r *http.Request) (foundMatch bool) {
	return el.Unmatched()[0].handle(w, r)
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

// Failed returns returns a human friendly string representation for the request that did not match
func (el *ExchangeList) Failed() string {
	return el.failed
}

// Matched returns all exchanges that triggered a match so far, useful for inspection
func (el *ExchangeList) Matched() []*Exchange {
	var ret []*Exchange
	for _, e := range el.exchanges {
		if e.matched {
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
	delay    time.Duration
}

// Expect configures how the incoming request is expected to look
func (e *Exchange) Expect(uri, key string, ops string) *Exchange {

	// If we can json marshal the ops string it will be more likely to conform to
	// the values we receive as they are likely, but not guarantied, to be from
	// other json marshalling operations

	// Check if we can match a JSON blob to maps
	unmarshalMap := make(map[string]any)
	err := json.Unmarshal([]byte(ops), &unmarshalMap)
	if err == nil {
		opsB, err := json.Marshal(unmarshalMap)
		if err == nil {
			ops = string(opsB)
		}
	}

	// Check if we can match a JSON blob to lists
	unmarshalSlice := []interface{}{}
	err = json.Unmarshal([]byte(ops), &unmarshalSlice)
	if err == nil {
		opsB, err := json.Marshal(unmarshalSlice)
		if err == nil {
			ops = string(opsB)
		}
	}

	// Crete expect and return self
	e.expect = expect{
		uri: uri,
		key: key,
		ops: ops,
	}

	return e
}

// Delay configures how long to wait before responding
func (e *Exchange) Delay(delay time.Duration) *Exchange {
	e.delay = delay

	return e
}

// Sexpect returns a human friendly string representation of the expect config
func (e *Exchange) Sexpect() string {
	return fmt.Sprintf("Expected Uri: %v\nExpected Key: %v\nExpected Ops: %v", e.expect.uri, e.expect.key, e.expect.ops)
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

	if e.delay > 0 {
		log.Printf("Mock srv:  Delaying response for: %dms", e.delay/time.Millisecond)
		time.Sleep(e.delay)
	}

	if !e.response.reply(w) {
		panic("Mock srv: unable to send reply")
	}

	log.Printf("Mock srv: Matched: %v", e.expect)
	e.matched = true
	return true
}

type expect struct {
	uri, key string
	ops      string
}

func (e expect) matches(r *http.Request) bool {
	// Check the simple parameters first to see if we can skip some work
	if !(e.uri == r.RequestURI && e.key == r.FormValue("key")) {
		return false
	}

	// sorting function for the cmp.Equal call
	// I do not know of a lazier way to do a botched "Deep Equals" that
	// don't care about slice order for deeply complex slices.
	// this should work fairly ok as we know anything that will be passed to it has
	// already been unmarashled from json.
	//
	// possible issue: If two objects contains the same
	// JSONified letters but with different meanings, eg: ["redhat", "thread"] and ["thread", "hatred"]
	// they will be considered equal for the sorting and might cause issues
	// if there are multiple of them
	sillySort := func(v1, v2 any) bool {
		v1Byte, err := json.Marshal(v1)
		if err != nil {
			panic(err)
		}
		v1StrSlice := strings.Split(string(v1Byte), "")
		slices.Sort(v1StrSlice)

		v2Byte, err := json.Marshal(v2)
		if err != nil {
			panic(err)
		}
		v2StrSlice := strings.Split(string(v2Byte), "")
		slices.Sort(v2StrSlice)

		return strings.Join(v1StrSlice, "") > strings.Join(v2StrSlice, "")
	}

	formData := r.FormValue("data")

	// Check if we can match a JSON blob to maps
	eOpsJSONM := make(map[string]any)
	formDataJSONM := make(map[string]any)
	opsErr := json.Unmarshal([]byte(e.ops), &eOpsJSONM)
	formErr := json.Unmarshal([]byte(formData), &formDataJSONM)
	if opsErr == nil && formErr == nil {
		log.Printf("Mock srv: map[string]any exchange check\n")
		return cmp.Equal(eOpsJSONM, formDataJSONM,
			cmpopts.SortMaps(sillySort),
			cmpopts.SortSlices(sillySort),
		)
	}

	// Check if we can match a JSON blob to lists
	eOpsJSONL := []interface{}{}
	formDataJSONL := []interface{}{}
	opsErr = json.Unmarshal([]byte(e.ops), &eOpsJSONL)
	formErr = json.Unmarshal([]byte(formData), &formDataJSONL)
	if opsErr == nil && formErr == nil {
		log.Printf("Mock srv: []interface{}{} exchange check\n")
		return cmp.Equal(eOpsJSONL, formDataJSONL,
			cmpopts.SortMaps(sillySort),
			cmpopts.SortSlices(sillySort),
		)
	}

	// Otherwise try to just match as a simple string
	log.Printf("Mock srv: comparing exchange expect as string\n")
	return e.ops == formData
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
// TODO change to using test server https://pkg.go.dev/net/http/httptest#Server
func Server(srv *http.Server, el *ExchangeList) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if !el.handle(w, r) {
			log.Printf("Mock srv: Request did not match the next expected exchange pattern:\n")
			el.failed = fmt.Sprintf("Requested Uri: %v\nRequested Key: %v\nRequested Ops: %v", r.RequestURI, r.FormValue("key"), r.FormValue("data"))

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("NO MATCH"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			log.Printf("Mock srv: Flushed\n")

			go srv.Shutdown(context.TODO())
			log.Printf("Mock srv: Srv Shutdown\n")
		}

		if len(el.Unmatched()) == 0 {
			// Only works if the response.reply() does a Flush() on the http.ResponseWriter,
			//   otherwise it will close the server before response is sent
			// Run as a go func as being in this handler and shutting down the server is self-blocking
			go srv.Shutdown(context.TODO())
			log.Printf("Mock srv: Srv Shutdown\n")
		}
	})

	srv.Handler = mux

	// Split out listen and serve functions to reduce chances of server not being ready when test starts
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatalf("Mock srv: error starting listner: %s\n", err)
	}

	go func() {
		err := srv.Serve(l)
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("Mock srv: server closed\n")
		} else if err != nil {
			log.Printf("Mock srv: error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	time.Sleep(50 * time.Millisecond)
}
