package recovery

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var (
	ErrSubject = errors.New("error email subject")
)

type Recovery struct {
	chainId    uint64
	client     *ethclient.Client
	transactor *bind.TransactOpts
}

func NewRecovery(keyFile, keyPassphrase, rpc string) (*Recovery, error) {
	keyData, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	keystore, err := keystore.DecryptKey(keyData, keyPassphrase)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(keystore.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	return &Recovery{
		chainId:    chainId.Uint64(),
		client:     client,
		transactor: transactor,
	}, nil
}

func (r *Recovery) Recover(server, subject string, data, signature []byte) (string, error) {
	prefix := fmt.Sprintf("01%d", r.chainId)
	if !strings.HasPrefix(subject, prefix) || len(subject) != len(prefix)+42+128 {
		return "", ErrSubject
	}
	accountAddr := subject[len(prefix) : len(prefix)+42]
	pubkey := subject[len(prefix)+42:]
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return "", err
	}
	account, err := NewAccount(common.HexToAddress(accountAddr), r.client)
	if err != nil {
		return "", err
	}

	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(server))
	var serverBytes [32]byte
	copy(serverBytes[:32], sha.Sum(nil)[:])
	log.Printf("recovery server: %s with bytes: %s\n", server, hex.EncodeToString(serverBytes[:]))

	tx, err := account.PendingRecovery(r.transactor, serverBytes, data, signature, pubkeyBytes)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}
