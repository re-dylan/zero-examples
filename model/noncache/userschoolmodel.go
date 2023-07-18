package noncache

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	UserSchoolModel interface {
		InsertUserSchool(ctx context.Context, u *User, s *School) error
	}

	customUserSchoolModel struct {
		conn        sqlx.SqlConn
		userModel   UserModel
		schoolModel *customSchoolModel
	}
)

func NewCustomUserSchoolModel(conn sqlx.SqlConn) UserSchoolModel {
	return &customUserSchoolModel{
		conn: conn,
	}
}

func (m *customUserSchoolModel) InsertUserSchool(ctx context.Context, u *User, s *School) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)
		r, err := NewUserModel(conn).Insert(ctx, u)
		if err != nil {
			return err
		}
		id, err := r.LastInsertId()
		if err != nil {
			return err
		}
		s.UserId = id

		err = NewSchoolModel(conn).Update(ctx, s)
		return err
	})
}
