# TCP Client-Server (Case Study for DFS Project)

## 📖 Overview
This project is a **basic TCP client-server application** written in Go.  
It serves as the **second step in my journey toward building a distributed file system (DFS)**.  

The goal of this case study is to build **networking foundations** — understanding how clients and servers communicate over TCP, how to manage multiple connections, and how to send/receive data reliably. These skills are essential before moving on to file transfer, replication, and distributed coordination in a DFS.  

---

## 🛠 Features
- **Server**:  
  - Listens for incoming TCP connections on port `8080`.  
  - Handles multiple clients concurrently using goroutines.  
  - Echoes messages back to the sender.  

- **Client**:  
  - Connects to the server.  
  - Lets the user type messages into the terminal.  
  - Receives and displays the server’s responses in real time.  

---

## ▶️ Usage

### 1. Run the server
```bash
go run server.go
```


### 2. Run the client
```bash
go run client.go
```

### 3. Chat

Type messages in the client terminal, and the server will echo them back.
