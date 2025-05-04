<h1>Chat and Notification Service for Flat Buddies Application</h1>

### Setup
```console
cd flat-buddies-chat-notification-service
docker-compose up -d
npm install -g wscat
```

<br><b> Run the code using the following command (Terminal 1)</b>

``` console
go run cmd/main.go
```
<br> Response on command line
```console
🚀 Server running on :8080 (Chat API + Notification Listener)
```
### Setup WebSocket for Notifications (Terminal 2)
```console
wscat -c "ws://localhost:8080/ws?user_id=user2"
```
<br> Response on Terminal 2
```console
Connected (press CTRL+C to quit)
```
<br> Response on Terminal 1
```console
WebSocket connected: user2
```

### To send a chat message (Terminal 3, if using curl on cmd, or use clients like Postman)
```console
curl -X POST http://localhost:8080/send -d '{                                     
  "sender_id":"user1",
  "receiver_id":"user2",              
  "content":"Hi there!"
}' -H "Content-Type: application/json"
```
<br> Response on command line - Kafka consumer of Notification Service (Terminal 1)
```console
New message for user2: Message from user1: Hi there!
```
<br> Response on command line - Web Socket forward of notification (Terminal 2)
```console
{"user_id":"user2","type":"message","content":"Message from user1: Hi there!","timestamp":1746341915}
```

### To send a flatmate match notification (Terminal 3, if using curl on cmd, or use clients like Postman)
```console
curl -X POST http://localhost:8080/match -d '{                                     
  "user_id": "user2",
  "content": "You matched with user1!"
}' -H "Content-Type: application/json"
```
<br> Response on command line - Kafka consumer of Notification Service (Terminal 1)
```console
Match found for user2: You matched with user1!
```
<br> Response on command line - Web Socket forward of notification (Terminal 2)
```console
{"user_id":"user2","type":"match","content":"You matched with user1!","timestamp":1746341941}
```
