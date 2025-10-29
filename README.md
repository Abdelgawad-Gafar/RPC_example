# Simple Go RPC Chatroom

This project implements a simple chatroom using Go with RPC communication between a server and multiple clients.

## Requirements

- Go installed on your system ([https://golang.org/dl/](https://golang.org/dl/))
- Any terminal application (such as Git Bash, Command Prompt, PowerShell, or VS Code Terminal)

## How to Run

1. **Prepare the Code**
   - Place both `server.go` and `client.go` in the same folder.

2. **Start the Server**
   - Open a terminal window in the project directory.
   - Run:
     ```
     go run server.go
     ```
   - You should see:
     ```
     Chat server running on port 1234...
     ```

3. **Start Clients**
   - Open two or more new terminal windows.
   - In each client terminal, start the client:
     ```
     go run client.go
     ```
   - Each client will be prompted to enter their name, then can send messages.
   - Every message will appear in the format:  
     `ClientName: message`
   - Type `exit` to close a client.

## Demo Video

Watch the demo video here:  
[Google Drive Demo Video](https://docs.google.com/videos/d/1EKEditXLn9a4Tgpxeb92eiW5UsDoNk1oRCLX550i_lQ/edit?usp=sharing)



## Notes

- All messages are shown to all clients in real-time, with the sender's name clearly indicated.
- To open multiple terminals:
  - Open several windows of your terminal application and resize them for side-by-side viewing, as shown in the screenshot.
- If you encounter connection issues, make sure the server is running and the IP/port matches (`127.0.0.1:1234`).

---

## License

This project is for educational purposes and RPC/distributed systems practice.
