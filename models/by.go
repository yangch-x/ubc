package models

import (
	"time"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建  人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"comment:最后更新时间"`
	DeletedAt *time.Time `json:"-" gorm:"comment:删除时间"`
}
