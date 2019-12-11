# docker build -f docker/Dockerfile -t onyxchain .

FROM golang:1-alpine AS builder

WORKDIR /go/src/app

COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./... && go install -v ./...


FROM alpine:latest

RUN apk add --no-cache bash curl jq \
    && mkdir -p /opt/OnyxChain /var/log/checkhealth \
    && curl -LR https://github.com/vi/websocat/releases/download/v1.5.0/websocat_amd64-linux-static -o /usr/local/bin/websocat \
    && chmod +x /usr/local/bin/websocat

WORKDIR /opt/OnyxChain

COPY --from=builder /go/bin/OnyxChain .

COPY docker/checkOnyxChainHealth.sh /opt

HEALTHCHECK --interval=60s --timeout=5s --start-period=60s --retries=1 CMD bash /opt/checkOnyxChainHealth.sh

CMD [ -z "${domain}" ] \
    && /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws \
    || /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws --cert-file /etc/letsencrypt/live/${domain}/fullchain.pem --key-file /etc/letsencrypt/live/${domain}/privkey.pem
