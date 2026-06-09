package model

import "gorm.io/gorm"

// Category 资产分类表，支持树形结构（自关联 parent_id）
type Category struct {
	gorm.Model
	// 分类名称
	Name string `gorm:"type:varchar(100);not null;comment:分类名称" json:"name"`
	// 父分类ID，可空。为空表示顶级分类
	ParentID *uint `gorm:"index;comment:父分类ID" json:"parent_id"`
	// 子分类列表（自关联）
	Children []Category `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"children,omitempty"`
}

// TableName 自定义表名
func (Category) TableName() string {
	return "categories"
}
