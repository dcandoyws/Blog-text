package test

import (
	"Blog-text/models"
	"testing"
	"time"
)

func TestAdd(ti *testing.T) {
	var textstest = []struct {
		text models.Text
		want int64
	}{

		{text: models.Text{1, "", "", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "dad", "", time.Now(), time.Now()}, want: 0},
		{text: models.Text{1, "ww", "2faf", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "dadc", "", time.Now(), time.Now()}, want: 0},
		{text: models.Text{1, "ca", " ", time.Now(), time.Now()}, want: 0},
		{text: models.Text{1, "1vsve", " ", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "aa", "222", time.Now(), time.Now()}, want: 0},
		{text: models.Text{1, " da", " ", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "333", "", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "666", "43", time.Now(), time.Now()}, want: 1},
		{text: models.Text{1, "", "8fa", time.Now(), time.Now()}, want: 1},
	}
	for _, test := range textstest {
		text := test.text
		if gt, _ := text.Addtext(); gt != test.want {
			ti.Errorf("Addtext = %d", gt)
		}
	}
}
