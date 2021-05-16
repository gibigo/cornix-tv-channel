FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o cornix-tv-forwarder .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./cornix-tv-forwarder"]