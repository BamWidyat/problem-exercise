package imaginedb

type Module struct {
	ID   int64
	Name string
}

type Feature struct {
	ID         int64
	Name       string
	Percentage float64
	Module     Module
}
