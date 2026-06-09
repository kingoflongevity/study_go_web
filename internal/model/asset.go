package model

import (
	"time"

	"gorm.io/gorm"
)

// 资产状态枚举常量
const (
	// AssetStatusInStock 在库（闲置状态）
	AssetStatusInStock = "in_stock"
	// AssetStatusInUse 使用中
	AssetStatusInUse = "in_use"
	// AssetStatusMaintenance 维修中
	AssetStatusMaintenance = "maintenance"
	// AssetStatusScrapped 已报废
	AssetStatusScrapped = "scrapped"
)

// Asset 资产表
type Asset struct {
	gorm.Model
	// 资产编号，唯一索引
	AssetNo string `gorm:"type:varchar(100);uniqueIndex;not null;comment:资产编号" json:"asset_no"`
	// 资产名称
	Name string `gorm:"type:varchar(200);not null;comment:资产名称" json:"name"`
	// 资产分类ID，外键关联 category.id
	CategoryID uint `gorm:"not null;index;comment:资产分类ID" json:"category_id"`
	// 资产分类（BelongsTo 关联）
	Category Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"category,omitempty"`
	// 资产状态：in_stock（在库）/ in_use（使用中）/ maintenance（维修中）/ scrapped（已报废）
	Status string `gorm:"type:varchar(20);not null;default:in_stock;comment:资产状态" json:"status"`
	// 采购日期，可空
	PurchaseDate *time.Time `gorm:"comment:采购日期" json:"purchase_date,omitempty"`
	// 采购价格，可空
	Price *float64 `gorm:"type:decimal(12,2);comment:采购价格" json:"price,omitempty"`
	// 当前使用人ID，可空。外键关联 employee.id
	CurrentHolderID *uint `gorm:"index;comment:当前使用人ID" json:"current_holder_id,omitempty"`
	// 当前使用人（BelongsTo 关联 Employee）
	CurrentHolder *Employee `gorm:"foreignKey:CurrentHolderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"current_holder,omitempty"`
}

// TableName 自定义表名
func (Asset) TableName() string {
	return "assets"
}
