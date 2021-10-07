FROM alpine:3.10

WORKDIR /usr/bin

COPY restsanity ./
RUN chmod +x ./restsanity
EXPOSE 5001
ENTRYPOINT ./restsanity --port 5001
