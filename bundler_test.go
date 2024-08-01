package goar

import (
	"testing"

	"github.com/everFinance/goether"
	"github.com/permadao/goar/schema"
	"github.com/permadao/goar/utils"
	"github.com/stretchr/testify/assert"
)

var (
	signer01 *goether.Signer
	signer02 *Signer
)

func init() {
	var err error
	signer01, err = goether.NewSigner("1f534ac18009182c07d266fe4a7903c0bcc8a66190f0967b719b2b3974a69c2f")
	if err != nil {
		return
	}

	signer02, err = NewSignerFromPath("testKey.json")
	if err != nil {
		return
	}
}

func TestBundle(t *testing.T) {
	b1, err := NewBundler(signer01)
	assert.NoError(t, err)
	item1, err := b1.CreateAndSignItem([]byte("eth foo"), "", "", []schema.Tag{{Name: "Content-Type", Value: "application/txt"}})
	assert.NoError(t, err)

	b2, err := NewBundler(signer02)
	assert.NoError(t, err)
	item2, err := b2.CreateAndSignItem([]byte("ar foo"), "", "", []schema.Tag{{Name: "Content-Type", Value: "application/txt"}})
	assert.NoError(t, err)

	bundle, err := utils.NewBundle(item1, item2)
	assert.NoError(t, err)

	resBundle, err := utils.DecodeBundle(bundle.Binary)
	assert.NoError(t, err)

	for _, item := range resBundle.Items {
		assert.NoError(t, utils.VerifyBundleItem(item))
	}

	// send to arweave
	// wal, err := NewWalletFromPath("test-keyfile.json", "https://arweave.net")
	// assert.NoError(t, err)
	// tx, err := wal.SendBundleTx(context.TODO(), 0, bundle.Binary, []schema.Tag{
	// 	{Name: "App", Value: "goar"},
	// })
	// assert.NoError(t, err)
	// t.Log(tx.ID)
}

func TestNestBundle(t *testing.T) {
	b1, err := NewBundler(signer01)
	assert.NoError(t, err)
	item1, err := b1.CreateAndSignItem([]byte("eth foo"), "", "", []schema.Tag{{Name: "Content-Type", Value: "application/txt"}})
	assert.NoError(t, err)

	b2, err := NewBundler(signer02)
	assert.NoError(t, err)
	item2, err := b2.CreateAndSignItem([]byte("ar foo"), "", "", []schema.Tag{{Name: "Content-Type", Value: "application/txt"}})
	assert.NoError(t, err)

	// nested
	nestedItem, err := b1.CreateAndSignNestedItem("", "", []schema.Tag{}, item1, item2)
	assert.NoError(t, err)

	item3, err := b2.CreateAndSignItem([]byte("ar foo2"), "", "", []schema.Tag{{Name: "Content-Type", Value: "application/txt"}})
	assert.NoError(t, err)

	// bundle
	bundle, err := utils.NewBundle(nestedItem, item3)
	assert.NoError(t, err)

	resBundle, err := utils.DecodeBundle(bundle.Binary)
	assert.NoError(t, err)

	for _, item := range resBundle.Items {
		assert.NoError(t, utils.VerifyBundleItem(item))
	}

	// send to arweave
	// wal, err := NewWalletFromPath("test-keyfile.json", "https://arweave.net")
	// assert.NoError(t, err)
	// tx, err := wal.SendBundleTx(context.TODO(), 0, bundle.Binary, []schema.Tag{
	// 	{Name: "App", Value: "goar"},
	// })
	// assert.NoError(t, err)
	// t.Log(tx.ID)
}
