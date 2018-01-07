package test

import (
	"Blog-text/models"
	"testing"
)

func TestTextGets(t *testing.T) {

	gts := make([]models.Text, 0)
	if gts, _ = models.GetAllText(); gts == nil {
		t.Errorf("没有文章！")
	}

}
