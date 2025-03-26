package gormdemo

type Student struct {
	ID uint `gorm:"primaryKey;column:ID"`
	//	StuNo   string `gorm:"column:STU_NO;size:50;not null"`
	Name    string `gorm:"column:NAME;size:100;not null"`
	Age     int    `gorm:"column:AGE"`
	Version int32  `gorm:"column:VERSION;default:0"`

	// gorm.Model // this is a struct that contains ID, CreatedAt, UpdatedAt, DeletedAt
}

func (Student) TableName() string {
	return "student"
}
