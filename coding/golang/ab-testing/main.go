package main

import (
	"fmt"

	"github.com/bamwidyat/coding-exercise/golang-ab-testing/algorithm"
	"github.com/bamwidyat/coding-exercise/golang-ab-testing/database/imaginedb"
)

func main() {
	db := imaginedb.Init()
	alg := algorithm.Init(db)

	// Run Tests Here
	feat1 := alg.GetFeature(1, 1)
	feat2 := alg.GetFeature(245, 1)
	feat3 := alg.GetFeature(3667765, 2)
	fmt.Println(feat1)
	fmt.Println(feat2)
	fmt.Println(feat3)
}
