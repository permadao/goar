package goar

import (
	"crypto/sha256"
	"errors"

	"github.com/everFinance/goether"
	"github.com/permadao/goar/schema"
	"github.com/permadao/goar/utils"
)

type Bundler struct {
	signType   int
	signer     interface{}
	owner      string // only rsa has owner
	signerAddr string
}

func NewBundler(signer interface{}) (*Bundler, error) {
	signType, signerAddr, owner, err := reflectSigner(signer)
	if err != nil {
		return nil, err
	}
	return &Bundler{
		signType:   signType,
		signer:     signer,
		owner:      owner,
		signerAddr: signerAddr,
	}, nil
}

func (b *Bundler) Sign(item *schema.BundleItem) error {
	signMsg, err := utils.BundleItemSignData(*item)
	if err != nil {
		return err
	}
	var sigData []byte
	switch b.signType {
	case schema.ArweaveSignType:
		arSigner, ok := b.signer.(*Signer)
		if !ok {
			return errors.New("signer must be goar signer")
		}
		sigData, err = utils.Sign(signMsg, arSigner.PrvKey)
		if err != nil {
			return err
		}

	case schema.EthereumSignType:
		ethSigner, ok := b.signer.(*goether.Signer)
		if !ok {
			return errors.New("signer not goether signer")
		}
		sigData, err = ethSigner.SignMsg(signMsg)
		if err != nil {
			return err
		}
	default:
		// todo come soon support ed25519
		return errors.New("not support this signType")
	}
	id := sha256.Sum256(sigData)
	item.Id = utils.Base64Encode(id[:])
	item.Signature = utils.Base64Encode(sigData)
	return nil
}

func (b *Bundler) CreateAndSignItem(data []byte, target string, anchor string, tags []schema.Tag) (bItem schema.BundleItem, err error) {
	item, err := utils.NewBundleItem(b.owner, b.signType, target, anchor, data, tags)
	if err != nil {
		return
	}
	// sign
	if err = b.Sign(item); err != nil {
		return
	}
	// get itemBinary
	itemBinary, err := utils.GenerateItemBinary(item)
	if err != nil {
		return
	}
	item.Binary = itemBinary
	return *item, nil
}

func (b *Bundler) CreateAndSignNestedItem(target string, anchor string, tags []schema.Tag, items ...schema.BundleItem) (schema.BundleItem, error) {
	bundleTags := []schema.Tag{
		{Name: "Bundle-Format", Value: "binary"},
		{Name: "Bundle-Version", Value: "2.0.0"},
	}

	if err := checkBundleTags(tags); err != nil {
		return schema.BundleItem{}, err
	}

	tags = append(tags, bundleTags...)

	bundle, err := utils.NewBundle(items...)
	if err != nil {
		return schema.BundleItem{}, err
	}
	return b.CreateAndSignItem(bundle.Binary, target, anchor, tags)
}

func reflectSigner(signer interface{}) (signType int, signerAddr, owner string, err error) {
	if s, ok := signer.(*Signer); ok {
		signType = schema.ArweaveSignType
		signerAddr = s.Address
		owner = s.Owner()
		return
	}
	if s, ok := signer.(*goether.Signer); ok {
		signType = schema.EthereumSignType
		signerAddr = s.Address.String()
		owner = utils.Base64Encode(s.GetPublicKey())
		return
	}
	err = errors.New("not support this signer")
	return
}

func checkBundleTags(tags []schema.Tag) error {
	mmap := map[string]struct{}{
		"Bundle-Format":  {},
		"Bundle-Version": {},
	}
	for _, tag := range tags {
		if _, ok := mmap[tag.Name]; ok {
			return errors.New("tags can not set bundleTags")
		}
	}

	return nil
}
