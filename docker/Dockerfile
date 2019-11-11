FROM alpine:latest

RUN mkdir -p /opt/OnyxChain && \
    mkdir -p /opt/OnyxChain/cert



ENV PATH /opt/OnyxChain:$PATH
WORKDIR /opt/OnyxChain

RUN apk update && apk add git curl jq npm bash && \
    npm install -g wscat

RUN apk add --no-cache --virtual .build-deps gcc musl-dev openssl go && \
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



RUN echo -e " \n\
"#!/bin/bash" \n\
checkStatus() { \n\
        if [ \"\${1}\" != \"SUCCESS\" ]; then \n\
                echo \"syncnode was reboot -- \$(date)\" >> /opt/OnyxChain/Log/checkOnyxChainHealth.log \n\
                logger -p user.error -t \$(basename \$0) \"syncnode was reboot -- \$(date)\" \n\
                reboot \n\
        fi \n\
} \n\
checkAPI=\$(curl --insecure -sS https://localhost:20334/api/v1/block/height | jq -r '.Desc') \n\
checkWebsocket=\$(wscat -c \"ws://localhost:20335\" -x '{\"Action\": \"heartbeat\",\"Version\": \"1.0.0\"}' | jq -r '.Desc') \n\
checkRPC=\$(curl --insecure -sS -X POST -H \"Content-type: application/json\"  -H \"Accept: application/json\" -d '{\"jsonrpc\":\"2.0\",\"method\":\"getbestblockhash\",\"params\":[],\"id\":1}'  https://localhost:20336 | jq -r '.desc') \n\
"# checkStatus api, websocket, rpc" \n\
checkStatus \${checkAPI} \n\
checkStatus \${checkWebsocket} \n\
checkStatus \${checkRPC} \n\
" > /opt/checkOnyxChainHealth.sh

RUN chmod +x /opt/checkOnyxChainHealth.sh && \
    echo '* * * * * /opt/checkOnyxChainHealth.sh' >> /etc/crontabs/root


CMD ['crond', '-l 2', '-f']
CMD [ -z "${domain}" ] && /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws \
|| /opt/OnyxChain/OnyxChain -w /opt/OnyxChain/wallet.json -a ${address} -p ${password} --networkid ${ntworkid} --gasprice 500 --rest --ws --cert-file /etc/letsencrypt/live/${domain}/fullchain.pem --key-file /etc/letsencrypt/live/${domain}/privkey.pem
