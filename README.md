# Ghat

A chat application in Go with React.js.

- Easy to build
- Provide Dockerfile 

## Backend Port 

See **backend/main.go**:

```go
log.Fatal(http.ListenAndServe(":8081", nil))
```

## Frontend API

See **frontend/src/api/index.js**:

```js
var socket = new WebSocket("ws://localhost:8081/ws");
```

## Docker Build 

First (in frontend):

```bash
npm run build 
```

Then:

```bash
docker build -t ghat .
```

## Docker Run 

```bash
docker run -d -p 8080:8080 -p 8081:8081 --restart=always ghat
```

- Frontend port: 8080
- Backend port: 8081 

Now visit http://localhost:8080 !

Of course port mapping and reverse proxy should be done when deploying the application on a server (modify **main\.go** and **api/index\.js** accordingly)\.
