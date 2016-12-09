package snapstats

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
)

const ApiVersion = "v1"

var (
	ErrNotFound     = errors.New("Not Found")
	ErrAuthFailure  = errors.New("Authentication failed")
	ErrAccessDenied = errors.New("Access denied")
	ErrNilResponse  = errors.New("Nil response")
)

type Message struct {
	Body *Body `json:"body"`
}

type Body struct {
	ScheduledTasks []Task
}

type Task struct {
	HitCount    int    `json:"hit_count"`
	FailedCount int    `json:"failed_count"`
	TaskState   string `json:"task_state"`
}

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
	u.Path = path.Clean(u.Path + "/" + ApiVersion)
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

func (c *Client) get(path string, query interface{}) ([]byte, error) {
	if query != nil {
		qstr, err := ToQueryString(query)
		if err != nil {
			return nil, err
		}
		path = path + "?" + qstr
	}
	log.Printf("sending request for %s", c.prefix+path)
	req, err := http.NewRequest("GET", c.prefix+path, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	return handleResp(rsp)
}

func handleResp(rsp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode == 401 {
		return nil, ErrAuthFailure
	}
	if rsp.StatusCode == 403 {
		return nil, ErrAccessDenied
	}
	if rsp.StatusCode == 404 {
		return nil, ErrNotFound
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Unknown error encountered. %s", rsp.Status)
	}

	return b, nil
}

// Convert an interface{} to a urlencoded querystring
func ToQueryString(q interface{}) (string, error) {
	v, err := query.Values(q)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (c *Client) GetTasks() ([]Task, error) {
	body, err := c.get("/tasks", nil)
	if err != nil {
		return nil, err
	}
	message := Message{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}

	return message.Body.ScheduledTasks, nil
}
