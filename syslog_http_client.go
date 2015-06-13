package httplog

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type httpConn struct {
	url    string
	client http.Client
}

func httpSyslog(proto, addr string) (serverConn, error) {
	h := &httpConn{
		url: proto + "://" + addr + "/syslog",
	}
	resp, err := h.client.Get(h.url)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK || resp.Header.Get("syslog") != "ready" {
		return nil, errors.New("server in not an httplogsys endpoint")
	}
	return h, nil
}

func (h *httpConn) writeString(p Priority, hostname, tag, msg, nl string) error {

	var body bytes.Buffer

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Fprintf(&body, "<%d>%s %s %s[%d]: %s%s", p, timestamp, hostname, tag, os.Getpid(), msg, nl)

	resp, err := h.client.Post(h.url, "text/syslog", &body)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusNoContent {
		return errors.New("rejected log")
	}
	return nil
}

func (h *httpConn) close() error {
	return nil
}
