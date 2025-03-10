FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/ai-agent
RUN cd /go/src/github.com/victorzhou123/ai-agent && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

COPY --from=BUILDER /go/src/github.com/victorzhou123/ai-agent/ai-agent /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/ai-agent/config/config.yaml /opt/app/config/

ENTRYPOINT ["/opt/app/ai-agent"]
