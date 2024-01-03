// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChatConfig = "chat_config"

// ChatConfig mapped from table <chat_config>
type ChatConfig struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:聊天配置ID" json:"id"`           // 聊天配置ID
	User      string    `gorm:"column:user;not null;comment:用户标识" json:"user"`                              // 用户标识
	AgentID   int64     `gorm:"column:agent_id;not null;comment:应用ID" json:"agent_id"`                      // 应用ID
	Model     string    `gorm:"column:model;not null;comment:模型" json:"model"`                              // 模型
	Prompt    string    `gorm:"column:prompt;not null;comment:系统初始设置" json:"prompt"`                        // 系统初始设置
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName ChatConfig's table name
func (*ChatConfig) TableName() string {
	return TableNameChatConfig
}
