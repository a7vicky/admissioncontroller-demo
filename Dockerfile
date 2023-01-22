FROM golang:1.19 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /work
COPY . ./

RUN go mod download
# Build admission-webhook
RUN  make build


# ---
FROM redhat/ubi8-minimal
ENV USER_UID=1001 \
    USER_NAME=webhooks

COPY --from=builder  /work/bin/admission-webhook-demo /usr/local/bin/
COPY build/bin/* /usr/local/bin/
RUN  chmod +x /usr/local/bin/user_setup
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}