package protomodel

// * This is a generated file, do not edit

// CoinbaseFees are fees rates and 30 days trailing volume.
type CoinbaseFees struct {
	MakerFeeRate float64 `json:"maker_fee_rate"`
	TakerFeeRate float64 `json:"taker_fee_rate"`
	UsdVolume    float64 `json:"usd_volume"`
}
