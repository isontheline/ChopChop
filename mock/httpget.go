package mock

import (
	"bytes"
	"fmt"
	"gochopchop/internal/httpget"
	"io/ioutil"
	"net/http"
)

type FakeNetClient map[string]*http.Response

func (f FakeNetClient) Get(url string) (*http.Response, error) {
	// implements IHTTPClient interface
	if res, ok := f[url]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("could not get url : %s", url)
}

func (f FakeNetClient) Do(req *http.Request) (*http.Response, error) {
	// implements IHTTPClient interface
	if res, ok := f[req.URL.String()]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("could not do on url : %s", req.URL.String())
}

var urls = FakeNetClient{
	"url1": &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("foo"))),
	},
}

var FakeFetcher = &httpget.Fetcher{
	Netclient: urls,
}
