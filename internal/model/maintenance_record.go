package model

import (
	"time"

	"gorm.io/gorm"
)

// MaintenanceRecord 维修记录表
type MaintenanceRecord struct {
	gorm.Model
	// 资产ID，外键关联 asset.id
	AssetID uint `gorm:"not null;index;comment:资产ID" json:"asset_id"`
	// 关联资产（BelongsTo 关联）
	Asset Asset `gorm:"foreignKey:AssetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"asset,omitempty"`
	// 维修描述
	Description string `gorm:"type:text;not null;comment:维修描述" json:"description"`
	// 维修费用
	Cost float64 `gorm:"type:decimal(12,2);not null;comment:维修费用" json:"cost"`
	// 维修日期
	Date time.Time `gorm:"not null;comment:维修日期" json:"date"`
}

// TableName 自定义表名
func (MaintenanceRecord) TableName() string {
	return "maintenance_records"
}
