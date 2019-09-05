FROM alpine:latest

RUN mkdir -p /opt/OnyxChain && \
    mkdir -p /opt/OnyxChain/cert



ENV PATH /opt/OnyxChain:$PATH
WORKDIR /opt/OnyxChain

RUN apk update && apk add git

RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go && \
    wget -O go.tgz https://dl.google.com/go/go1.10.3.src.tar.gz  && \
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


CMD /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws


