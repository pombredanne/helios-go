package helios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

const (
	srvService  = "helios"
	srvProtocol = "http"
)

type Client struct {
	// HTTP client used to talk to Helios.
	client *http.Client

	BaseURL *url.URL

	Hosts *HostsService
}

// NewClient returns a new Helios client for talking to (one of) the
// Helios masters in the provided domain. The Helios masters are discovered
// via SRV lookup. Optionally you may also provide your own http.Client.
func NewClient(domain string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	_, addrs, err := net.LookupSRV(srvService, srvProtocol, domain)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(fmt.Sprintf("http://%v:%d/", addrs[0].Target, addrs[0].Port))
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.Hosts = &HostsService{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(p)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if v != nil {
		return json.Unmarshal(body, v)
	}

	return nil
}
