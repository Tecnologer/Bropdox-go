package models

import (
	"time"

	"github.com/tecnologer/bropdox/models/proto"
)

type Clients []*Client

func NewClientCollection() *Clients {
	clients := Clients(make([]*Client, 0))
	return &clients
}

func (cc *Clients) Add(c *Client) {
	(*cc) = append(*cc, c)
}

func (cc *Clients) AddNewClient(id string, stream proto.Bropdox_NotificationsServer) *Client {
	client := NewClient(id, stream)
	cc.Add(client)
	return client
}

func (cc *Clients) Exists(client *Client) bool {
	for _, c := range *cc {
		if c.ID == client.ID {
			return true
		}
	}

	return false
}

type Client struct {
	ID        string
	Stream    proto.Bropdox_NotificationsServer
	lastNotif *notification
}

func NewClient(id string, stream proto.Bropdox_NotificationsServer) *Client {
	return &Client{
		ID:     id,
		Stream: stream,
	}
}

func (c *Client) Send(res *proto.Response) {
	fileRes := res.GetFileResponse()

	if c.lastNotif != nil && fileRes != nil && c.lastNotif.equals(fileRes) {
		return
	}

	c.Stream.Send(res)
}

type notification struct {
	_type     proto.TypeResponse
	path      string
	timestamp time.Time
}

func (n *notification) equals(res *proto.FileResponse) bool {
	file := res.File
	since := time.Since(n.timestamp)
	return n._type == res.Type && n.path == file.Path && since >= time.Second
}
