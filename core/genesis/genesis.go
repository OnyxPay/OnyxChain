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

package genesis

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/OnyxPay/OnyxChain-crypto/keypair"
	"github.com/OnyxPay/OnyxChain/common"
	"github.com/OnyxPay/OnyxChain/common/config"
	"github.com/OnyxPay/OnyxChain/common/constants"
	"github.com/OnyxPay/OnyxChain/consensus/vbft/config"
	"github.com/OnyxPay/OnyxChain/core/types"
	"github.com/OnyxPay/OnyxChain/core/utils"
	"github.com/OnyxPay/OnyxChain/smartcontract/service/native/global_params"
	"github.com/OnyxPay/OnyxChain/smartcontract/service/native/governance"
	"github.com/OnyxPay/OnyxChain/smartcontract/service/native/onx"
	nutils "github.com/OnyxPay/OnyxChain/smartcontract/service/native/utils"
	"github.com/OnyxPay/OnyxChain/smartcontract/service/neovm"
)

const (
	BlockVersion uint32 = 0
	GenesisNonce uint64 = 2083236893
)

var (
	ONXToken   = newGoverningToken()
	OXGToken   = newUtilityToken()
	ONXTokenID = ONXToken.Hash()
	OXGTokenID = OXGToken.Hash()
)

var GenBlockTime = (config.DEFAULT_GEN_BLOCK_TIME * time.Second)

var INIT_PARAM = map[string]string{
	"gasPrice": "0",
}

var GenesisBookkeepers []keypair.PublicKey

// BuildGenesisBlock returns the genesis block with default consensus bookkeeper list
func BuildGenesisBlock(defaultBookkeeper []keypair.PublicKey, genesisConfig *config.GenesisConfig) (*types.Block, error) {
	//getBookkeeper
	GenesisBookkeepers = defaultBookkeeper
	nextBookkeeper, err := types.AddressFromBookkeepers(defaultBookkeeper)
	if err != nil {
		return nil, fmt.Errorf("[Block],BuildGenesisBlock err with GetBookkeeperAddress: %s", err)
	}
	conf := bytes.NewBuffer(nil)
	if genesisConfig.VBFT != nil {
		genesisConfig.VBFT.Serialize(conf)
	}
	govConfig := newGoverConfigInit(conf.Bytes())
	consensusPayload, err := vconfig.GenesisConsensusPayload(govConfig.Hash(), 0)
	if err != nil {
		return nil, fmt.Errorf("consensus genesis init failed: %s", err)
	}
	//blockdata
	genesisHeader := &types.Header{
		Version:          BlockVersion,
		PrevBlockHash:    common.Uint256{},
		TransactionsRoot: common.Uint256{},
		Timestamp:        constants.GENESIS_BLOCK_TIMESTAMP,
		Height:           uint32(0),
		ConsensusData:    GenesisNonce,
		NextBookkeeper:   nextBookkeeper,
		ConsensusPayload: consensusPayload,

		Bookkeepers: nil,
		SigData:     nil,
	}

	//block
	onx := newGoverningToken()
	oxg := newUtilityToken()
	param := newParamContract()
	oid := deployOnxIDContract()
	auth := deployAuthContract()
	govConfigTx := newGovConfigTx()

	genesisBlock := &types.Block{
		Header: genesisHeader,
		Transactions: []*types.Transaction{
			onx,
			oxg,
			param,
			oid,
			auth,
			govConfigTx,
			newGoverningInit(),
			newUtilityInit(),
			newParamInit(),
			govConfig,
		},
	}
	genesisBlock.RebuildMerkleRoot()
	return genesisBlock, nil
}

func newGoverningToken() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.OnxContractAddress[:], "ONYX", "1.0",
		"OnyxChain Team", "contact@onyxpay.co", "OnyxChain Network ONYX Token", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis governing token transaction error ")
	}
	return tx
}

func newUtilityToken() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.OxgContractAddress[:], "OXG", "1.0",
		"OnyxChain Team", "contact@onyxpay.co", "OnyxChain Network OXG Token", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis utility token transaction error ")
	}
	return tx
}

