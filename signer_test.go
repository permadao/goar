package goar

import (
	"testing"

	"github.com/permadao/goar/schema"
	"github.com/permadao/goar/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewSigner(t *testing.T) {
	signer, err := NewSignerFromPath("testKey.json")
	assert.NoError(t, err)
	tags := []schema.Tag{
		{Name: "GOAR", Value: "sendAR"},
	}
	tx := &schema.Transaction{
		Format:    2,
		ID:        "EyPNVxI-zv1WGjiMmeb20VimIGjfvnkFQtgVnlSMX1o",
		LastTx:    "0cJJGyXdeIJ-azX3I-jJ3XJEIYti-_qjHZhaRsXQxi1D3wXQVv6Px5WQhj_j1W8O",
		Owner:     signer.Owner(),
		Tags:      utils.TagsEncode(tags),
		Target:    "cSYOy8-p1QFenktkDBFyRM3cwZSTrQ_J4EsELLho_UE",
		Quantity:  "1000000000",
		Data:      "",
		DataSize:  "0",
		DataRoot:  "",
		Reward:    "747288",
		Signature: "",
	}
	err = signer.SignTx(tx)
	assert.NoError(t, err)

	err = utils.VerifyTransaction(*tx)
	assert.NoError(t, err)
}
