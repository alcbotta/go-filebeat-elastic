# syntax=docker/dockerfile:1
FROM golang:1.16
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /workspace/app ./
CMD ["./app"]  