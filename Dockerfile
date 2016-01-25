FROM busybox
WORKDIR /app
ADD ./keys /app/keys
ADD ./hangouts-call-center_linux-amd64 /app/run
CMD ["/app/run", "http"]
