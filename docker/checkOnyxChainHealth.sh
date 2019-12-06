#!/bin/bash
hostUrl="https://localhost"
checkHttp=$(curl --connect-timeout 10 --insecure -sS $hostUrl:20334/api/v1/block/height)
if grep -q "$checkHttp" <<< "SSL"; then
        hostUrl="http://localhost"
fi
checkStatus() {
        if [ "${1}" != "SUCCESS" ]; then
                echo "syncnode was reboot -- $(date)" >> /opt/OnyxChain/Log/checkOnyxChainHealth.log
                logger -p user.error -t $(basename $0) "syncnode was reboot -- $(date)"
                kill -9 $(ps ax | grep OnyxChain | awk '{print $1}')
        fi
}
checkAPI=$(curl --connect-timeout 10 --insecure -sS $hostUrl:20334/api/v1/block/height | jq -r '.Desc')
checkRPC=$(curl --connect-timeout 10 --insecure -sS -X POST -H "Content-type: application/json"  -H "Accept: application/json" -d '{"jsonrpc":"2.0","method":"getbestblockhash","params":[],"id":1}'  $hostUrl:20336 | jq -r '.desc')
# checkStatus api, rpc
checkStatus ${checkAPI}
checkStatus ${checkRPC}
# check websocket
checkWebsocket=$(echo '{"Action": "heartbeat","Version": "1.0.0"}' | websocat wss://localhost:20335 | jq -r '.Desc')
if [ -z "$checkWebsocket" ]; then
checkWebsocket=$(echo '{"Action": "heartbeat","Version": "1.0.0"}' | websocat ws://localhost:20335 | jq -r '.Desc')
fi
checkStatus ${checkWebsocket}
