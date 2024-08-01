package goar

import (
	"encoding/base64"

	"github.com/stretchr/testify/assert"

	"testing"
)

var testWallet *Wallet
var err error

func init() {
	clientUrl := "https://arweave.net"
	testWallet, err = NewWalletFromPath("testKey.json", clientUrl)

	if err != nil {
		panic(err)
	}
}

func TestPubKey(t *testing.T) {
	pubKey := testWallet.Signer.PubKey
	assert.Equal(t, "nQ9iy1fRM2xrgggjHhN1xZUnOkm9B4KFsJzH70v7uLMVyDqfyIJEVXeJ4Jhk_8KpjzYQ1kYfnCMjeXnhTUfY3PbeqY4PsK5nTje0uoOe1XGogeGAyKr6mVtKPhBku-aq1gz7LLRHndO2tvLRbLwX1931vNk94bSfJPYgMfU7OXxFXbTdKU38W6u9ShoaJGgUQI1GObd_sid1UVniCmu7P-99XPkixqyacsrkHzBajGz1S7jGmpQR669KWE9Z0unvH0KSHxAKoDD7Q7QZO7_4ujTBaIFwy_SJUxzVV8G33xvs7edmRdiqMdVK5W0LED9gbS4dv_aee9IxUJQqulSqZphPgShIiGNl9TcL5iUi9gc9cXR7ISyavos6VGiem_A-S-5f-_OKxoeZzvgAQda8sD6jtBTTuM5eLvgAbosbaSi7zFYCN7zeFdB72OfvCh72ZWSpBMH3dkdxsKCDmXUXvPdDLEnnRS87-MP5RV9Z6foq_YSEN5MFTMDdo4CpFGYl6mWTP6wUP8oM3Mpz3-_HotwSZEjASvWtiff2tc1fDHulVMYIutd52Fis_FKj6K1fzpiDYVA1W3cV4P28Q1-uF3CZ8nJEa5FXchB9lFrXB4HvsJVG6LPSt-y2R9parGi1_kEc6vOYIesKspgZ0hLyIKtqpTQFiPgKRlyUc-WEn5E", base64.RawURLEncoding.EncodeToString(pubKey.N.Bytes()))
}

func TestAddress(t *testing.T) {
	addr := testWallet.Signer.Address
	assert.Equal(t, "eIgnDk4vSKPe0lYB6yhCHDV1dOw3JgYHGocfj7WGrjQ", addr)
}

// test sand ar without data
func TestWallet_SendAR(t *testing.T) {
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("./example/testKey.json", arNode) // your wallet private key
	// assert.NoError(t, err)
	//
	// target := "cSYOy8-p1QFenktkDBFyRM3cwZSTrQ_J4EsELLho_UE"
	// amount := big.NewFloat(0.001)
	// tags := []schema.Tag{
	// 	{Name: "GOAR", Value: "sendAR"},
	// }
	// tx,  err := w.SendAR(amount, target, tags)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s \n", tx.ID)
}

// test send small size file
func TestWallet_SendDataSpeedUp01(t *testing.T) {
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("./example/testKey.json", arNode) // your wallet private key
	// assert.NoError(t, err)
	//
	// // data := []byte("aaa this is a goar test small size file data") // small file
	// data := make([]byte, 255*1024)
	// for i := 0; i < len(data); i++ {
	// 	data[i] = byte('b' + i)
	// }
	// tags := []schema.Tag{
	// 	{Name: "GOAR", Value: "SMDT"},
	// }
	// tx, err := w.SendDataSpeedUp(data, tags, 50)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s", tx.ID)
}

// test send big size file
func TestWallet_SendDataSpeedUp02(t *testing.T) {
	// proxyUrl := "http://127.0.0.1:8001"
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("./wallet/account1.json", arNode, proxyUrl) // your wallet private key
	// assert.NoError(t, err)
	//
	// data, err := os.ReadFile("/Users/sandyzhou/Downloads/abc.jpeg")
	// if err != nil {
	// 	panic(err)
	// }
	// tags := []schema.Tag{
	// 	{Name: "Sender", Value: "Jie"},
	// 	{Name: "Data-Introduce", Value: "Happy anniversary, my google and dearest! I‘m so grateful to have you in my life. I love you to infinity and beyond! (⁎⁍̴̛ᴗ⁍̴̛⁎)"},
	// }
	// tx, err := w.SendDataSpeedUp(data, tags, 10)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s", tx.ID)
}

func Test_SendPstTransfer(t *testing.T) {
	// w, err := NewWalletFromPath("./wallet/account1.json","https://arweave.net")
	// assert.NoError(t, err)
	//
	// contractId := "usjm4PCxUd5mtaon7zc97-dt-3qf67yPyqgzLnLqk5A"
	// target := "Ii5wAMlLNz13n26nYY45mcZErwZLjICmYd46GZvn4ck"
	// qty := big.NewInt(1)
	// arTx, err := w.SendPst(contractId,target,qty,nil,50)
	// assert.NoError(t, err)
	// t.Log(arTx.ID)
}

