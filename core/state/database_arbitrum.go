package state

import (
	"errors"

	"github.com/paxosglobal/go-ethereum-arbitrum/common"
	"github.com/paxosglobal/go-ethereum-arbitrum/core/rawdb"
	"github.com/paxosglobal/go-ethereum-arbitrum/ethdb"
)

func (db *CachingDB) ActivatedAsm(target ethdb.WasmTarget, moduleHash common.Hash) ([]byte, error) {
	cacheKey := activatedAsmCacheKey{moduleHash, target}
	if asm, _ := db.activatedAsmCache.Get(cacheKey); len(asm) > 0 {
		return asm, nil
	}
	if asm := rawdb.ReadActivatedAsm(db.wasmdb, target, moduleHash); len(asm) > 0 {
		db.activatedAsmCache.Add(cacheKey, asm)
		return asm, nil
	}
	return nil, errors.New("not found")
}
