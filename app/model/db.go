package model

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var db gdb.DB

func init() {
	db = g.DB()
}
