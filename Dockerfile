FROM alpine

COPY message-persistence-service /message-persistence-service

ENTRYPOINT ["./message-persistence-service"]
