// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	extNet        = flag.Bool("extnet", false, "do external network tests")
	transportHost = flag.String("transporthost", "http2.golang.org", "hostname to use for TestTransport")
	insecure      = flag.Bool("insecure", false, "insecure TLS dials") // TODO: dead code. remove?
)

var tlsConfigInsecure = &tls.Config{InsecureSkipVerify: true}

func TestTransportExternal(t *testing.T) {
	if !*extNet {
		t.Skip("skipping external network test")
	}
	req, _ := http.NewRequest("GET", "https://"+*transportHost+"/", nil)
	rt := &Transport{TLSClientConfig: tlsConfigInsecure}
	res, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("%v", err)
	}
	res.Write(os.Stdout)
}

func TestTransport(t *testing.T) {
	const body = "sup"
	st := newServerTester(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, body)
	}, optOnlyServer)
	defer st.Close()

	tr := &Transport{TLSClientConfig: tlsConfigInsecure}
	defer tr.CloseIdleConnections()

	req, err := http.NewRequest("GET", st.ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	t.Logf("Got res: %+v", res)
	if g, w := res.StatusCode, 200; g != w {
		t.Errorf("StatusCode = %v; want %v", g, w)
	}
	if g, w := res.Status, "200 OK"; g != w {
		t.Errorf("Status = %q; want %q", g, w)
	}
	wantHeader := http.Header{
		"Content-Length": []string{"3"},
		"Content-Type":   []string{"text/plain; charset=utf-8"},
	}
	if !reflect.DeepEqual(res.Header, wantHeader) {
		t.Errorf("res Header = %v; want %v", res.Header, wantHeader)
	}
	if res.Request != req {
		t.Errorf("Response.Request = %p; want %p", res.Request, req)
	}
	if res.TLS == nil {
		t.Error("Response.TLS = nil; want non-nil")
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Body read: %v", err)
	} else if string(slurp) != body {
		t.Errorf("Body = %q; want %q", slurp, body)
	}

}

func TestTransportReusesConns(t *testing.T) {
	st := newServerTester(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.RemoteAddr)
	}, optOnlyServer)
	defer st.Close()
	tr := &Transport{TLSClientConfig: tlsConfigInsecure}
	defer tr.CloseIdleConnections()
	get := func() string {
		req, err := http.NewRequest("GET", st.ts.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		res, err := tr.RoundTrip(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		slurp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Body read: %v", err)
		}
		addr := strings.TrimSpace(string(slurp))
		if addr == "" {
			t.Fatalf("didn't get an addr in response")
		}
		return addr
	}
	first := get()
	second := get()
	if first != second {
		t.Errorf("first and second responses were on different connections: %q vs %q", first, second)
	}
}

func TestTransportAbortClosesPipes(t *testing.T) {
	shutdown := make(chan struct{})
	st := newServerTester(t,
		func(w http.ResponseWriter, r *http.Request) {
			w.(http.Flusher).Flush()
			<-shutdown
		},
		optOnlyServer,
	)
	defer st.Close()
	defer close(shutdown) // we must shutdown before st.Close() to avoid hanging

	done := make(chan struct{})
	requestMade := make(chan struct{})
	go func() {
		defer close(done)
		tr := &Transport{TLSClientConfig: tlsConfigInsecure}
		req, err := http.NewRequest("GET", st.ts.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		res, err := tr.RoundTrip(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		close(requestMade)
		_, err = ioutil.ReadAll(res.Body)
		if err == nil {
			t.Error("expected error from res.Body.Read")
		}
	}()

	<-requestMade
	// Now force the serve loop to end, via closing the connection.
	st.closeConn()
	// deadlock? that's a bug.
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("timeout")
	}
}

// TODO: merge this with TestTransportBody to make TestTransportRequest? This
// could be a table-driven test with extra goodies.
func TestTransportPath(t *testing.T) {
	gotc := make(chan *url.URL, 1)
	st := newServerTester(t,
		func(w http.ResponseWriter, r *http.Request) {
			gotc <- r.URL
		},
		optOnlyServer,
	)
	defer st.Close()

	tr := &Transport{TLSClientConfig: tlsConfigInsecure}
	defer tr.CloseIdleConnections()
	const (
		path  = "/testpath"
		query = "q=1"
	)
	surl := st.ts.URL + path + "?" + query
	req, err := http.NewRequest("POST", surl, nil)
	if err != nil {
		t.Fatal(err)
	}
	c := &http.Client{Transport: tr}
	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	got := <-gotc
	if got.Path != path {
		t.Errorf("Read Path = %q; want %q", got.Path, path)
	}
	if got.RawQuery != query {
		t.Errorf("Read RawQuery = %q; want %q", got.RawQuery, query)
	}
}

var bodyTests = []struct {
	body         string
	noContentLen bool
}{
	{body: "some message"},
	{body: "some message", noContentLen: true},
	{body: ""},
	{body: "", noContentLen: true},
	{body: strings.Repeat("a", 1<<20), noContentLen: true},
	{body: strings.Repeat("a", 1<<20)},
}

func TestTransportBody(t *testing.T) {
	gotc := make(chan interface{}, 1)
	st := newServerTester(t,
		func(w http.ResponseWriter, r *http.Request) {
			slurp, err := ioutil.ReadAll(r.Body)
			if err != nil {
				gotc <- err
			} else {
				gotc <- string(slurp)

			}
		},
		optOnlyServer,
	)
	defer st.Close()

	for i, tt := range bodyTests {
		tr := &Transport{TLSClientConfig: tlsConfigInsecure}
		defer tr.CloseIdleConnections()

		var body io.Reader = strings.NewReader(tt.body)
		if tt.noContentLen {
			body = struct{ io.Reader }{body} // just a Reader, hiding concrete type and other methods
		}
		req, err := http.NewRequest("POST", st.ts.URL, body)
		if err != nil {
			t.Fatalf("#%d: %v", i, err)
		}
		c := &http.Client{Transport: tr}
		res, err := c.Do(req)
		if err != nil {
			t.Fatalf("#%d: %v", i, err)
		}
		defer res.Body.Close()
		got := <-gotc
		if err, ok := got.(error); ok {
			t.Fatalf("#%d: %v", i, err)
		} else if got.(string) != tt.body {
			t.Errorf("#%d: Read body = %q; want %q", i, got, tt.body)
		}
	}
}

func TestTransportDialTLS(t *testing.T) {
	var mu sync.Mutex // guards following
	var gotReq, didDial bool

	ts := newServerTester(t,
		func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			gotReq = true
			mu.Unlock()
		},
		optOnlyServer,
	)
	defer ts.Close()
	tr := &Transport{
		DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
			mu.Lock()
			didDial = true
			mu.Unlock()
			cfg.InsecureSkipVerify = true
			c, err := tls.Dial(netw, addr, cfg)
			if err != nil {
				return nil, err
			}
			return c, c.Handshake()
		},
	}
	defer tr.CloseIdleConnections()
	client := &http.Client{Transport: tr}
	res, err := client.Get(ts.ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	mu.Lock()
	if !gotReq {
		t.Error("didn't get request")
	}
	if !didDial {
		t.Error("didn't use dial hook")
	}
}
