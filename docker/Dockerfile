FROM golang

COPY root /

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh

RUN go get github.com/oxequa/realize

WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]

CMD ["realize", "start"]
