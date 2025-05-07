<h1>Chat and Notification Service for Flat Buddies Application</h1>

### Chat service deployment on minikube
```console
cd flat-buddies-chat-notification-service
minikube start
eval $(minikube docker-env)
cd chat-service
docker build -t chat-service:latest .
docker images
cd ..
kubectl apply -f chat-deployment.yaml
kubectl get pods
```

### Notification service deployment on minikube
```console
cd flat-buddies-chat-notification-service
eval $(minikube docker-env)
cd notification-service
docker build -t notification-service:latest .
docker images
cd ..
kubectl apply -f notification-deployment.yaml
kubectl get pods
```

### Kafka and Zookeeper deployment on minikube
```console
kubectl apply -f kafka-deployment.yaml
kubectl get pods
```
### Test Chat service - messages(Terminal 1 - use curl or clients like POSTMAN)
```console
minikube service chat-service --url
```
<br> Note the URL to send requests
```console
curl -X POST http://127.0.0.1:<port>/send -d '{
  "sender_id":"user1",
  "receiver_id":"user2",
  "content":"Hi there! Is this flat available?"
}' -H "Content-Type: application/json"
```
<br> Response
```console
Message sent to Kafka
```

### Test notification service for chat messages (Terminal 2)
```console
kubectl get pods
```
<br> Note the pod_name for notification-service
```console
kubectl logs notification-service-<pod-id>
```
<br> Response
```
📬 New chat from user1 to user2: Hi there! Is this flat available?
```
### Test Chat service - match (Terminal 1 - use curl or clients like POSTMAN)

```console
curl -X POST http://127.0.0.1:<port>/notification -d '{
  "user1_id": "user1",
  "user2_id": "user2",
  "content": "You matched with user2"
}' -H "Content-Type: application/json"
```
<br> Response
```console
{
    "message": "Match notification sent to Kafka",
    "user1": "user1",
    "user2": "user2"
}
```

### Test notification service for matches (Terminal 2)
<br> Response
```
📬 New chat from user1 to user2: You are matched with user1
```

