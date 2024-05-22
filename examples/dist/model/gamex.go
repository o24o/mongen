package model

type Gamex struct {
	BaseDoc        `bson:",inline"`
	LocalCategory  string `json:"LocalCategory" bson:"localCategory"`
	LocalBrand     string `json:"LocalBrand" bson:"localBrand"`
	LocalGame      string `json:"localGame" bson:"localGame"`
	VendorCode     string `json:"vendorCode" bson:"vendorCode"`
	GameCode       string `json:"gameCode" bson:"gameCode"`
	GameName       string `json:"gameName" bson:"gameName"`
	CategoryCode   string `json:"categoryCode" bson:"categoryCode"`
	ImageSquare    string `json:"imageSquare" bson:"imageSquare"`
	ImageLandscape string `json:"imageLandscape" bson:"imageLandscape"`
	LanguageCode   string `json:"languageCode" bson:"languageCode"`
	PlatformCode   string `json:"platformCode" bson:"platformCode"`
	CurrencyCode   string `json:"currencyCode" bson:"currencyCode"`
}
