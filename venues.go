package fsvenues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FSClient struct {
	v             string
	client_id     string
	client_secret string
	apiUrl        string
	Client        *http.Client
}

func NewFSVenuesClient(id, secret string) FSClient {
	return FSClient{
		v:             "20131014",
		client_id:     id,
		client_secret: secret,
		apiUrl:        "https://api.foursquare.com/v2/venues/",
		Client:        http.DefaultClient,
	}
}

func (c *FSClient) GetVenues(params map[string]string) (interface{}, error) {
	return c.dispatchRequest(params, "search?")
}

func (c *FSClient) GetCategories(params map[string]string) (interface{}, error) {
	return c.dispatchRequest(params, "/categories/")
}

func (c *FSClient) ExploreVenues(params map[string]string) (interface{}, error) {
	return c.dispatchRequest(params, "/explore/")
}

func (c *FSClient) GetVenue(params map[string]string) (interface{}, error) {
	return c.dispatchRequest(params, "")
}

func (c *FSClient) dispatchRequest(params map[string]string, endpoint string) (interface{}, error) {
	var reqUrl bytes.Buffer
	reqUrl.WriteString(c.apiUrl + endpoint)
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	values.Set("client_id", c.client_id)
	values.Set("client_secret", c.client_secret)
	values.Set("v", c.v)
	reqUrl.WriteString(values.Encode())
	r, e := c.Client.Get(reqUrl.String())
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode >= 200 && r.StatusCode <= 400 && e == nil {
		var jsonResult interface{}
		json.Unmarshal(body, &jsonResult)
		return jsonResult, nil
	} else {
		return nil, fmt.Errorf("venues.go: code:%d error:%v body:%s", r.StatusCode, e, body)
	}
}
