package main

import (
	"context"
	"fmt"
	"main/dist/dao"
	"main/tools"
)

func main() {
	ctx := context.Background()
	colVendor := tools.GetClient().Database("oneapi").Collection("vendor")
	first, err := dao.Vendor.Collection(colVendor).WithContext(ctx).Where(dao.Vendor.ID().Eq(1792842643722469376)).First()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", first)
}