func TestCreateUploader(t *testing.T) {
	// w, err := NewWalletFromPath("./wallet/account1.json", "https://arweave.net")
	// assert.NoError(t, err)
	// t.Log(w.Signer.Address)

	// data, err := os.ReadFile("/Users/sandyzhou/Downloads/44444.mp4")
	// if err != nil {
	// 	panic(err)
	// }
	// tags := []schema.Tag{
	// 	{Name: "Content-Type", Value: "video/mpeg4"},
	// }
	// tx, err := w.SendData(data, tags)
	// assert.NoError(t, err)
	// t.Log(tx.ID)
}

func TestNewWallet(t *testing.T) {
	// cli := NewClient("https://arseed-dev.web3infra.dev")
	// data, err := cli.GetTransactionData("SAk5DdgYiKZTBFIxpVmiOQKsJdVbXr9qj5jTA5ACDmY")
	// assert.NoError(t, err)
	// os.WriteFile("/Users/sandyzhou/Downloads/55555.mp4", data, 0666)

	// -I6guxsTtnaLazLN4HgCHtYXNXQGALcHDrpi6oz7Cbk
	// 91s

	// _OXH-KzWyJVZsHEJ257uMcsi1VqdvMLroxVTeAYODi4
	// 408s

	// 355oQfW-6XM659iG4Yy69vyAeuTC5v7c_QqGU4iRI9Q
	// 48s

	// _bIcfrydLSGRzwLDYvwCB16SpKAeClJdZHLRC0lgOnQ
	// 24s
}

// func TestTransactionUploader_ConcurrentUploadChunks(t *testing.T) {
// 	w, err := NewWalletFromPath("./wallet/account1.json", "https://arweave.net")
// 	assert.NoError(t, err)
// 	t.Log(w.Signer.Address)
// 	signer01 := w.Signer
// 	// sig item01 by ecc signer
// 	itemSigner01, err := NewItemSigner(signer01)
// 	assert.NoError(t, err)
// 	d1, err := os.ReadFile("/Users/sandyzhou/Downloads/1.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item01, err := itemSigner01.CreateAndSignItem(d1, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item01", "id", item01.Id)

// 	d2, err := os.ReadFile("/Users/sandyzhou/Downloads/2.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item02, err := itemSigner01.CreateAndSignItem(d2, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item02", "id", item02.Id)

// 	d3, err := os.ReadFile("/Users/sandyzhou/Downloads/3.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item03, err := itemSigner01.CreateAndSignItem(d3, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item03", "id", item03.Id)

// 	d4, err := os.ReadFile("/Users/sandyzhou/Downloads/4.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item04, err := itemSigner01.CreateAndSignItem(d4, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)

// 	d5, err := os.ReadFile("/Users/sandyzhou/Downloads/5.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item05, err := itemSigner01.CreateAndSignItem(d5, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item05", "id", item05.Id)

// 	d6, err := os.ReadFile("/Users/sandyzhou/Downloads/6.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item06, err := itemSigner01.CreateAndSignItem(d6, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item06", "id", item06.Id)

// 	d7, err := os.ReadFile("/Users/sandyzhou/Downloads/7.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item07, err := itemSigner01.CreateAndSignItem(d7, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item07", "id", item07.Id)

// 	d8, err := os.ReadFile("/Users/sandyzhou/Downloads/8.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item08, err := itemSigner01.CreateAndSignItem(d8, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item08", "id", item08.Id)

// 	d9, err := os.ReadFile("/Users/sandyzhou/Downloads/9.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item09, err := itemSigner01.CreateAndSignItem(d9, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item09", "id", item09.Id)

// 	d10, err := os.ReadFile("/Users/sandyzhou/Downloads/10.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item10, err := itemSigner01.CreateAndSignItem(d10, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item10", "id", item10.Id)

// 	d11, err := os.ReadFile("/Users/sandyzhou/Downloads/11.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	item11, err := itemSigner01.CreateAndSignItem(d11, "", "", []schema.Tag{
// 		{Name: "Content-Type", Value: "image/jpeg"},
// 		{Name: "Owner", Value: "Vv"},
// 	})
// 	assert.NoError(t, err)
// 	t.Log("item11", "id", item11.Id)

// 	// assemble bundle
// 	bundle, err := utils.NewBundle(item01, item02, item03, item04, item05, item06, item07, item08, item09, item10, item11)
// 	assert.NoError(t, err)

// 	t.Log(len(bundle.BundleBinary))
// 	// send to arweave
// 	// ctx ,cancel := context.WithTimeout(context.Background(),100*time.Millisecond)
// 	// defer cancel()
// 	// tx, err := w.SendBundleTx(ctx, 0,bundle.BundleBinary, []schema.Tag{
// 	// 	{Name: "APP", Value: "Goar"},
// 	// 	{Name: "Protocol-Name", Value: "BAR"},
// 	// 	{Name: "Action", Value: "Burn"},
// 	// 	{Name: "App-Name", Value: "SmartWeaveAction"},
// 	// 	{Name: "App-Version", Value: "0.3.0"},
// 	// 	{Name: "Input", Value: `{"function":"mint"}`},
// 	// 	{Name: "Contract", Value: "VFr3Bk-uM-motpNNkkFg4lNW1BMmSfzqsVO551Ho4hA"},
// 	// })
// 	// assert.NoError(t, err)
// 	// t.Log(tx.ID)
// }
