FROM scratch

COPY message-persistence-service /message-persistence-service

ENTRYPOINT ["./message-persistence-service"]
