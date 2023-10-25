package client

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type client struct {
	cfg    Config
	client *ethclient.Client
}

func New(opts ...Option) (*client, error) {
	cfg := defaultConfig()

	for _, opt := range opts {
		opt(&cfg)
	}

	c, err := ethclient.Dial(cfg.endpoint)
	if err != nil {
		return nil, err
	}

	return &client{
		cfg:    cfg,
		client: c,
	}, nil
}

// InvokeContract Invoke an EVM smart contract
func (c *client) InvokeContract(nonce uint64, invokeFunc func(opts *bind.TransactOpts) (*types.Transaction, error)) ([]byte, error) {
	networkID, err := c.chainID(context.TODO())
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(c.cfg.privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	if nonce == 0 {
		nonce, err = c.Nonce()
		if err != nil {
			return nil, err
		}
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	signer := types.LatestSignerForChainID(big.NewInt(networkID))
	opts := &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, privateKey)
		},
		From:    fromAddress,
		Context: context.Background(),
		Nonce:   big.NewInt(int64(nonce)),
	}

	tx, err := invokeFunc(opts)
	if err != nil {
		return nil, err
	}

	return tx.MarshalJSON()
}

func (c *client) EthClient() *ethclient.Client {
	return c.client
}

func (c *client) Address() (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(c.cfg.privateKey)
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func (c *client) chainID(ctx context.Context) (int64, error) {
	chainID, err := c.client.ChainID(ctx)
	if err != nil {
		return 0, err
	}

	return chainID.Int64(), nil
}

func (c *client) Nonce() (uint64, error) {
	addr, err := c.Address()
	if err != nil {
		return 0, err
	}

	nonce, err := c.client.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}
