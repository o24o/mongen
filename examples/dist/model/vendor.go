package model

type ID int64
type BaseDoc struct {
	ID        ID    `json:"id" bson:"_id"`
	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"`
	Version   int64 `json:"version" bson:"version"`
}

type Vendor struct {
	BaseDoc      `bson:",inline"`
	LocalName    string `json:"localName" bson:"localName"`
	Name         string `json:"name" bson:"name"`
	CurrencyCode string `json:"currencyCode" bson:"currencyCode"`
	Code         string `json:"code" bson:"code"`
	CategoryCode string `json:"categoryCode" bson:"categoryCode"`
}
