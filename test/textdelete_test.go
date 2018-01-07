package test

import (
	"Blog-text/models"
	"strconv"
	"testing"
)

func TestDel(t *testing.T) {
	var textstest = []struct {
		want int64
		id   int
	}{
		{want: 1, id: 32},
		{want: 3, id: 3},
		{want: 2, id: 24},
		{want: 0, id: 42},
		{want: 0, id: 78}}

	for _, test := range textstest {
		Id := strconv.Itoa(test.id)
		if w, _ := models.DeleteText(Id); w != test.want {
			t.Errorf("DelBlog = %d", w)
		}
	}
}
