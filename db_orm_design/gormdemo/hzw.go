package gormdemo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Hzw struct {
	//gorm.Model      // this is a struct that contains ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:NAME;size:100;not null"`
	Age       int       `gorm:"column:AGE"`
	Version   int32     `gorm:"column:VERSION;default:0"`
	CreatedAt time.Time `gorm:"column:CREATED_AT"`           // CreatedAt is a field that contains create time
	UpdatedAt time.Time `gorm:"column:UPDATED_AT"`           // UpdatedAt is a field that contains update time
	Time1     time.Time `gorm:"AUTOUPDATETIME;column:TIME1"` // AUTOUPDATETIME means update time when update, like UpdatedAt
	Time2     time.Time `gorm:"AUTOCREATETIME;column:TIME2"` // AUTOCREATETIME means create time when insert, like CreatedAt
	Time3     time.Time `gorm:"column:TIME3"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Hzw) TableName() string {
	return "hzw"
}

func (h *Hzw) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Printf("====== BeforeSave  name=%s ======\n", h.Name)
	return
}
func (h *Hzw) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("====== BeforeCreate  name=%s ======\n", h.Name)
	return
}
func (h *Hzw) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("====== AfterCreate  name=%s id=%d ======\n", h.Name, h.ID)
	return
}
func (h *Hzw) AfterSave(tx *gorm.DB) (err error) {
	fmt.Printf("====== AfterSave  name=%s id=%d ======\n", h.Name, h.ID)
	return
}
