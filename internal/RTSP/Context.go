package RTSP

import (
	"net/url"
)

type Context struct {
	url    *url.URL
	method string
	req    Request
	resp   Response

	Keys map[string]interface{}
}

func NewContext() *Context {
	c := &Context{
		Keys: make(map[string]interface{}),
	}
	return c
}