func newParamContract() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.ParamContractAddress[:],
		"ParamConfig", "1.0", "OnyxChain Team", "contact@onyxpay.co",
		"Chain Global Environment Variables Manager ", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis param transaction error ")
	}
	return tx
}

func newGovConfigTx() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.GovernanceContractAddress[:], "CONFIG", "1.0",
		"OnyxChain Team", "contact@onyxpay.co", "OnyxChain Network Consensus Config", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis config transaction error ")
	}
	return tx
}

func deployAuthContract() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.AuthContractAddress[:], "AuthContract", "1.0",
		"OnyxChain Team", "contact@onyxpay.co", "OnyxChain Network Authorization Contract", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis auth transaction error ")
	}
	return tx
}

func deployOnxIDContract() *types.Transaction {
	mutable := utils.NewDeployTransaction(nutils.OnxIDContractAddress[:], "OID", "1.0",
		"OnyxChain Team", "contact@onyxpay.co", "OnyxChain Network ONX ID", true)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis onxid transaction error ")
	}
	return tx
}

func newGoverningInit() *types.Transaction {
	bookkeepers, _ := config.DefConfig.GetBookkeepers()

	var addr common.Address
	if len(bookkeepers) == 1 {
		addr = types.AddressFromPubKey(bookkeepers[0])
	} else {
		m := (5*len(bookkeepers) + 6) / 7
		temp, err := types.AddressFromMultiPubKeys(bookkeepers, m)
		if err != nil {
			panic(fmt.Sprint("wrong bookkeeper config, caused by", err))
		}
		addr = temp
	}

	distribute := []struct {
		addr  common.Address
		value uint64
	}{{addr, constants.ONX_TOTAL_SUPPLY}}

	args := bytes.NewBuffer(nil)
	nutils.WriteVarUint(args, uint64(len(distribute)))
	for _, part := range distribute {
		nutils.WriteAddress(args, part.addr)
		nutils.WriteVarUint(args, part.value)
	}

	mutable := utils.BuildNativeTransaction(nutils.OnxContractAddress, onx.INIT_NAME, args.Bytes())
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis governing token transaction error ")
	}
	return tx
}

func newUtilityInit() *types.Transaction {
	mutable := utils.BuildNativeTransaction(nutils.OxgContractAddress, onx.INIT_NAME, []byte{})
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis utility token transaction error ")
	}

	return tx
}

func newParamInit() *types.Transaction {
	params := new(global_params.Params)
	var s []string
	for k := range INIT_PARAM {
		s = append(s, k)
	}

	for k, v := range neovm.INIT_GAS_TABLE {
		INIT_PARAM[k] = strconv.FormatUint(v, 10)
		s = append(s, k)
	}

	sort.Strings(s)
	for _, v := range s {
		params.SetParam(global_params.Param{Key: v, Value: INIT_PARAM[v]})
	}
	bf := new(bytes.Buffer)
	params.Serialize(bf)

	bookkeepers, _ := config.DefConfig.GetBookkeepers()
	var addr common.Address
	if len(bookkeepers) == 1 {
		addr = types.AddressFromPubKey(bookkeepers[0])
	} else {
		m := (5*len(bookkeepers) + 6) / 7
		temp, err := types.AddressFromMultiPubKeys(bookkeepers, m)
		if err != nil {
			panic(fmt.Sprint("wrong bookkeeper config, caused by", err))
		}
		addr = temp
	}
	nutils.WriteAddress(bf, addr)

	mutable := utils.BuildNativeTransaction(nutils.ParamContractAddress, global_params.INIT_NAME, bf.Bytes())
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis governing token transaction error ")
	}
	return tx
}

func newGoverConfigInit(config []byte) *types.Transaction {
	mutable := utils.BuildNativeTransaction(nutils.GovernanceContractAddress, governance.INIT_CONFIG, config)
	tx, err := mutable.IntoImmutable()
	if err != nil {
		panic("construct genesis governing token transaction error ")
	}
	return tx
}
