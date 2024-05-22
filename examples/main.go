package main

import "gen"

func main() {

	type ID int64
	type BaseDoc struct {
		ID        ID    `json:"id" bson:"_id"`
		CreatedAt int64 `json:"createdAt" bson:"createdAt"`
		UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"`
		Version   int64 `json:"version" bson:"version"`
	}

	type vendor struct {
		BaseDoc      `bson:",inline"`
		LocalName    string `json:"localName" bson:"localName"`
		Name         string `json:"name" bson:"name"`
		CurrencyCode string `json:"currencyCode" bson:"currencyCode"`
		Code         string `json:"code" bson:"code"`
		CategoryCode string `json:"categoryCode" bson:"categoryCode"`
	}

	type gamex struct {
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

	gen.Gen(gamex{}, "")
}
