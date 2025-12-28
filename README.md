# PoE Backend (Golang) – Proof of Existence

A minimal **blockchain backend portfolio project** that demonstrates how a Golang service interacts with an Ethereum smart contract.

## What this project does

* Generate a **SHA-256 hash** from uploaded data
* Record the hash **on-chain** via a Solidity smart contract
* Retrieve on-chain record information (recorder address + timestamp)
* Enforce **idempotency**: the same hash cannot be recorded twice

This repository focuses on the **Golang backend side**:
transaction signing, sending, waiting for mining, and reading smart contract state.

---

## Architecture

### On-chain (Solidity)

* Stores only a `bytes32` hash and minimal metadata:

  * recorder address
  * timestamp
* Emits an event when a hash is recorded
* Prevents duplicate records at the contract level

### Off-chain (Golang)

* Accepts file uploads via HTTP (Gin)
* Computes SHA-256 hash
* Sends `recordHash(bytes32)` transactions
* Maps contract reverts to HTTP semantics (e.g. `409 Conflict`)
* Reads on-chain data via `getRecord(bytes32)`

---

## Prerequisites

* WSL2 (Ubuntu)
* Go
* Anvil + Foundry (for local blockchain)
* A deployed PoE contract on Anvil

---

## Environment Variables

Set these in your WSL shell:

* `RPC_URL` – Ethereum RPC endpoint
  (e.g. `http://127.0.0.1:8545`)
* `POE_ADDRESS` – deployed PoE contract address
* `DEPLOYER_PRIVATE_KEY` – private key of the Anvil account
* `LISTEN_ADDR` (optional) – API listen address (default `:8080`)

Example:

```bash
export RPC_URL=http://127.0.0.1:8545
export POE_ADDRESS=0x5F~~~~
export DEPLOYER_PRIVATE_KEY=0x...
export LISTEN_ADDR=:8080
```

---

## Run

Start the API server:

```bash
go run ./cmd/poe-backend
```

Expected behavior:

* The server listens on `:8080`
* Incoming requests trigger on-chain transactions or reads

---

## API Usage

### POST /v1/proofs

Upload a file and record its SHA-256 hash on-chain.

#### Request

```bash
echo "hello api" > /tmp/hello.txt
curl -i -F "file=@/tmp/hello.txt" http://localhost:8080/v1/proofs
```

#### Success Response (first upload)

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "hash": "7c6a~~~~",
  "tx_hash": "0x0a~~~~",
  "already_recorded": false
}
```

#### Duplicate Upload (idempotent)

```bash
curl -i -F "file=@/tmp/hello.txt" http://localhost:8080/v1/proofs
```

```http
HTTP/1.1 409 Conflict
Content-Type: application/json

{
  "hash": "7c6a~~~~",
  "tx_hash": "",
  "already_recorded": true
}
```

---

### GET /v1/proofs/:hash

Retrieve on-chain record information for a given hash.

```bash
curl -i http://localhost:8080/v1/proofs/7c6a~~~~
```

#### Success Response

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "hash": "7c6a~~~",
  "recorder": "0xf3~~",
  "timestamp": 1766828014
}
```

---

## Notes

* Only **hashes** are stored on-chain; raw files are never written to the blockchain.
* Idempotency is enforced at the **smart contract level**, not just in the backend.
* HTTP semantics reflect blockchain state:

  * `201 Created` → new on-chain record
  * `409 Conflict` → hash already recorded
* This project is designed to be fully reproducible on a local Anvil chain.

---

## Future Improvements

* Asynchronous transaction handling (respond before mining)
* Event indexing + PostgreSQL for fast queries
* Authentication and rate limiting
* Support for public testnets (Sepolia)