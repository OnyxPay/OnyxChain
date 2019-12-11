/*
 * Copyright (C) 2019 The onyxchain Authors
 * This file is part of The onyxchain library.
 *
 * The onyxchain is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The onyxchain is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The onyxchain.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package jsonrpc privides a function to start json rpc server
package jsonrpc

import (
	"context"
	"net/http"
	"strconv"

	"fmt"

	"github.com/OnyxPay/OnyxChain/common/aux"
	cfg "github.com/OnyxPay/OnyxChain/common/config"
	"github.com/OnyxPay/OnyxChain/common/log"
	"github.com/OnyxPay/OnyxChain/http/base/rpc"
)

func StartRPCServer() (aux.Stopper, error) {
	log.Debug()
	http.HandleFunc("/", rpc.Handle)

	rpc.HandleFunc("getbestblockhash", rpc.GetBestBlockHash)
	rpc.HandleFunc("getblock", rpc.GetBlock)
	rpc.HandleFunc("getblockcount", rpc.GetBlockCount)
	rpc.HandleFunc("getblockhash", rpc.GetBlockHash)
	rpc.HandleFunc("getconnectioncount", rpc.GetConnectionCount)
	//HandleFunc("getrawmempool", GetRawMemPool)

	rpc.HandleFunc("getrawtransaction", rpc.GetRawTransaction)
	rpc.HandleFunc("sendrawtransaction", rpc.SendRawTransaction)
	rpc.HandleFunc("getstorage", rpc.GetStorage)
	rpc.HandleFunc("getversion", rpc.GetNodeVersion)
	rpc.HandleFunc("getnetworkid", rpc.GetNetworkId)

	rpc.HandleFunc("getcontractstate", rpc.GetContractState)
	rpc.HandleFunc("getmempooltxcount", rpc.GetMemPoolTxCount)
	rpc.HandleFunc("getmempooltxstate", rpc.GetMemPoolTxState)
	rpc.HandleFunc("getsmartcodeevent", rpc.GetSmartCodeEvent)
	rpc.HandleFunc("getblockheightbytxhash", rpc.GetBlockHeightByTxHash)

	rpc.HandleFunc("getbalance", rpc.GetBalance)
	rpc.HandleFunc("getallowance", rpc.GetAllowance)
	rpc.HandleFunc("getmerkleproof", rpc.GetMerkleProof)
	rpc.HandleFunc("getblocktxsbyheight", rpc.GetBlockTxsByHeight)
	rpc.HandleFunc("getgasprice", rpc.GetGasPrice)
	rpc.HandleFunc("getunboundoxg", rpc.GetUnboundOxg)
	rpc.HandleFunc("getgrantoxg", rpc.GetGrantOxg)

	port := int(cfg.DefConfig.Rpc.HttpJsonPort)
	certPath := cfg.DefConfig.Rpc.HttpCertPath
	keyPath := cfg.DefConfig.Rpc.HttpKeyPath
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: nil}
	var err error
	if len(certPath) > 0 && len(keyPath) > 0 {
		err = server.ListenAndServeTLS(certPath, keyPath)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		return nil, fmt.Errorf("ListenAndServe error:%s", err)
	}
	stopper := func() error {
		return server.Shutdown(context.Background())
	}
	return stopper, nil
}
