# goar

### Install

```
go get github.com/everFinance/goar
```

### Example

#### Send AR or Winston

```golang
package main

import (
	"fmt"
	"math/big"
	"github.com/permadao/goar/schema"
	"github.com/permadao/goar"
)

func main() {
	wallet, err := goar.NewWalletFromPath("./test-keyfile.json", "https://arweave.net")
	if err != nil {
		panic(err)
	}

	tx, err := wallet.SendAR(
  //id, err := wallet.SendWinston( 
		big.NewFloat(1.0), // AR amount
		{{target}}, // target address
		[]schema.Tag{},
	)

	fmt.Println(tx.ID, err)
}

```

#### Send Data

```golang
tx, err := wallet.SendData(
  []byte("123"), // Data bytes
  []schema.Tag{
    schema.Tag{
      Name:  "testSendData",
      Value: "123",
    },
  },
)

fmt.Println(id, err) // {{id}}, nil
```

#### Send Data SpeedUp

Arweave occasionally experiences congestion, and a low Reward can cause a transaction to fail; use speedUp to accelerate the transaction.

```golang
speedUp := int64(50) // means reward = reward * 150%
tx, err := wallet.SendDataSpeedUp(
  []byte("123"), // Data bytes
  []schema.Tag{
    schema.Tag{
      Name:  "testSendDataSpeedUp",
      Value: "123",
    },
  },speedUp)

fmt.Println(tx.ID, err)
```
### Components

#### Client

- [x] GetInfo
- [x] GetTransactionByID
- [x] GetTransactionStatus
- [x] GetTransactionField
- [x] GetTransactionData
- [x] GetTransactionPrice
- [x] GetTransactionAnchor
- [x] SubmitTransaction
- [x] Arql(Deprecated)
- [x] GraphQL
- [x] GetWalletBalance
- [x] GetLastTransactionID
- [x] GetBlockByID
- [x] GetBlockByHeight
- [x] BatchSendItemToBundler
- [x] GetBundle
- [x] GetTxDataFromPeers
- [x] BroadcastData
- [x] GetUnconfirmedTx
- [x] GetPendingTxIds
- [x] GetBlockHashList
- [x] ConcurrentDownloadChunkData

Initialize the instance:

```golang
arClient := goar.NewClient("https://arweave.net")

// if your network is not good, you can config http proxy
proxyUrl := "http://127.0.0.1:8001"
arClient := goar.NewClient("https://arweave.net", proxyUrl)
```

#### Wallet

- [x] SendAR
- [x] SendARSpeedUp
- [x] SendWinston
- [x] SendWinstonSpeedUp
- [x] SendData
- [x] SendDataSpeedUp
- [x] SendTransaction
- [x] CreateAndSignBundleItem
- [x] SendBundleTxSpeedUp
- [x] SendBundleTx
- [x] SendPst

Initialize the instance, use a keyfile.json:

```golang
arWallet := goar.NewWalletFromPath("./keyfile.json")

// if your network is not good, you can config http proxy
proxyUrl := "http://127.0.0.1:8001"
arWallet := NewWalletFromPath("./keyfile.json", "https://arweave.net", proxyUrl)
```

#### Signer

- [x] SignTx
- [x] SignMsg
- [x] Owner

```golang
signer := goar.NewSignerFromPath("./keyfile.json")
```

#### Utils

Package for Arweave develop toolkit.

- [x] Base64Encode
- [x] Base64Decode
- [x] Sign
- [x] Verify
- [x] DeepHash
- [x] GenerateChunks
- [x] ValidatePath
- [x] OwnerToAddress
- [x] OwnerToPubKey
- [x] TagsEncode
- [x] TagsDecode
- [x] PrepareChunks
- [x] GetChunk
- [x] SignTransaction
- [x] GetSignatureData
- [x] VerifyTransaction
- [x] NewBundle
- [x] NewBundleItem
- [x] SubmitItemToBundlr
- [x] SubmitItemToArSeed

### Development

#### Test

