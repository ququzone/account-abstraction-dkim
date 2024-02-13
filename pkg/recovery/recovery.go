package recovery

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var (
	ErrSubject = errors.New("error email subject")
	ErrChainId = errors.New("error chainId")
)

type Recoveror struct {
	client     *ethclient.Client
	transactor *bind.TransactOpts
}

type Recovery struct {
	recoveror map[uint64]*Recoveror
}

func NewRecovery(key, rpc string) (*Recovery, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, err
	}

	rpcs := strings.Split(rpc, ",")
	recoveror := make(map[uint64]*Recoveror, len(rpcs))
	for _, v := range rpcs {
		client, err := ethclient.Dial(v)
		if err != nil {
			return nil, err
		}

		chainId, err := client.ChainID(context.Background())
		if err != nil {
			return nil, err
		}

		transactor, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
		if err != nil {
			return nil, err
		}

		recoveror[chainId.Uint64()] = &Recoveror{
			client,
			transactor,
		}
	}

	return &Recovery{
		recoveror,
	}, nil
}

func (r *Recovery) PendingRecover(server, subject string, data, signature []byte) (uint64, string, string, error) {
	if !strings.Contains(subject, "0x") {
		return 0, "", "", ErrSubject
	}
	chainIdStr := subject[2:strings.Index(subject, "0x")]
	chainId, err := strconv.ParseUint(chainIdStr, 10, 64)
	if err != nil {
		return 0, "", "", ErrChainId
	}
	prefix := fmt.Sprintf("01%d", chainId)

	if !strings.HasPrefix(subject, prefix) || len(subject) <= len(prefix)+42 {
		return 0, "", "", ErrSubject
	}

	recoveror, ok := r.recoveror[chainId]
	if !ok {
		return 0, "", "", ErrChainId
	}
	accountAddr := subject[len(prefix) : len(prefix)+42]
	pubkey := subject[len(prefix)+42:]
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return 0, "", "", err
	}
	account, err := NewAccount(common.HexToAddress(accountAddr), recoveror.client)
	if err != nil {
		return 0, "", "", err
	}

	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(server))
	var serverBytes [32]byte
	copy(serverBytes[:32], sha.Sum(nil)[:])
	log.Printf("recovery server: %s with bytes: %s\n", server, hex.EncodeToString(serverBytes[:]))

	tx, err := account.PendingRecovery(recoveror.transactor, serverBytes, data, signature, pubkeyBytes)
	if err != nil {
		return 0, "", "", err
	}

	return chainId, accountAddr, tx.Hash().String(), nil
}

func (r *Recovery) Recover(chainId uint64, accountAddr string) (string, error) {
	recoveror, ok := r.recoveror[chainId]
	if !ok {
		return "", ErrChainId
	}

	account, err := NewAccount(common.HexToAddress(accountAddr), recoveror.client)
	if err != nil {
		return "", err
	}

	tx, err := account.Recovery(recoveror.transactor)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}
