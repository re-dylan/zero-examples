package cache

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	UserSchoolModel interface {
		InsertUserSchool(ctx context.Context, u *User, s *School) error
	}

	customUserSchoolModel struct {
		conn sqlc.CachedConn
	}
)

func NewCustomUserSchoolModel(conn sqlc.CachedConn) UserSchoolModel {
	return &customUserSchoolModel{
		conn: conn,
	}
}

func (m *customUserSchoolModel) InsertUserSchool(ctx context.Context, u *User, s *School) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := m.conn.WithSession(session)
		r, err := NewUserModelForConn(conn).Insert(ctx, u)
		if err != nil {
			return err
		}
		id, err := r.LastInsertId()
		if err != nil {
			return err
		}
		s.UserId = id

		err = NewSchoolModelForConn(conn).Update(ctx, s)
		return err
	})
}
