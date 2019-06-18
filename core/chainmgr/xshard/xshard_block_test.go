/*
 * Copyright (C) 2019 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package xshard

import (
	"testing"

	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/config"
	crossshard "github.com/ontio/ontology/core/chainmgr/message"
	"github.com/ontio/ontology/core/ledger"
	"github.com/ontio/ontology/core/types"
)

func newTestShardMsg(t *testing.T) *types.CrossShardMsg {
	shardMsg := &types.CrossShardMsg{
		CrossShardMsgInfo: &types.CrossShardMsgInfo{
			FromShardID:   common.NewShardIDUnchecked(0),
			MsgHeight:     uint32(90),
			SignMsgHeight: uint32(100),
		},
	}
	return shardMsg
}

func TestCrossShardPool(t *testing.T) {
	InitCrossShardPool(common.NewShardIDUnchecked(1), 100)
	shardMsg := newTestShardMsg(t)
	acc1 := account.NewAccount("")
	tx, err := crossshard.NewCrossShardTxMsg(acc1, uint32(120), shardMsg.CrossShardMsgInfo.FromShardID, 0, 20000, nil)
	if err != nil {
		t.Errorf("crossShardPool NewCrossShardTxMsg err:%s", err)
		return
	}
	ldg, err := ledger.NewLedger(config.DEFAULT_DATA_DIR, 0)
	if err != nil {
		t.Errorf("failed to new ledger")
		return
	}
	if err = AddCrossShardInfo(ldg, shardMsg, tx); err != nil {
		t.Fatalf("failed add CrossShardInfo:%s", err)
	}
}