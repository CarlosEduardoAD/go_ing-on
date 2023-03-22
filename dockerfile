FROM golang:1.20.1-bullseye  

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

CMD ["tail", "-f", "/dev/null"]