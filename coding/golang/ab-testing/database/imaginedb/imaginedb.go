package imaginedb

type Interface interface {
	GetAll() []Feature
}

type imaginedb struct {
}

func Init() Interface {
	return &imaginedb{}
}

func (db *imaginedb) GetAll() []Feature {
	return []Feature{
		{ID: 1, Name: "Red Banner", Module: Module{ID: 1, Name: "Banner Color"}},
		{ID: 2, Name: "Blue Banner", Module: Module{ID: 1, Name: "Banner Color"}},
		{ID: 3, Name: "Green Banner", Module: Module{ID: 1, Name: "Banner Color"}},
		{ID: 4, Name: "Yellow Banner", Module: Module{ID: 1, Name: "Banner Color"}},
		{ID: 5, Name: "Square Button", Module: Module{ID: 2, Name: "Button Shape"}},
		{ID: 6, Name: "Round Button", Module: Module{ID: 2, Name: "Button Shape"}},
		{ID: 7, Name: "Oval Button", Module: Module{ID: 2, Name: "Button Shape"}},
		{ID: 8, Name: "Small Font", Module: Module{ID: 3, Name: "Font Size"}},
		{ID: 9, Name: "Large Font", Module: Module{ID: 3, Name: "Font Size"}},
	}
}
