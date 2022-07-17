package svc

import (
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/model"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookModel // 手动代码
}

const postgresDriverName = "postgres"

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookModel(sqlx.NewSqlConn(postgresDriverName, c.DataSource), c.Cache), // 手动代码
	}
}
