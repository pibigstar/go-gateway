package model

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// TODO: 加入trace
func GetDB(ctx context.Context) gdb.DB {
	db := g.DB()
	return db
}
