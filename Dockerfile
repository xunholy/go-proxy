FROM golang:alpine AS build_base
 
RUN apk update --no-cache && apk add git
 
WORKDIR /app
 
ENV GO111MODULE=on
 
COPY go.mod .
 
COPY go.sum .
 
RUN go mod download
 
FROM build_base as build
  
COPY ./ /app
 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o proxy .
 
FROM alpine

RUN apk update --no-cache && apk add cntlm && \
    mkdir -p /usr/local/etc && \
    cp /etc/cntlm.conf /usr/local/etc/cntlm.conf
 
WORKDIR /app
 
COPY --from=build /app/proxy /app
 
ENTRYPOINT ["/app/proxy"]
 
EXPOSE 8000