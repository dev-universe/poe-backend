package app

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/dev-universe/poe-backend/internal/poe"
	"github.com/dev-universe/poe-backend/internal/service"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	RPCURL     string
	POEAddress string
	PrivateKey string
	ListenAddr string
	MaxUpload  int64
}

func getEnv(key, def string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	return v
}

func LoadConfig() Config {
	return Config{
		RPCURL:     getEnv("RPC_URL", "http://127.0.0.1:8545"),
		POEAddress: getEnv("POE_ADDRESS", ""),
		PrivateKey: getEnv("DEPLOYER_PRIVATE_KEY", ""),
		ListenAddr: getEnv("LISTEN_ADDR", ":8080"),
		MaxUpload:  5 << 20, // 5MB
	}
}

type Built struct {
	Cfg     Config
	Service *service.PoEService
}

func Build(ctx context.Context, cfg Config) (*Built, error) {
	if cfg.POEAddress == "" {
		return nil, fmt.Errorf("POE_ADDRESS is required")
	}
	if cfg.PrivateKey == "" {
		return nil, fmt.Errorf("DEPLOYER_PRIVATE_KEY is required")
	}

	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	pkHex := strings.TrimPrefix(cfg.PrivateKey, "0x")
	pk, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, err
	}

	// best-effort fees
	if tip, err := client.SuggestGasTipCap(ctx); err == nil {
		auth.GasTipCap = tip
	}
	if gp, err := client.SuggestGasPrice(ctx); err == nil {
		auth.GasFeeCap = new(big.Int).Mul(gp, big.NewInt(2))
	}

	contract, err := poe.NewProofOfExistence(common.HexToAddress(cfg.POEAddress), client)
	if err != nil {
		return nil, err
	}

	svc := service.NewPoEService(contract, client, auth)
	log.Printf("configured: rpc=%s contract=%s listen=%s", cfg.RPCURL, cfg.POEAddress, cfg.ListenAddr)

	return &Built{Cfg: cfg, Service: svc}, nil
}
