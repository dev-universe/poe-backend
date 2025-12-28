package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var ErrAlreadyRecorded = errors.New("already recorded")

type ProofResult struct {
	HashHex         string `json:"hash"`
	TxHash          string `json:"tx_hash"`
	AlreadyRecorded bool   `json:"already_recorded"`
}

type ProofContract interface {
	RecordHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error)
	GetRecord(opts *bind.CallOpts, hash [32]byte) (struct {
		Recorder  common.Address
		Timestamp *big.Int
	}, error)
}

type PoEService struct {
	Contract ProofContract
	Backend  bind.DeployBackend
	Auth     *bind.TransactOpts
}

func NewPoEService(contract ProofContract, backend bind.DeployBackend, auth *bind.TransactOpts) *PoEService {
	return &PoEService{Contract: contract, Backend: backend, Auth: auth}
}

func sha256Bytes32FromBytes(b []byte) ([32]byte, string) {
	sum := sha256.Sum256(b)
	return sum, hex.EncodeToString(sum[:])
}

func (s *PoEService) RecordBytes(ctx context.Context, data []byte) (*ProofResult, error) {
	h32, hHex := sha256Bytes32FromBytes(data)

	tx, err := s.Contract.RecordHash(s.Auth, h32)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "Already recorded") || strings.Contains(msg, "execution reverted") {
			return &ProofResult{
				HashHex:         hHex,
				TxHash:          "",
				AlreadyRecorded: true,
			}, ErrAlreadyRecorded
		}
		return nil, err
	}

	if _, err := bind.WaitMined(ctx, s.Backend, tx); err != nil {
		return nil, err
	}

	return &ProofResult{
		HashHex: hHex,
		TxHash:  tx.Hash().Hex(),
	}, nil
}

type OnchainRecord struct {
	Recorder  string `json:"recorder"`
	Timestamp uint64 `json:"timestamp"`
}

func (s *PoEService) GetRecord(ctx context.Context, hash32 [32]byte) (*OnchainRecord, error) {
	rec, err := s.Contract.GetRecord(&bind.CallOpts{Context: ctx}, hash32)
	if err != nil {
		return nil, err
	}
	return &OnchainRecord{
		Recorder:  rec.Recorder.Hex(),
		Timestamp: rec.Timestamp.Uint64(),
	}, nil
}
