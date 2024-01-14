// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		Bot:            newBot(db, opts...),
		BotsPrompt:     newBotsPrompt(db, opts...),
		BotsWithCustom: newBotsWithCustom(db, opts...),
		BotsWithModel:  newBotsWithModel(db, opts...),
		Chat:           newChat(db, opts...),
		ChatConfig:     newChatConfig(db, opts...),
		PromptConfig:   newPromptConfig(db, opts...),
		User:           newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Bot            bot
	BotsPrompt     botsPrompt
	BotsWithCustom botsWithCustom
	BotsWithModel  botsWithModel
	Chat           chat
	ChatConfig     chatConfig
	PromptConfig   promptConfig
	User           user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Bot:            q.Bot.clone(db),
		BotsPrompt:     q.BotsPrompt.clone(db),
		BotsWithCustom: q.BotsWithCustom.clone(db),
		BotsWithModel:  q.BotsWithModel.clone(db),
		Chat:           q.Chat.clone(db),
		ChatConfig:     q.ChatConfig.clone(db),
		PromptConfig:   q.PromptConfig.clone(db),
		User:           q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Bot:            q.Bot.replaceDB(db),
		BotsPrompt:     q.BotsPrompt.replaceDB(db),
		BotsWithCustom: q.BotsWithCustom.replaceDB(db),
		BotsWithModel:  q.BotsWithModel.replaceDB(db),
		Chat:           q.Chat.replaceDB(db),
		ChatConfig:     q.ChatConfig.replaceDB(db),
		PromptConfig:   q.PromptConfig.replaceDB(db),
		User:           q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Bot            *botDo
	BotsPrompt     *botsPromptDo
	BotsWithCustom *botsWithCustomDo
	BotsWithModel  *botsWithModelDo
	Chat           *chatDo
	ChatConfig     *chatConfigDo
	PromptConfig   *promptConfigDo
	User           *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Bot:            q.Bot.WithContext(ctx),
		BotsPrompt:     q.BotsPrompt.WithContext(ctx),
		BotsWithCustom: q.BotsWithCustom.WithContext(ctx),
		BotsWithModel:  q.BotsWithModel.WithContext(ctx),
		Chat:           q.Chat.WithContext(ctx),
		ChatConfig:     q.ChatConfig.WithContext(ctx),
		PromptConfig:   q.PromptConfig.WithContext(ctx),
		User:           q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
