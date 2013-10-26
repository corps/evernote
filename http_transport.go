package evernote

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	DEFAULT_BUFFER_SIZE = 1024
)

type THttpClientTransportError struct {
	Message string
	Err     error
}

func (e THttpClientTransportError) Error() string {
	msg := e.Message
	if len(msg) == 0 {
		msg = e.Err.Error()
	}
	return fmt.Sprintf("THttpClientTransport Error: %s\n", msg)
}

type THttpClientTransport struct {
	Client         *http.Client
	URL            string
	requestBuffer  *bytes.Buffer
	responseBuffer *bytes.Buffer
	responses      chan *bytes.Buffer
}

func NewTHttpClientTransport(URL string) *THttpClientTransport {
	return NewTHttpClientTransportWithHttpClient(URL, http.DefaultClient)
}

func NewTHttpClientTransportWithHttpClient(URL string, Client *http.Client) *THttpClientTransport {
	return &THttpClientTransport{Client: Client, URL: URL,
		requestBuffer:  bytes.NewBuffer(make([]byte, 0, DEFAULT_BUFFER_SIZE)),
		responseBuffer: nil,
		responses:      make(chan *bytes.Buffer),
	}
}

func (t *THttpClientTransport) Close() error {
	t.Client = nil
	t.requestBuffer = nil
	t.responseBuffer = nil
	return nil
}

func (t *THttpClientTransport) Flush() error {
	requestData, err := ioutil.ReadAll(t.requestBuffer)
	if err != nil {
		return THttpClientTransportError{Err: err}
	}

	request, err := http.NewRequest("POST", t.URL, bytes.NewReader(requestData))
	if err != nil {
		return THttpClientTransportError{Err: err}
	}

	request.Header.Add("Content-Type", "application/x-thrift")
	request.Header.Add("Accept", "application/x-thrift")
	request.Header.Add("User-Agent", "github.com/corps/evernote/THttpClientTransport")

	resp, err := t.Client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return THttpClientTransportError{Err: err}
	}

	if resp.StatusCode != 200 {
		return THttpClientTransportError{Message: fmt.Sprintf("Unexpected status code: %s", resp.StatusCode)}
	}

	responseBuffer := bytes.NewBuffer(make([]byte, 0, DEFAULT_BUFFER_SIZE))
	_, err = io.Copy(responseBuffer, resp.Body)
	go func() { t.responses <- responseBuffer }()

	if err != nil {
		return THttpClientTransportError{Err: err}
	}

	return nil
}

func (t *THttpClientTransport) Write(p []byte) (n int, e error) {
	if t.requestBuffer == nil {
		return 0, io.EOF
	}

	return t.requestBuffer.Write(p)
}

func (t *THttpClientTransport) Read(p []byte) (n int, e error) {
	if t.requestBuffer == nil {
		return 0, io.EOF
	}

	if t.responseBuffer == nil || t.responseBuffer.Len() == 0 {
		t.responseBuffer = <-t.responses
	}

	return t.responseBuffer.Read(p)
}
