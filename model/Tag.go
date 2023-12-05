package model

import "encoding/json"

// 标签
type Tag struct {
	Tags json.RawMessage `gorm:"type:json" json:"tags,omitempty"` // 标签
}
