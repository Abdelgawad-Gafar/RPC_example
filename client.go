package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

// الهيكل المستخدم لإرسال الاسم والرسالة للسيرفر
type MessageArgs struct {
	Name    string
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Unable to connect to server:", err)
	}
	defer client.Close()

	// أخذ اسم المستخدم في البداية
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg == "exit" {
			fmt.Println("Exiting chat.")
			break
		}

		args := MessageArgs{Name: name, Message: msg}
		var history []string
		err := client.Call("ChatService.SendMessage", args, &history)
		if err != nil {
			log.Println("Error sending message or fetching history:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, line := range history {
			fmt.Println(line)
		}
		fmt.Println("--------------------")
	}
}
