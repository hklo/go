FROM golang

ARG projectDir

ENV PROJECT_DIR ${projectDir}

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get github.com/oxequa/realize

COPY entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]

CMD ["realize", "start"]
