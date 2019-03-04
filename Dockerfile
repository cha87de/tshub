FROM alpine:latest
RUN apk update
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# add binary
ADD dist/linux_amd64/tshub /bin/tshub

# start tshub
CMD ["/bin/tshub"]
