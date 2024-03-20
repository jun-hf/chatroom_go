package main

import "net"

type Room struct {
	name string
	members map[net.Addr]*Client
}

func (room *Room) broadcast(sender *Client, msg string) {
	for address, member := range(room.members) {
		if sender.conn.RemoteAddr() != address {
			member.msg(msg)
		}
	}
}