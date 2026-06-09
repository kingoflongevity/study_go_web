package model

import "gorm.io/gorm"

// Employee 员工表，属于某个部门
type Employee struct {
	gorm.Model
	// 员工姓名
	Name string `gorm:"type:varchar(100);not null;comment:员工姓名" json:"name"`
	// 所属部门ID，外键关联 department.id
	DepartmentID uint `gorm:"not null;index;comment:所属部门ID" json:"department_id"`
	// 所属部门（BelongsTo 关联）
	Department Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"department,omitempty"`
}

// TableName 自定义表名
func (Employee) TableName() string {
	return "employees"
}
