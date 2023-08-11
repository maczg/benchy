FROM golang:1.19 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o benchy

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/benchy .
USER 65532:65532

ENTRYPOINT ["/benchy"]