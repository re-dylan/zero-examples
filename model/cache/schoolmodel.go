package cache

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchoolModel = (*customSchoolModel)(nil)

type (
	// SchoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchoolModel.
	SchoolModel interface {
		schoolModel
	}

	customSchoolModel struct {
		*defaultSchoolModel
	}
)

// NewSchoolModel returns a model for the database table.
func NewSchoolModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SchoolModel {
	return &customSchoolModel{
		defaultSchoolModel: newSchoolModel(conn, c, opts...),
	}
}

func newSchoolModelFromCachedConn(conn sqlc.CachedConn) SchoolModel {
	return &customSchoolModel{
		defaultSchoolModel: &defaultSchoolModel{
			CachedConn: conn,
			table:      "`school`",
		},
	}
}
