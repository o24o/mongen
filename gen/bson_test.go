package gen

import (
	"testing"
)

type Doc struct {
	ID int64 `bson:"_id"`
}

type Member struct {
	Doc  `bson:",inline"`
	Name string `bson:"name"`
}

func Test_getBsonTag(t *testing.T) {

	m := Member{}
	data := getBsonTag(&m, &m.ID)
	t.Log(data)
}
