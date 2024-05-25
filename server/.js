const socket = new WebSocket("ws://localhost:3000/ws");

// Event handler for when the WebSocket connection is opened
socket.addEventListener('open', function(event) {
    console.log('WebSocket connection opened');
    
    // Send a message to the server after the connection is opened
    socket.send("Hello from Chrome console");
});

// Event handler for when a message is received from the server
socket.addEventListener('message', function(event) {
    console.log('Message from server:', event.data);
});

// Event handler for errors that occur with the WebSocket connection
socket.addEventListener('error', function(event) {
    console.error('WebSocket error:', event);
});

// Event handler for when the WebSocket connection is closed
socket.addEventListener('close', function(event) {
    console.log('WebSocket connection closed');
});
