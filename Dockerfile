FROM golang

ARG projectDir

RUN go get github.com/oxequa/realize

WORKDIR $projectDir

CMD ["realize", "start"]
