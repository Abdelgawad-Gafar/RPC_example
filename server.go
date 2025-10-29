package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// ChatService manages messages
type ChatService struct {
	messages []string
	mu       sync.Mutex
}

// Args structure to receive both name and message
type MessageArgs struct {
	Name    string
	Message string
}

// SendMessage method for sending a message and returning history
func (c *ChatService) SendMessage(args MessageArgs, reply *[]string) error {
	c.mu.Lock()
	formattedMsg := fmt.Sprintf("%s: %s", args.Name, args.Message)
	c.messages = append(c.messages, formattedMsg)
	*reply = make([]string, len(c.messages))
	copy(*reply, c.messages)
	c.mu.Unlock()
	fmt.Printf("%s\n", formattedMsg)
	return nil
}

// GetHistory RPC method for fetching chat history only
func (c *ChatService) GetHistory(_ int, reply *[]string) error {
	c.mu.Lock()
	*reply = make([]string, len(c.messages))
	copy(*reply, c.messages)
	c.mu.Unlock()
	return nil
}

func main() {
	chatService := new(ChatService)
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Chat server running on port 1234...")

	rpc.Register(chatService)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
