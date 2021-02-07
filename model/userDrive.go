package model

type UserDrive struct {
	DriveId      int64	`json:"DriveId" gorm:"primary_key;not_null;auto_increment"`
	UserName    string	`json:"userName" gorm:"unique;not_null"`
	DriveName string	`json:"driveName"`
	CpuModel string	`json:"cpuModel"`
	MemorySize string	`json:"memorySize"`
	DisplayCardModel string	`json:"displayCardModel"`
	HardSize string	`json:"diskSize"`
	SsdSize string	`json:"ssdSize"`
	Display1 string	`json:"display1"`
	Display2 string	`json:"display2"`
}
