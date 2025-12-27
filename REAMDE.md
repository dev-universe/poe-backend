# PoE Backend (Golang) - Proof of Existence

A minimal blockchain-backend portfolio project:
- Create a SHA-256 hash from data (later: uploaded files)
- Record the hash on-chain via a Solidity contract
- Read the on-chain record (recorder address + timestamp)

This repo focuses on the **Golang backend** side:
transaction signing/sending, waiting for mining, and reading contract state.

## Architecture

- On-chain (Solidity):
  - Stores only `bytes32` hash + metadata (recorder, timestamp)
  - Emits an event when a hash is recorded

- Off-chain (Golang):
  - Generates hash
  - Sends `recordHash(bytes32)` transaction
  - Reads `getRecord(bytes32)` from the contract

## Prerequisites

- WSL2 (Ubuntu)
- Go
- Anvil + Foundry (contract side)
- A deployed PoE contract on Anvil (local)

## Environment Variables

Set these in your WSL shell:

- `RPC_URL` (e.g. `http://127.0.0.1:8545`)
- `POE_ADDRESS` (deployed contract address)
- `DEPLOYER_PRIVATE_KEY` (anvil account private key)

Example:

```bash
export RPC_URL=http://127.0.0.1:8545
export POE_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3
export DEPLOYER_PRIVATE_KEY=0x...
```

## Run (CLI)
```bash
go run ./cmd/poe-backend
```
Expected output:
- input string
- SHA-256 hash
- tx hash
- mined block + status
- recorder + timestamp

## Notes
- This project intentionally stores only hashes on-chain (not the raw files) to keep on-chain data minimal and immutable.
- Next steps (planned):
  - Gin API: POST /v1/proofs (file upload -> SHA-256 -> recordHash)
  - Idempotency: return 409 if already recorded
  - Event indexing + Postgres for fast queries