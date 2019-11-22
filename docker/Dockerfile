# cd ./docker
# docker build -t onyxchain .

FROM alpine:latest

RUN mkdir -p /opt/OnyxChain && \
    mkdir -p /opt/OnyxChain/cert



ENV PATH /opt/OnyxChain:$PATH
WORKDIR /opt/OnyxChain

RUN apk update && apk add git curl jq npm bash && \
    npm install -g wscat

ENV GO111MODULE="on"

RUN apk add --no-cache --virtual .build-deps gcc musl-dev openssl go && \
    wget -O go.tgz https://dl.google.com/go/go1.12.src.tar.gz  && \
    tar -C /usr/local -xzf go.tgz && \
    cd /usr/local/go/src/ && \
    ./make.bash && \
    export PATH="/usr/local/go/bin:$PATH" && \
    export GOPATH=/opt/go/  && \
    export PATH=$PATH:$GOPATH/bin && \
    apk del .build-deps && \
    go get github.com/OnyxPay/OnyxChain && \
    mv /opt/go/bin/OnyxChain /opt/OnyxChain/ && \
    rm -rf /opt/go/ && \
    rm -rf /usr/local/go/ && \
    rm -r /opt/OnyxChain/go.tgz



COPY ./checkOnyxChainHealth.sh /opt/checkOnyxChainHealth.sh

RUN chmod +x /opt/checkOnyxChainHealth.sh && \
    echo '* * * * * /opt/checkOnyxChainHealth.sh' >> /etc/crontabs/root


CMD ['crond', '-l 2', '-f']
CMD [ -z "${domain}" ] && /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws \
|| /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws --cert-file /etc/letsencrypt/live/${domain}/fullchain.pem --key-file /etc/letsencrypt/live/${domain}/privkey.pem
