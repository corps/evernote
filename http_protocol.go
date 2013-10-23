package evernote

import (
	"io"
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"
)

const (
	DEFAULT_BUFFER_SIZE = 1024
)

type THttpClientTransportError struct {
	Message			string
	Err   			error
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
	URL 			string
	RequestBuffer  *bytes.Buffer	
	ResponseBuffer *bytes.Buffer
}

func NewTHttpClientTransport(URL string) *THttpClientTransport {
	return &THttpClientTransport{Client: http.DefaultClient, URL: URL, 
		RequestBuffer: bytes.NewBuffer(make([]byte, 0, DEFAULT_BUFFER_SIZE)),
		ResponseBuffer: bytes.NewBuffer(make([]byte, 0, DEFAULT_BUFFER_SIZE)),
	}
}

func (t *THttpClientTransport) Close() error {
	t.Client = nil
	t.RequestBuffer = nil
	t.ResponseBuffer = nil
	return nil
}

func (t *THttpClientTransport) doHttpFlush() error {
	fmt.Println("Doing flush")	
	requestData, err := ioutil.ReadAll(t.RequestBuffer)
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

	_, err = io.Copy(t.ResponseBuffer, resp.Body)
	fmt.Println("Finished copying flush")	

	if err != nil {
		return THttpClientTransportError{Err: err}
	}

	return nil
}

func (t *THttpClientTransport) Write(p []byte) (n int, e error) {
	fmt.Println("Doing write")	
	return t.RequestBuffer.Write(p)
}

func (t *THttpClientTransport) Read(p []byte) (n int, e error) {
	fmt.Println("Doing read")	
	if len(t.RequestBuffer.Bytes()) > 0 {
		e = t.doHttpFlush()
		if e != nil {
			return
		}
	}
	return t.ResponseBuffer.Read(p)
}
