# docker build -f docker/Dockerfile -t onyxchain .

FROM golang:1-alpine AS builder

WORKDIR /go/src/app

COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./... && go install -v ./...


FROM alpine:latest

RUN apk add --no-cache jq \
    && mkdir -p /opt/OnyxChain /var/log/cron /var/log/checkhealth

WORKDIR /opt/OnyxChain

ADD https://github.com/vi/websocat/releases/download/v1.5.0/websocat_amd64-linux-static /usr/local/bin/websocat

COPY --from=builder /go/bin/OnyxChain .

COPY docker/checkOnyxChainHealth.sh /opt

RUN chmod +x /usr/local/bin/websocat /opt/checkOnyxChainHealth.sh && \
    echo '* * * * * /opt/checkOnyxChainHealth.sh >> /var/log/checkhealth/check.log' >> /etc/crontabs/root

ENV PATH=/opt/OnyxChain:$PATH

CMD crond -l 2 -L /var/log/cron/cron.log \
    && [ -z "${domain}" ] \
    && /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws \
    || /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws --cert-file /etc/letsencrypt/live/${domain}/fullchain.pem --key-file /etc/letsencrypt/live/${domain}/privkey.pem
