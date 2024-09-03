package algorithm

import "github.com/bamwidyat/coding-exercise/golang-ab-testing/database/imaginedb"

type Interface interface {
	GetFeature(userId, moduleId int64) imaginedb.Feature
}

type algorithm struct {
	db imaginedb.Interface
}

func Init(db imaginedb.Interface) Interface {
	return &algorithm{
		db: db,
	}
}

func (a *algorithm) GetFeature(userId, moduleId int64) imaginedb.Feature {
	// Get Data From Imaginary Database
	// features := a.db.GetAll()

	// Your Code Here
	// ...
	return imaginedb.Feature{}
}
