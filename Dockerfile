# Build Stage
FROM ziwon/crawli:1.12.4 AS build-stage

LABEL app="build-crawli"
LABEL REPO="https://github.com/ziwon/crawli"

ENV PROJPATH=/go/src/github.com/ziwon/crawli

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/ziwon/crawli
WORKDIR /go/src/github.com/ziwon/crawli

RUN make build-alpine

# Final Stage
FROM ziwon/crawli:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/ziwon/crawli"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/crawli/bin

WORKDIR /opt/crawli/bin

COPY --from=build-stage /go/src/github.com/ziwon/crawli/bin/crawli /opt/crawli/bin/
RUN chmod +x /opt/crawli/bin/crawli

# Create appuser
RUN adduser -D -g '' crawli
USER crawli

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/crawli/bin/crawli"]
