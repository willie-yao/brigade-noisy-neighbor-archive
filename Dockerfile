FROM brigadecore/go-tools:v0.1.0

ARG VERSION
ARG COMMIT
ENV CGO_ENABLED=0

WORKDIR /
COPY . /
COPY go.mod go.mod
COPY go.sum go.sum

RUN go build \
  -o bin/gateway \
  -ldflags "-w -X github.com/willie-yao/brigade-noisy-neighbor/internal/version.version=$VERSION -X github.com/willie-yao/brigade-noisy-neighbor/internal/version.commit=$COMMIT" \
  .

EXPOSE 8080

FROM scratch
COPY --from=0 /bin/ /brigade-noisy-neighbor/bin/
ENTRYPOINT ["/brigade-noisy-neighbor/bin/gateway"]
