package hooks

import (
	"gorm.io/gorm"
)

func HelloHook(m Models, tx *gorm.DB) error {
	m.TableName()
	// TODO 记录审计日志
	return nil
}

// =====框架的一部分=======
type Models interface {
	TableName() string
}

// TODO 根据表映射获取hook
type HookManager struct {
	BeforeCreate []func(Models, *gorm.DB) error
	AfterCreate  []func(Models, *gorm.DB) error
	// ...
}

func (hm *HookManager) RegisterHook(hooktype int, f func(Models, *gorm.DB) error) {
	if hooktype == 1 {
		hm.BeforeCreate = append(hm.BeforeCreate, f)
	}
}

func (hm *HookManager) SelectHook(tbname string) []func(Models, *gorm.DB) error {
	return nil
}