```
make test
```
---
### About chunks
1. First, we use Chunk transactions for all types of transactions in this library, so we only support transactions where format equals 2.
2. Second, the library already encapsulates a common interface for sending transactions : e.g `SendAR; SendData`. The user only needs to call this interface to send the transaction and do not need to worry about the usage of chunks.
3. The third，If the user needs to control the transaction such as breakpoint retransmission and breakpoint continuation operations. Here is how to do it.

#### chunked uploading advanced options
##### upload all transaction data
The method of submitting a data transaction is to use chunk uploading. This method will allow larger transaction sizes, resuming a transaction upload if it's interrupted and give progress updates while uploading.
Simple example:

```golang
arNode := "https://arweave.net"
w, err := goar.NewWalletFromPath("../example/testKey.json", arNode) // your wallet private key
anchor, err := w.Client.GetTransactionAnchor()
if err != nil {
  return
}
data, err := ioutil.ReadFile("./2.3MBPhoto.jpg")
if err != nil {
  return
}

reward, err := w.Client.GetTransactionPrice(data, nil)
if err != nil {
  return
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

tx.LastTx = anchor
tx.Owner = utils.Base64Encode(w.PubKey.N.Bytes())

if err = utils.SignTransaction(tx, w.PubKey, w.Signer.PrvKey); err != nil {
  return
}

id = tx.ID

uploader, err := goar.CreateUploader(w.Client, tx, nil)
if err != nil {
  return
}

err = uploader.Once()
if err != nil {
  return
}
```

##### Breakpoint continuingly

You can resume an upload from a saved uploader object, that you have persisted in storage some using json.marshal(uploader) at any stage of the upload. To resume, parse it back into an object and pass it to getUploader() along with the transactions data:

```golang
uploaderBuf, err := ioutil.ReadFile("./jsonUploaderFile.json")
lastUploader := &txType.TransactionUploader{}
err = json.Unmarshal(uploaderBuf, lastUploader)
assert.NoError(t, err)

// new uploader object by last time uploader
newUploader, err := txType.CreateUploader(wallet.Client, lastUploader.FormatSerializedUploader(), bigData)
assert.NoError(t, err)
for !newUploader.IsComplete() {
  err := newUploader.UploadChunk()
  assert.NoError(t, err)
}
```

When resuming the upload, you must provide the same data as the original upload. When you serialize the uploader object with json.marshal() to save it somewhere, it will not include the data.

##### Breakpoint retransmission

You can also resume an upload from just the transaction ID and data, once it has been mined into a block. This can be useful if you didn't save the uploader somewhere but the upload got interrupted. This will re-upload all of the data from the beginning, since we don't know which parts have been uploaded:

```golang
bigData, err := ioutil.ReadFile(filePath)
txId := "myTxId"

// get uploader by txId and post big data by chunks
uploader, err := goar.CreateUploader(wallet.Client, txId, bigData)
assert.NoError(t, err)
assert.NoError(t, uploader.Once())
```

---
### About Arweave Bundles
1. `goar` implemented creating, editing, reading and verifying bundles tx
2. This is the [ANS-104](https://github.com/joshbenaron/arweave-standards/blob/ans104/ans/ANS-104.md) standard protocol and refers to the [arbundles](https://github.com/Bundler-Network/arbundles) js-lib implement

#### Create Bundle Item
```go
signer, err := goar.NewSignerFromPath("./testKey.json") // rsa signer
// or 
signer, err := goether.NewSigner("0x.....") // ecdsa signer

bundler, err := goar.NewBundler(signer)

// Create Item
data := []byte("aa bb cc dd")
target := "" // option 
anchor := "" // option
tags := []schema.Tags{}{} // option bundle item tags
item01, err := bundler.CreateAndSignItem(data, target, anchor, tags)    
// Same as create item
item02
item03
....

```
#### assemble bundle and send to arweave network 
You can send items directly to the arweave network
```go

items := []schema.BundleItem{item01, item02, item03 ...}
bundle, err := utils.NewBundle(items...)

w, err := goar.NewWalletFromPath("./key.json", arNode)

arTxTags := []schema.Tags{}{} // option
tx, err := w.SendBundleTx(bd.BundleBinary, arTxtags)

```

#### Verify Bundle Items
```go

// verify
for _, item := range bundle.Items {
  err = utils.VerifyBundleItem(item)
  assert.NoError(t, err)
}
```
