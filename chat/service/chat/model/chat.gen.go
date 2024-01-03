// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChat = "chat"

// Chat mapped from table <chat>
type Chat struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:聊天记录ID" json:"id"`               // 聊天记录ID
	User       string    `gorm:"column:user;not null;comment:weCom用户标识/customer用户标识" json:"user"`                // weCom用户标识/customer用户标识
	MessageID  string    `gorm:"column:message_id;not null;comment:message_id customer消息唯一ID" json:"message_id"` // message_id customer消息唯一ID
	OpenKfID   string    `gorm:"column:open_kf_id;not null;comment:客服标识" json:"open_kf_id"`                      // 客服标识
	AgentID    int64     `gorm:"column:agent_id;not null;comment:应用ID" json:"agent_id"`                          // 应用ID
	ReqContent string    `gorm:"column:req_content;not null;comment:用户发送内容" json:"req_content"`                  // 用户发送内容
	ResContent string    `gorm:"column:res_content;not null;comment:openai响应内容" json:"res_content"`              // openai响应内容
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`     // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`     // 更新时间
}

// TableName Chat's table name
func (*Chat) TableName() string {
	return TableNameChat
}
