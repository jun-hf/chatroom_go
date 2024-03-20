package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	rooms   map[string]*Room
	actions chan Action
}

func CreateServer() *Server {
	return &Server{
		rooms:   make(map[string]*Room),
		actions: make(chan Action),
	}
}

func (server *Server) run() {
	for action := range server.actions {
		switch action.id {
		case ACT_NAME:
			server.name(action.client, action.args)
		case ACT_JOIN:
			server.join(action.client, action.args)
		case ACT_ROOMS:
			server.listRoom(action.client)
		case ACT_MSG:
			server.msg(action.client, action.args)
		case ACT_LEAVE:
			server.leave(action.client)
		}

	}
}

func (server *Server) newClient(conn net.Conn) *Client {
	log.Printf("New client had joined the server: %s", conn.RemoteAddr().String())

	return &Client{
		conn:    conn,
		name:    "anonymous",
		actions: server.actions,
	}
}

func (server *Server) name(client *Client, args []string) {
	if len(args) < 2 {
		client.msg("name is required. usage:/name <NAME>")
		return
	}
	client.name = args[1]
	client.msg(fmt.Sprintf("Hello 😀 %s", client.name))
}

func (server *Server) join(client *Client, args []string) {
	if len(args) < 2 {
		client.msg("room name is required. usage:/join <Room Name>")
		return
	}

	roomName := args[1]

	r, ok := server.rooms[roomName]
	if !ok {
		r = &Room{
			name: roomName,
			members: make(map[net.Addr]*Client),
		}
		server.rooms[roomName] = r
	}
	r.members[client.conn.RemoteAddr()] = client
	server.leaveCurrentRoom(client)
	client.room = r

	r.broadcast(client, fmt.Sprintf("%s joined the room", client.name))

	client.msg(fmt.Sprintf("Welcome to %s", roomName))
}

func (server *Server) leaveCurrentRoom(c *Client) {
	if c.room != nil {
		oldRoom := server.rooms[c.room.name]
		delete(server.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.name))
	}
}

func (server *Server) listRoom(c *Client) {
	var rooms []string
	for name := range(server.rooms) {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("available rooms to join: %s", strings.Join(rooms, ", ")))
}

func (server *Server) msg(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("message is required. usage:/msg <MSG>")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.name + ": "+msg)
}

func (server *Server) leave(c *Client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	server.leaveCurrentRoom(c)

	c.msg("Goodbye")
	c.conn.Close()
}


