package repository

import (
	config "akshidas/e-com"
	"fmt"
	"testing"
)

func TestInsertImage(t *testing.T) {
	config := config.NewTestConfig()
	store := New(config)
	images := []string{"008_ie42n6", "005_rgjfhk"}
	fmt.Println(config)
	fmt.Println(images)
	store.Product.InsertImages(5, images)
}
