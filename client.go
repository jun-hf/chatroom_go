package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	name string
	room *Room
	actions chan Action
}

func (client *Client) readInput() {
	for {
		msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		action := strings.TrimSpace(args[0])

		switch action {
		case "/name":
			client.actions <- Action{
				id: ACT_NAME,
				client: client,
				args: args,
			}
		case "/join":
			client.actions <- Action{
				id: ACT_JOIN,
				client: client,
				args: args,
			} 
		default:
			client.error(fmt.Errorf("invalid action: %s", action))
		}
	}
}

func (client *Client) error(err error) {
	client.conn.Write([]byte("Error!!: " + err.Error() + "\n"))
}

func (client *Client) msg(msg string) {
	client.conn.Write([]byte(">>> " + msg +"\n"))
}