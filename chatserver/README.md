# WebSocket Server & Client in Go and JavaScript
This repository contains a simple WebSocket server in Go and two client implementations (one in Go and one in HTML/JavaScript), providing a base for any application that requires real-time, bi-directional communication between the server and clients.

# Server
The WebSocket server keeps track of all active connections, broadcasts received messages to all connected clients, and handles connection errors. This server listens for connections on port 3001.

To run the server:

bash
```
go run server.go
```
# Go Client
The Go WebSocket client connects to the server, sends a greeting message, and then continuously listens for incoming messages from the server.

To run the Go client:

```
go run client.go
```

# HTML/JavaScript Client
The HTML/JavaScript client opens a WebSocket connection to the server, sends a greeting message when the connection is opened, and logs any messages received from the server to the console.

To use the HTML/JavaScript client:

Copy the client.html file to your web server directory.
```
sudo cp client.html /var/www/html/
sudo chmod 755 /var/www/html/client.html
```
Open a web browser and navigate to http://localhost/client.html.
You can also send messages from the browser console with:

```
socket.send("New message!")
```
# Dependencies
The only external dependency for the Go server and client is golang.org/x/net/websocket, which provides the basic WebSocket implementation.

You can install it with:

```
go get golang.org/x/net/websocket
```
# Note
Please ensure that you have Go installed on your machine and that your GOPATH is properly set up before running the Go server and client. Also, ensure that you have a web server installed and running for the HTML/JavaScript client.