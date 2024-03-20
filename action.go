package main

type ActionId int

const (
	ACT_NAME ActionId = iota
	ACT_JOIN
	ACT_ROOMS
	ACT_MSG
	ACT_LEAVE
)

type Action struct {
	id ActionId
	client *Client
	args []string
}