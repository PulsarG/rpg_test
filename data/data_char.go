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
		mainPercent: 15,

		pausCount:   0,
		pausPercent: 10,
	}
}

// ** GETERS

func (d *DataChar) GetMainCount() int {
	return d.mainCount
}

func (d *DataChar) GetMainPercent() int {
	return d.mainPercent
}

func (d *DataChar) GetPausCount() int {
	return d.pausCount
}

func (d *DataChar) GetPausPercent() int {
	return d.pausPercent
}

// ** SETERS

func (d *DataChar) SetMainCount(count int) {
	d.mainCount += count
}

func (d *DataChar) SetMainPercent(i int) {
	d.mainPercent = i
}

func (d *DataChar) SetPausCount(i int) {
	d.pausCount = i
}

func (d *DataChar) SetPausPercent(i int) {
	d.pausPercent = i
}
