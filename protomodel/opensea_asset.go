package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/serial"
)

// * This is a generated file, do not edit

// Asset is the primary object in the OpenSea API is the, representing a unique
// digital item whose ownership is managed by the blockchain. The below
// CryptoSaga hero is an example of an asset shown on OpenSea.
type OpenseaAsset struct {
	AnimationOriginalUrl string               `json:"animation_original_url"`
	AnimationUrl         string               `json:"animation_url"`
	BackgroundColor      string               `json:"background_color"`
	Decimals             string               `json:"decimals"`
	Description          string               `json:"description"`
	ExternalLink         string               `json:"external_link"`
	Id                   int                  `json:"id"`
	ImageOriginalUrl     string               `json:"image_original_url"`
	ImagePreviewUrl      string               `json:"image_preview_url"`
	ImageThumbnailUrl    string               `json:"image_thumbnail_url"`
	ImageUrl             string               `json:"image_url"`
	IsPresale            bool                 `json:"is_presale"`
	LastSale             float64              `json:"last_sale"`
	ListingDate          time.Time            `json:"listing_date"`
	Name                 string               `json:"name"`
	NumSales             float64              `json:"num_sales"`
	Permalink            string               `json:"permalink"`
	ProtoAssetContract   OpenseaAssetContract `json:"asset_contract"`
	ProtoCollection      OpenseaCollection    `json:"collection"`
	ProtoCreator         OpenseaCreator       `json:"creator"`
	ProtoOwner           OpenseaOwner         `json:"owner"`
	TokenId              string               `json:"token_id"`
	TokenMetadata        string               `json:"token_metadata"`
	TopBid               float64              `json:"top_bid"`
	TransferFee          float64              `json:"transfer_fee"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaAsset model
func (openseaAsset *OpenseaAsset) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                   = "id"
		tokenIdJsonTag              = "token_id"
		numSalesJsonTag             = "num_sales"
		backgroundColorJsonTag      = "background_color"
		imageUrlJsonTag             = "image_url"
		imagePreviewUrlJsonTag      = "image_preview_url"
		imageThumbnailUrlJsonTag    = "image_thumbnail_url"
		imageOriginalUrlJsonTag     = "image_original_url"
		animationUrlJsonTag         = "animation_url"
		animationOriginalUrlJsonTag = "animation_original_url"
		nameJsonTag                 = "name"
		descriptionJsonTag          = "description"
		externalLinkJsonTag         = "external_link"
		assetContractJsonTag        = "asset_contract"
		permalinkJsonTag            = "permalink"
		collectionJsonTag           = "collection"
		decimalsJsonTag             = "decimals"
		tokenMetadataJsonTag        = "token_metadata"
		ownerJsonTag                = "owner"
		creatorJsonTag              = "creator"
		lastSaleJsonTag             = "last_sale"
		topBidJsonTag               = "top_bid"
		listingDateJsonTag          = "listing_date"
		isPresaleJsonTag            = "is_presale"
		transferFeeJsonTag          = "transfer_fee"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(isPresaleJsonTag, &openseaAsset.IsPresale)
	data.UnmarshalFloat(lastSaleJsonTag, &openseaAsset.LastSale)
	data.UnmarshalFloat(numSalesJsonTag, &openseaAsset.NumSales)
	data.UnmarshalFloat(topBidJsonTag, &openseaAsset.TopBid)
	data.UnmarshalFloat(transferFeeJsonTag, &openseaAsset.TransferFee)
	data.UnmarshalInt(idJsonTag, &openseaAsset.Id)
	data.UnmarshalString(animationOriginalUrlJsonTag, &openseaAsset.AnimationOriginalUrl)
	data.UnmarshalString(animationUrlJsonTag, &openseaAsset.AnimationUrl)
	data.UnmarshalString(backgroundColorJsonTag, &openseaAsset.BackgroundColor)
	data.UnmarshalString(decimalsJsonTag, &openseaAsset.Decimals)
	data.UnmarshalString(descriptionJsonTag, &openseaAsset.Description)
	data.UnmarshalString(externalLinkJsonTag, &openseaAsset.ExternalLink)
	data.UnmarshalString(imageOriginalUrlJsonTag, &openseaAsset.ImageOriginalUrl)
	data.UnmarshalString(imagePreviewUrlJsonTag, &openseaAsset.ImagePreviewUrl)
	data.UnmarshalString(imageThumbnailUrlJsonTag, &openseaAsset.ImageThumbnailUrl)
	data.UnmarshalString(imageUrlJsonTag, &openseaAsset.ImageUrl)
	data.UnmarshalString(nameJsonTag, &openseaAsset.Name)
	data.UnmarshalString(permalinkJsonTag, &openseaAsset.Permalink)
	data.UnmarshalString(tokenIdJsonTag, &openseaAsset.TokenId)
	data.UnmarshalString(tokenMetadataJsonTag, &openseaAsset.TokenMetadata)
	err = data.UnmarshalTime(time.RFC3339Nano, listingDateJsonTag, &openseaAsset.ListingDate)
	if err != nil {
		return err
	}
	openseaAsset.ProtoAssetContract = OpenseaAssetContract{}
	if err := data.UnmarshalStruct(assetContractJsonTag, &openseaAsset.ProtoAssetContract); err != nil {
		return err
	}
	openseaAsset.ProtoCollection = OpenseaCollection{}
	if err := data.UnmarshalStruct(collectionJsonTag, &openseaAsset.ProtoCollection); err != nil {
		return err
	}
	openseaAsset.ProtoCreator = OpenseaCreator{}
	if err := data.UnmarshalStruct(creatorJsonTag, &openseaAsset.ProtoCreator); err != nil {
		return err
	}
	openseaAsset.ProtoOwner = OpenseaOwner{}
	if err := data.UnmarshalStruct(ownerJsonTag, &openseaAsset.ProtoOwner); err != nil {
		return err
	}
	return nil
}
