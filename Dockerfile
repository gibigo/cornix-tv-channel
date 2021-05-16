FROM golang:alpine as builder
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
# add build dependencies
RUN apk add --no-cache gcc musl-dev upx
# build the binary
RUN go build -o cornix-tv-channel -ldflags="-s" .
# compress the binary
RUN upx -9 -k cornix-tv-channel

# build the runtime image (apline because of cgo dependecies)
FROM alpine:latest
COPY --from=builder /app/cornix-tv-channel .
EXPOSE 3000
ENTRYPOINT [ "./cornix-tv-channel" ]