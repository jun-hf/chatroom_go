# chatroom_go
A chatroom built with Go

## Design
When a connection comes into the server, the server will create a client.
![image](https://github.com/jun-hf/chatroom_go/assets/86782267/b5831c7f-f580-42bb-86af-8875c805e3ed)

Within the connection, the client can publish a message into the server's channel. As an action and the server will process the action concurrently.
![image](https://github.com/jun-hf/chatroom_go/assets/86782267/a007fb1a-ba60-4d23-a881-8a2adde65fdf)

## Get started
1. `git clone` the latest version of this codebase
2. run `go run .`
3. Open a new terminal and enter `telnet localhost 8080`
4. Given yourself a name `/name <name>`
5. Join a room `/join <room name>`
6. Send a message to the room `/msg <msg>`
7. leave the room `/leave`

