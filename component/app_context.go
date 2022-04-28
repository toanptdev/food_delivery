package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
}

func NewAppCtx(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db: db, secretKey: secretKey}
}

func (a *appCtx) GetMainDBConnection() *gorm.DB {
	return a.db
}

func (a *appCtx) SecretKey() string {
	return a.secretKey
}
