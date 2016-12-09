package snapstats

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	URL    *url.URL
	http   *http.Client
	prefix string
}

func NewClient(serverUrl string, insecure bool) (*Client, error) {
	u, err := url.Parse(serverUrl)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("URL %s is not in the format of http(s)://<ip>:<port>", serverUrl)
	}
	c := &Client{
		URL: u,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
		prefix: u.String(),
	}
	return c, nil
}

func (c *Client) GetTasks() {
	// body, err := c.get("/monitoring/jobs", nil)
	// if err != nil {
	// 	return nil, err
	// }

}
