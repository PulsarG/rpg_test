package data

type DataChar struct {
	mainCount   int
	mainPercent int
	hardCount   int

	pausCount   int
	pausPercent int
}

func GetStartChar() *DataChar {
	return &DataChar{
		mainCount:   0,
		mainPercent: 60, // 60

		hardCount: 0,

		pausCount:   0,
		pausPercent: 20, // 20
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

func (d *DataChar) GetHardCount() int {
	return d.hardCount
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

func (d *DataChar) SetHardCount(i int) {
	d.hardCount = i
}
