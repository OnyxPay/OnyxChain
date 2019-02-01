/*
 * Copyright (C) 2018 The OnyxChain Authors
 * This file is part of The OnyxChain library.
 *
 * The OnyxChain is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OnyxChain is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The OnyxChain.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package error privides error code for http
package error

import onyxErrors "github.com/OnyxPay/OnyxChain/errors"

const (
	SUCCESS            int64 = 0
	SESSION_EXPIRED    int64 = 41001
	SERVICE_CEILING    int64 = 41002
	ILLEGAL_DATAFORMAT int64 = 41003
	INVALID_VERSION    int64 = 41004

	INVALID_METHOD int64 = 42001
	INVALID_PARAMS int64 = 42002

	INVALID_TRANSACTION int64 = 43001
	INVALID_ASSET       int64 = 43002
	INVALID_BLOCK       int64 = 43003

	UNKNOWN_TRANSACTION int64 = 44001
	UNKNOWN_ASSET       int64 = 44002
	UNKNOWN_BLOCK       int64 = 44003
	UNKNOWN_CONTRACT    int64 = 44004

	INTERNAL_ERROR  int64 = 45001
	SMARTCODE_ERROR int64 = 47001
	PRE_EXEC_ERROR  int64 = 47002
)

var ErrMap = map[int64]string{
	SUCCESS:            "SUCCESS",
	SESSION_EXPIRED:    "SESSION EXPIRED",
	SERVICE_CEILING:    "SERVICE CEILING",
	ILLEGAL_DATAFORMAT: "ILLEGAL DATAFORMAT",
	INVALID_VERSION:    "INVALID VERSION",

	INVALID_METHOD: "INVALID METHOD",
	INVALID_PARAMS: "INVALID PARAMS",

	INVALID_TRANSACTION: "INVALID TRANSACTION",
	INVALID_ASSET:       "INVALID ASSET",
	INVALID_BLOCK:       "INVALID BLOCK",

	UNKNOWN_TRANSACTION: "UNKNOWN TRANSACTION",
	UNKNOWN_ASSET:       "UNKNOWN ASSET",
	UNKNOWN_BLOCK:       "UNKNOWN BLOCK",
	UNKNOWN_CONTRACT:    "UNKNOWN CONTRACT",

	INTERNAL_ERROR:                           "INTERNAL ERROR",
	SMARTCODE_ERROR:                          "SMARTCODE EXEC ERROR",
	PRE_EXEC_ERROR:                           "SMARTCODE PREPARE EXEC ERROR",
	int64(onyxErrors.ErrNoCode):               "INTERNAL ERROR, ErrNoCode",
	int64(onyxErrors.ErrUnknown):              "INTERNAL ERROR, ErrUnknown",
	int64(onyxErrors.ErrDuplicatedTx):         "INTERNAL ERROR, ErrDuplicatedTx",
	int64(onyxErrors.ErrDuplicateInput):       "INTERNAL ERROR, ErrDuplicateInput",
	int64(onyxErrors.ErrAssetPrecision):       "INTERNAL ERROR, ErrAssetPrecision",
	int64(onyxErrors.ErrTransactionBalance):   "INTERNAL ERROR, ErrTransactionBalance",
	int64(onyxErrors.ErrAttributeProgram):     "INTERNAL ERROR, ErrAttributeProgram",
	int64(onyxErrors.ErrTransactionContracts): "INTERNAL ERROR, ErrTransactionContracts",
	int64(onyxErrors.ErrTransactionPayload):   "INTERNAL ERROR, ErrTransactionPayload",
	int64(onyxErrors.ErrDoubleSpend):          "INTERNAL ERROR, ErrDoubleSpend",
	int64(onyxErrors.ErrTxHashDuplicate):      "INTERNAL ERROR, ErrTxHashDuplicate",
	int64(onyxErrors.ErrStateUpdaterVaild):    "INTERNAL ERROR, ErrStateUpdaterVaild",
	int64(onyxErrors.ErrSummaryAsset):         "INTERNAL ERROR, ErrSummaryAsset",
	int64(onyxErrors.ErrXmitFail):             "INTERNAL ERROR, ErrXmitFail",
	int64(onyxErrors.ErrNoAccount):            "INTERNAL ERROR, ErrNoAccount",
}