package model

type OwnershipResponse struct {
	ID            string        `json:"id"`
	Blockchain    string        `json:"blockchain"`
	ItemID        string        `json:"itemId"`
	Contract      string        `json:"contract"`
	Collection    string        `json:"collection"`
	TokenID       string        `json:"tokenId"`
	Owner         string        `json:"owner"`
	Value         string        `json:"value"`
	Source        string        `json:"source"`
	CreatedAt     string        `json:"createdAt"`
	LastUpdatedAt string        `json:"lastUpdatedAt"`
	LazyValue     string        `json:"lazyValue"`
	Pending       []Pending     `json:"pending"`
	BestSellOrder *Order        `json:"bestSellOrder"`
	OriginOrders  []OriginOrder `json:"originOrders"`
	Version       int           `json:"version"`
}

type Pending struct {
	Type      string    `json:"@type"`
	Royalties []Royalty `json:"royalties,omitempty"`
	From      string    `json:"from,omitempty"`
}

type Royalty struct {
	Account string `json:"account"`
	Value   int64  `json:"value"`
}

type Order struct {
	ID                string    `json:"id"`
	Fill              float64   `json:"fill"`
	Platform          string    `json:"platform"`
	Status            string    `json:"status"`
	StartedAt         string    `json:"startedAt"`
	EndedAt           string    `json:"endedAt"`
	MakeStock         float64   `json:"makeStock"`
	Cancelled         bool      `json:"cancelled"`
	OptionalRoyalties bool      `json:"optionalRoyalties"`
	CreatedAt         string    `json:"createdAt"`
	LastUpdatedAt     string    `json:"lastUpdatedAt"`
	DbUpdatedAt       string    `json:"dbUpdatedAt"`
	MakePrice         float64   `json:"makePrice"`
	TakePrice         float64   `json:"takePrice"`
	MakePriceUsd      float64   `json:"makePriceUsd"`
	TakePriceUsd      float64   `json:"takePriceUsd"`
	Maker             string    `json:"maker"`
	Taker             string    `json:"taker"`
	Make              Asset     `json:"make"`
	Take              Asset     `json:"take"`
	Salt              string    `json:"salt"`
	Signature         string    `json:"signature"`
	FeeTakers         []string  `json:"feeTakers"`
	Data              OrderData `json:"data"`
	Version           int       `json:"version"`
}

type Asset struct {
	Type  AssetType `json:"type"`
	Value float64   `json:"value"`
}

type AssetType struct {
	Blockchain string `json:"blockchain"`
	Contract   string `json:"contract"`
	Type       string `json:"@type"`
}

type OrderData struct {
	Type string      `json:"@type"`
	Data interface{} `json:"data"`
}

type OriginOrder struct {
	Origin        string `json:"origin"`
	BestSellOrder *Order `json:"bestSellOrder,omitempty"`
	BestBidOrder  *Order `json:"bestBidOrder,omitempty"`
}

type RarityTrait struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity string `json:"rarity"`
}

type RarityResponse struct {
	Continuation string        `json:"continuation"`
	Traits       []RarityTrait `json:"traits"`
}

type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RarityRequest struct {
	CollectionId string     `json:"collectionId"`
	Properties   []Property `json:"properties"`
}

type OwnershipRequest struct {
	OwnershipId string `json:"ownershipId"`
}
