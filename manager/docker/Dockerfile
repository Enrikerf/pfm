FROM debian:buster AS compiler

ARG GO_VERSION
RUN apt update && apt install -y curl
WORKDIR /tmp
RUN URL=https://storage.googleapis.com/golang/go1.${GO_VERSION}.linux-amd64.tar.gz \
    && curl ${URL} -o go.tar.gz  \
    && tar -C /usr/local -xzf go.tar.gz  \
    && rm -rf go.tar.gz

ENV PATH $PATH:/usr/local/go/bin
# If you enable this, then gcc is needed to debug your app
ENV CGO_ENABLED 0

FROM compiler AS dev

WORKDIR /app
EXPOSE 8000 40000
RUN apt update && apt install -y curl vim fish \
    && groupadd -g 1000 docker-user \
    && useradd -d /home/docker-user -s /bin/bash -u 1000 -g 1000 docker-user \
    && mkdir /home/docker-user \
    && chown -R docker-user:docker-user /home/docker-user \
    && go install github.com/go-delve/delve/cmd/dlv@latest

FROM compiler AS compile

WORKDIR /app
COPY ./app .
RUN go build -ldflags "-s -w" -o final.sh

FROM debian:buster AS prod

WORKDIR /
EXPOSE 8000
COPY --from=compile /app/final.sh /
CMD ["./final.sh"]