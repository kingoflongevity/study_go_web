package model

import "gorm.io/gorm"

// Department 部门表，支持树形结构（自关联 parent_id）
type Department struct {
	gorm.Model
	// 部门名称
	Name string `gorm:"type:varchar(100);not null;comment:部门名称" json:"name"`
	// 父部门ID，可空。为空表示顶级部门
	ParentID *uint `gorm:"index;comment:父部门ID" json:"parent_id"`
	// 子部门列表（自关联）
	Children []Department `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"children,omitempty"`
}

// TableName 自定义表名，统一使用 snake_case 复数形式
func (Department) TableName() string {
	return "departments"
}
