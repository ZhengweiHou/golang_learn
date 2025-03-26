package gormdemo

type Teacher struct {
	Name    string `gorm:"primaryKey;column:NAME;size:255;not null"`
	Age     int    `gorm:"column:AGE"`
	Version int32  `gorm:"column:VERSION;default:0"`

	// gorm.Model // this is a struct that contains ID, CreatedAt, UpdatedAt, DeletedAt
}

func (Teacher) TableName() string {
	return "teacher"
}
