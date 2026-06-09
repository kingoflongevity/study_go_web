package model

import (
	"time"

	"gorm.io/gorm"
)

// AssetAllocation 资产分配记录表，记录资产的分配和归还历史
type AssetAllocation struct {
	gorm.Model
	// 资产ID，外键关联 asset.id
	AssetID uint `gorm:"not null;index;comment:资产ID" json:"asset_id"`
	// 关联资产（BelongsTo 关联）
	Asset Asset `gorm:"foreignKey:AssetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"asset,omitempty"`
	// 员工ID，外键关联 employee.id
	EmployeeID uint `gorm:"not null;index;comment:员工ID" json:"employee_id"`
	// 关联员工（BelongsTo 关联）
	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
	// 分配时间
	AllocatedAt time.Time `gorm:"not null;comment:分配时间" json:"allocated_at"`
	// 归还时间，可空。为空表示尚未归还
	ReturnedAt *time.Time `gorm:"comment:归还时间" json:"returned_at,omitempty"`
}

// TableName 自定义表名
func (AssetAllocation) TableName() string {
	return "asset_allocations"
}
