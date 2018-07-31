package livy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"io/ioutil"

	"strconv"

	"github.com/pkg/errors"
)

const (
	apiVersion     = "0.5.0-incubating"
	libraryVersion = "0.1.0"
	mediaType      = "application/json"
)

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	userAgent string
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	u, _ := url.Parse(baseURL)

	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL:   u,
		userAgent: fmt.Sprintf("go-livy/%s livy/%s %s (%s/%s)", libraryVersion, apiVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}

func (c *Client) NewRequest(method, resource string, payload interface{}) (*http.Request, error) {
	rel, err := url.Parse(resource)
	if err != nil {
		return nil, err
	}
	u := c.baseURL.ResolveReference(rel)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling JSON")
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil

}

func (c *Client) Do(req *http.Request, into interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		errorMessage, _ := strconv.Unquote(string(body))
		return errors.Errorf("server returned status: %s: %s", resp.Status, errorMessage)
	}

	if into == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "reading response body")
	}

	//fmt.Println(string(body))

	if err := json.Unmarshal(body, into); err != nil {
		return errors.Wrap(err, "decoding response body")
	}

	return nil
}
