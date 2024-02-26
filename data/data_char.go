package data

type DataChar struct {
	mainCount   int
	mainPercent int

	pausCount   int
	pausPercent int
}

func GetStartChar() *DataChar {
	return &DataChar{
		mainCount:   0,
		mainPercent: 60,

		pausCount:   0,
		pausPercent: 20,
	}
}
