FROM alpine:latest
RUN apk update
RUN apk add libvirt-client ncurses5-libs git gettext curl bash make
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN ln -s /usr/lib/libncurses.so.5 /usr/lib/libtinfo.so.5

# add binary
ADD dist/linux_amd64/tshub /bin/tshub
ADD dist/linux_amd64/tshub-fs /bin/tshub-fs

# start tshub
CMD ["/bin/tshub"]