package dao

import (
	"gen/common"
	"main/dist/model"
)

type vendor struct {
	common.Q[model.Vendor]
}

var Vendor = &vendor{}

func (q *vendor) ID() common.Field[model.ID] {
	return common.Field[model.ID]{Name: "ID", Bson: "_id"}
}
func (q *vendor) CreatedAt() common.Field[int64] {
	return common.Field[int64]{Name: "CreatedAt", Bson: "createdAt"}
}
func (q *vendor) UpdatedAt() common.Field[int64] {
	return common.Field[int64]{Name: "UpdatedAt", Bson: "updatedAt"}
}
func (q *vendor) Version() common.Field[int64] {
	return common.Field[int64]{Name: "Version", Bson: "version"}
}
func (q *vendor) LocalName() common.Field[string] {
	return common.Field[string]{Name: "LocalName", Bson: "localName"}
}
func (q *vendor) Name() common.Field[string] {
	return common.Field[string]{Name: "Name", Bson: "name"}
}
func (q *vendor) CurrencyCode() common.Field[string] {
	return common.Field[string]{Name: "CurrencyCode", Bson: "currencyCode"}
}
func (q *vendor) Code() common.Field[string] {
	return common.Field[string]{Name: "Code", Bson: "code"}
}
func (q *vendor) CategoryCode() common.Field[string] {
	return common.Field[string]{Name: "CategoryCode", Bson: "categoryCode"}
}
