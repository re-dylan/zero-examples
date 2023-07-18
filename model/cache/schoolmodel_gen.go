// Code generated by goctl. DO NOT EDIT.

package cache

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	schoolFieldNames          = builder.RawFieldNames(&School{})
	schoolRows                = strings.Join(schoolFieldNames, ",")
	schoolRowsExpectAutoSet   = strings.Join(stringx.Remove(schoolFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	schoolRowsWithPlaceHolder = strings.Join(stringx.Remove(schoolFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSchoolIdPrefix = "cache:school:id:"
)

type (
	schoolModel interface {
		Insert(ctx context.Context, data *School) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*School, error)
		Update(ctx context.Context, data *School) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSchoolModel struct {
		sqlc.CachedConn
		table string
	}

	School struct {
		Id       int64          `db:"id"`
		Name     sql.NullString `db:"name"`    // The username
		UserId   int64          `db:"user_id"` // The user id
		Type     int64          `db:"type"`    // The user type, 0:normal,1:vip, for test golang keyword
		CreateAt sql.NullTime   `db:"create_at"`
		UpdateAt time.Time      `db:"update_at"`
	}
)

func newSchoolModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSchoolModel {
	return &defaultSchoolModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`school`",
	}
}

func (m *defaultSchoolModel) withSession(session sqlx.Session) *defaultSchoolModel {
	return &defaultSchoolModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`school`",
	}
}

func (m *defaultSchoolModel) Delete(ctx context.Context, id int64) error {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, schoolIdKey)
	return err
}

func (m *defaultSchoolModel) FindOne(ctx context.Context, id int64) (*School, error) {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, id)
	var resp School
	err := m.QueryRowCtx(ctx, &resp, schoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schoolRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSchoolModel) Insert(ctx context.Context, data *School) (sql.Result, error) {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, schoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.UserId, data.Type)
	}, schoolIdKey)
	return ret, err
}

func (m *defaultSchoolModel) Update(ctx context.Context, data *School) error {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, schoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.UserId, data.Type, data.Id)
	}, schoolIdKey)
	return err
}

func (m *defaultSchoolModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSchoolIdPrefix, primary)
}

func (m *defaultSchoolModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schoolRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSchoolModel) tableName() string {
	return m.table
}
