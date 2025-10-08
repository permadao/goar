package goar

import (
	"crypto/sha256"
	"github.com/permadao/goar/schema"
	"github.com/permadao/goar/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyEcdsaTx(t *testing.T) {
	// id := "hmtw7VXo-yfn_Gj5g5ZpcWU1P5ZMziU4EMhap1oPjyE"
	tx := &schema.Transaction{
		Format: 2,
		ID:     "",
		LastTx: "Pu9WXD63YpX4TQe3iwQKF8-bkbtPx7WpqHtUvBcuqO9Bvq44SyX6dTOMfOCERX1t",
		Owner:  "",
		Tags: []schema.Tag{
			{Name: "VGVzdA", Value: "ZWNkc2EtdHg"},
		},
		Target:     "Qa8AAZv-sEhQRIm7xZr3CVLtlzIH8NezaY0GZhURcAc",
		Quantity:   "1",
		Data:       "",
		DataReader: nil,
		DataSize:   "0",
		DataRoot:   "",
		Reward:     "3429726",
		Signature:  "ElI9De-2Z5AT3aoq5TYCz-qHJGa_tmYqqYwC7Rf89_05O5wVr2U-SuH9WPF-iRr1HQGjR-7_XVw0a2LyW0rf4gE",
	}

	sig, _ := utils.Base64Decode(tx.Signature)
	sigHash := sha256.Sum256(sig)
	txId := utils.Base64Encode(sigHash[:])
	assert.Equal(t, "hmtw7VXo-yfn_Gj5g5ZpcWU1P5ZMziU4EMhap1oPjyE", txId)
	owner, err := GetEcdsaTxOwner(tx)
	assert.NoError(t, err)
	assert.Equal(t, "AlFhxNH-6NmRDVEukOvvgsrEOWgxLi5x_h-r9Cg_dP27", owner)
	address, err := OwnerToAddress(owner)
	assert.NoError(t, err)
	assert.Equal(t, "RymI02hes920xGugzRJ3L54eGg-jVU-_R2uCI057_nU", address)
	/*
			{
		        "address": "RymI02hes920xGugzRJ3L54eGg-jVU-_R2uCI057_nU",
		        "key": "AlFhxNH-6NmRDVEukOvvgsrEOWgxLi5x_h-r9Cg_dP27"
		      }
	*/
	err = VerifyEcdsaTxSig(tx)
	assert.NoError(t, err)
}

func TestEcdsaArTx(t *testing.T) {
	private := "ccb43edab9f7fd5c24388c43579b9d50ad3c8d94cf5fce9bb0bcb8e6ddb57bce"
	ecSigner, err := NewEcSigner(private)
	assert.NoError(t, err)
	t.Log("address: ", ecSigner.Address())

	tx := &schema.Transaction{
		Format: 2,
		LastTx: "5Yx84D8oaCAlPqMe7vmDGuEwXtIOy_rMOHlpgm7yovYoW__Mj8S2gC8UuWmLs-61",
		Owner:  "",
		Tags: []schema.Tag{
			{Name: "VGVzdA", Value: "ZWNkc2EtdHg"},
		},
		Target:    "cSYOy8-p1QFenktkDBFyRM3cwZSTrQ_J4EsELLho_UE",
		Quantity:  "100000000",
		Data:      "",
		DataSize:  "0",
		DataRoot:  "",
		Reward:    "7050359",
		Signature: "",
	}
	err = ecSigner.SignTx(tx)
	assert.NoError(t, err)

	t.Log("txId: ", tx.ID)
	t.Log("txSig: ", tx.Signature)

	owner, err := GetEcdsaTxOwner(tx)
	assert.NoError(t, err)
	t.Log("owner: ", owner)
	address, err := OwnerToAddress(owner)
	assert.NoError(t, err)
	t.Log("address: ", address)

	err = VerifyEcdsaTxSig(tx)
	assert.NoError(t, err)

	// submit ecdsa arTx
	cli := NewClient("https://arweave.net")
	status, code, err := cli.SubmitTransaction(tx)
	t.Log("status: ", status)
	t.Log("code: ", code)
	assert.NoError(t, err)
}
