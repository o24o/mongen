package dao

import (
	"gen/common"
	"main/dist/model"
)

type gamex struct {
	common.Q[model.Gamex]
}

var Gamex = &gamex{}

func (q *gamex) ID() common.Field[model.ID] {
	return common.Field[model.ID]{Name: "ID", Bson: "_id"}
}
func (q *gamex) CreatedAt() common.Field[int64] {
	return common.Field[int64]{Name: "CreatedAt", Bson: "createdAt"}
}
func (q *gamex) UpdatedAt() common.Field[int64] {
	return common.Field[int64]{Name: "UpdatedAt", Bson: "updatedAt"}
}
func (q *gamex) Version() common.Field[int64] {
	return common.Field[int64]{Name: "Version", Bson: "version"}
}
func (q *gamex) LocalCategory() common.Field[string] {
	return common.Field[string]{Name: "LocalCategory", Bson: "localCategory"}
}
func (q *gamex) LocalBrand() common.Field[string] {
	return common.Field[string]{Name: "LocalBrand", Bson: "localBrand"}
}
func (q *gamex) LocalGame() common.Field[string] {
	return common.Field[string]{Name: "LocalGame", Bson: "localGame"}
}
func (q *gamex) VendorCode() common.Field[string] {
	return common.Field[string]{Name: "VendorCode", Bson: "vendorCode"}
}
func (q *gamex) GameCode() common.Field[string] {
	return common.Field[string]{Name: "GameCode", Bson: "gameCode"}
}
func (q *gamex) GameName() common.Field[string] {
	return common.Field[string]{Name: "GameName", Bson: "gameName"}
}
func (q *gamex) CategoryCode() common.Field[string] {
	return common.Field[string]{Name: "CategoryCode", Bson: "categoryCode"}
}
func (q *gamex) ImageSquare() common.Field[string] {
	return common.Field[string]{Name: "ImageSquare", Bson: "imageSquare"}
}
func (q *gamex) ImageLandscape() common.Field[string] {
	return common.Field[string]{Name: "ImageLandscape", Bson: "imageLandscape"}
}
func (q *gamex) LanguageCode() common.Field[string] {
	return common.Field[string]{Name: "LanguageCode", Bson: "languageCode"}
}
func (q *gamex) PlatformCode() common.Field[string] {
	return common.Field[string]{Name: "PlatformCode", Bson: "platformCode"}
}
func (q *gamex) CurrencyCode() common.Field[string] {
	return common.Field[string]{Name: "CurrencyCode", Bson: "currencyCode"}
}
