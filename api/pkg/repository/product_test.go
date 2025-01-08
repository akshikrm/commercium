package repository

import (
	config "akshidas/e-com"
	"testing"
)

func TestInsertImage(t *testing.T) {
	config := config.NewTestConfig()
	store := New(config)
	images := []string{"008_ie42n6", "005_rgjfhk"}
	err := store.Product.InsertImages(5, images)
	if err != nil {
		t.Error("failed to insert images")
	}
}
