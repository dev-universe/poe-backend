package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	poe "github.com/dev-universe/poe-backend/internal/poe"
)

func mustEnv(key string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}

func sha256Bytes32(input string) [32]byte {
	return sha256.Sum256([]byte(input))
}

func main() {
	ctx := context.Background()

	// ---- env ----
	rpcURL := mustEnv("RPC_URL")
	contractAddr := common.HexToAddress(mustEnv("POE_ADDRESS"))
	privateKeyHex := mustEnv("DEPLOYER_PRIVATE_KEY")

	// ---- connect RPC ----
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("rpc dial failed: %v", err)
	}
	defer client.Close()

	// ---- private key ----
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("failed to create transactor: %v", err)
	}

	// ---- gas (EIP-1559 best-effort) ----
	if tip, err := client.SuggestGasTipCap(ctx); err == nil {
		auth.GasTipCap = tip
	}
	if fee, err := client.SuggestGasPrice(ctx); err == nil {
		auth.GasFeeCap = new(big.Int).Mul(fee, big.NewInt(2))
	}

	// ---- contract ----
	contract, err := poe.NewProofOfExistence(contractAddr, client)
	if err != nil {
		log.Fatalf("failed to bind contract: %v", err)
	}

	// ---- input -> hash ----
	input := "hello world (from go)"
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	}

	hash := sha256Bytes32(input)

	fmt.Println("input   :", input)
	fmt.Println("sha256  :", hex.EncodeToString(hash[:]))

	// ---- send tx ----
	tx, err := contract.RecordHash(auth, hash)
	if err != nil {
		// "Already recorded"는 정상적인 도메인 상태이므로 종료하지 않고 안내만 출력
		if strings.Contains(err.Error(), "Already recorded") ||
			strings.Contains(err.Error(), "execution reverted") {
			fmt.Println("already recorded on-chain; skipping tx")
		} else {
			log.Fatalf("recordHash tx failed: %v", err)
		}
	} else {
		fmt.Println("tx hash :", tx.Hash().Hex())

		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			log.Fatalf("tx mining failed: %v", err)
		}
		fmt.Printf("mined   : block=%d status=%d\n",
			receipt.BlockNumber.Uint64(),
			receipt.Status,
		)
	}

	// ---- read ----
	rec, err := contract.GetRecord(&bind.CallOpts{Context: ctx}, hash)
	if err != nil {
		log.Fatalf("getRecord failed: %v", err)
	}

	fmt.Println("recorder :", rec.Recorder.Hex())
	fmt.Println("timestamp:", rec.Timestamp.Uint64(), "=>", time.Unix(int64(rec.Timestamp.Uint64()), 0).UTC())

}
