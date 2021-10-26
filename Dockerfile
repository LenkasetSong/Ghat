### build ###
FROM golang:alpine AS build-env 

# 1. build backend 
RUN mkdir -p /backend 
WORKDIR /backend
ADD ./backend/ /backend/
RUN go build -o backend_docker 
# 2. build server for frontend 
RUN mkdir -p /server 
WORKDIR /server 
ADD ./server /server/
RUN go build -o server_docker 


### run ###
FROM alpine
RUN mkdir -p /build 
WORKDIR /
# server 
COPY --from=build-env /server/server_docker /
# backend 
COPY --from=build-env /backend/backend_docker /
# frontend 
ADD ./frontend/build/ /build/
RUN echo -e "#!/bin/sh\n ./server_docker & \n ./backend_docker" > /start.sh 
# frontend port 
EXPOSE 8080 
# backend port 
EXPOSE 8081
CMD [ "sh", "start.sh" ]
