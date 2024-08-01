package goar

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/permadao/goar/schema"
	"github.com/permadao/goar/utils"
)

type Wallet struct {
	Client *Client
	Signer *Signer
}

// proxyUrl: option
func NewWalletFromPath(path string, clientUrl string, proxyUrl ...string) (*Wallet, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewWallet(b, clientUrl, proxyUrl...)
}

func NewWallet(b []byte, clientUrl string, proxyUrl ...string) (w *Wallet, err error) {
	signer, err := NewSigner(b)
	if err != nil {
		return nil, err
	}

	w = &Wallet{
		Client: NewClient(clientUrl, proxyUrl...),
		Signer: signer,
	}

	return
}

func (w *Wallet) Owner() string {
	return w.Signer.Owner()
}

func (w *Wallet) SendAR(amount *big.Float, target string, tags []schema.Tag) (schema.Transaction, error) {
	return w.SendWinstonSpeedUp(utils.ARToWinston(amount), target, tags, 0)
}

func (w *Wallet) SendARSpeedUp(amount *big.Float, target string, tags []schema.Tag, speedFactor int64) (schema.Transaction, error) {
	return w.SendWinstonSpeedUp(utils.ARToWinston(amount), target, tags, speedFactor)
}

func (w *Wallet) SendWinston(amount *big.Int, target string, tags []schema.Tag) (schema.Transaction, error) {
	return w.SendWinstonSpeedUp(amount, target, tags, 0)
}

func (w *Wallet) SendWinstonSpeedUp(amount *big.Int, target string, tags []schema.Tag, speedFactor int64) (schema.Transaction, error) {
	reward, err := w.Client.GetTransactionPrice(0, &target)
	if err != nil {
		return schema.Transaction{}, err
	}

	tx := &schema.Transaction{
		Format:   2,
		Target:   target,
		Quantity: amount.String(),
		Tags:     utils.TagsEncode(tags),
		Data:     "",
		DataSize: "0",
		Reward:   fmt.Sprintf("%d", reward*(100+speedFactor)/100),
	}

	return w.SendTransaction(tx)
}

func (w *Wallet) SendData(data []byte, tags []schema.Tag) (schema.Transaction, error) {
	return w.SendDataSpeedUp(data, tags, 0)
}

func (w *Wallet) SendDataStream(data *os.File, tags []schema.Tag) (schema.Transaction, error) {
	return w.SendDataStreamSpeedUp(data, tags, 0)
}

// SendDataSpeedUp set speedFactor for speed up
// eg: speedFactor = 10, reward = 1.1 * reward
func (w *Wallet) SendDataSpeedUp(data []byte, tags []schema.Tag, speedFactor int64) (schema.Transaction, error) {
	reward, err := w.Client.GetTransactionPrice(len(data), nil)
	if err != nil {
		return schema.Transaction{}, err
	}

	tx := &schema.Transaction{
		Format:   2,
		Target:   "",
		Quantity: "0",
		Tags:     utils.TagsEncode(tags),
		Data:     utils.Base64Encode(data),
		DataSize: fmt.Sprintf("%d", len(data)),
		Reward:   fmt.Sprintf("%d", reward*(100+speedFactor)/100),
	}

	return w.SendTransaction(tx)
}

func (w *Wallet) SendDataStreamSpeedUp(data *os.File, tags []schema.Tag, speedFactor int64) (schema.Transaction, error) {
	fileInfo, err := data.Stat()
	if err != nil {
		return schema.Transaction{}, err
	}
	reward, err := w.Client.GetTransactionPrice(int(fileInfo.Size()), nil)
	if err != nil {
		return schema.Transaction{}, err
	}

	tx := &schema.Transaction{
		Format:     2,
		Target:     "",
		Quantity:   "0",
		Tags:       utils.TagsEncode(tags),
		Data:       "",
		DataReader: data,
		DataSize:   fmt.Sprintf("%d", fileInfo.Size()),
		Reward:     fmt.Sprintf("%d", reward*(100+speedFactor)/100),
	}

	return w.SendTransaction(tx)
}

func (w *Wallet) SendDataConcurrentSpeedUp(ctx context.Context, concurrentNum int, data interface{}, tags []schema.Tag, speedFactor int64) (schema.Transaction, error) {
	var reward int64
	var dataLen int
	isByteArr := true
	if _, isByteArr = data.([]byte); isByteArr {
		dataLen = len(data.([]byte))
	} else {
		fileInfo, err := data.(*os.File).Stat()
		if err != nil {
			return schema.Transaction{}, err
		}
		dataLen = int(fileInfo.Size())
	}
	reward, err := w.Client.GetTransactionPrice(dataLen, nil)
	if err != nil {
		return schema.Transaction{}, err
	}

	tx := &schema.Transaction{
		Format:   2,
		Target:   "",
		Quantity: "0",
		Tags:     utils.TagsEncode(tags),
		DataSize: fmt.Sprintf("%d", dataLen),
		Reward:   fmt.Sprintf("%d", reward*(100+speedFactor)/100),
	}

	if isByteArr {
		tx.Data = utils.Base64Encode(data.([]byte))
	} else {
		tx.DataReader = data.(*os.File)
	}

	return w.SendTransactionConcurrent(ctx, concurrentNum, tx)
}

// SendTransaction: if send success, should return pending
func (w *Wallet) SendTransaction(tx *schema.Transaction) (schema.Transaction, error) {
	uploader, err := w.getUploader(tx)
	if err != nil {
		return schema.Transaction{}, err
	}
	err = uploader.Once()
	return *tx, err
}

func (w *Wallet) SendTransactionConcurrent(ctx context.Context, concurrentNum int, tx *schema.Transaction) (schema.Transaction, error) {
	uploader, err := w.getUploader(tx)
	if err != nil {
		return schema.Transaction{}, err
	}
	err = uploader.ConcurrentOnce(ctx, concurrentNum)
	return *tx, err
}

func (w *Wallet) getUploader(tx *schema.Transaction) (*TransactionUploader, error) {
	anchor, err := w.Client.GetTransactionAnchor()
	if err != nil {
		return nil, err
	}
	tx.LastTx = anchor
	tx.Owner = w.Owner()
	if err = w.Signer.SignTx(tx); err != nil {
		return nil, err
	}
	return CreateUploader(w.Client, tx, nil)
}
