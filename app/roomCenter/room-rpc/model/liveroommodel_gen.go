// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	liveRoomFieldNames          = builder.RawFieldNames(&LiveRoom{})
	liveRoomRows                = strings.Join(liveRoomFieldNames, ",")
	liveRoomRowsExpectAutoSet   = strings.Join(stringx.Remove(liveRoomFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	liveRoomRowsWithPlaceHolder = strings.Join(stringx.Remove(liveRoomFieldNames, "`roomId`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLiveRoomRoomIdPrefix = "cache:liveRoom:roomId:"
)

type (
	liveRoomModel interface {
		Insert(ctx context.Context, data *LiveRoom) (sql.Result, error)
		FindOne(ctx context.Context, roomId int64) (*LiveRoom, error)
		Update(ctx context.Context, data *LiveRoom) error
		Delete(ctx context.Context, roomId int64) error
		Query(ctx context.Context, page int64, pageSize int64) (*[]LiveRoom, error)
	}

	defaultLiveRoomModel struct {
		sqlc.CachedConn
		table string
	}

	LiveRoom struct {
		RoomId    int64          `db:"roomId"`
		RoomName  sql.NullString `db:"roomName"`
		RoomOwner sql.NullString `db:"roomOwner"`
		ImgPath   sql.NullString `db:"imgPath"`
		ImgName   sql.NullString `db:"imgName"`
	}
)

func newLiveRoomModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLiveRoomModel {
	return &defaultLiveRoomModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`liveRoom`",
	}
}

func (m *defaultLiveRoomModel) Query(ctx context.Context, page int64, pageSize int64) (*[]LiveRoom, error) {
	var resp []LiveRoom
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT ? OFFSET ?", liveRoomRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, pageSize, (page-1)*pageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultLiveRoomModel) Delete(ctx context.Context, roomId int64) error {

	liveRoomRoomIdKey := fmt.Sprintf("%s%v", cacheLiveRoomRoomIdPrefix, roomId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `roomId` = ?", m.table)
		return conn.ExecCtx(ctx, query, roomId)
	}, liveRoomRoomIdKey)
	return err
}

func (m *defaultLiveRoomModel) FindOne(ctx context.Context, roomId int64) (*LiveRoom, error) {
	liveRoomRoomIdKey := fmt.Sprintf("%s%v", cacheLiveRoomRoomIdPrefix, roomId)
	var resp LiveRoom
	err := m.QueryRowCtx(ctx, &resp, liveRoomRoomIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `roomId` = ? limit 1", liveRoomRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, roomId)
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

func (m *defaultLiveRoomModel) Insert(ctx context.Context, data *LiveRoom) (sql.Result, error) {
	liveRoomRoomIdKey := fmt.Sprintf("%s%v", cacheLiveRoomRoomIdPrefix, data.RoomId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, liveRoomRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RoomId, data.RoomName, data.RoomOwner, data.ImgPath, data.ImgName)
	}, liveRoomRoomIdKey)
	return ret, err
}

func (m *defaultLiveRoomModel) Update(ctx context.Context, data *LiveRoom) error {
	liveRoomRoomIdKey := fmt.Sprintf("%s%v", cacheLiveRoomRoomIdPrefix, data.RoomId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `roomId` = ?", m.table, liveRoomRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RoomName, data.RoomOwner, data.ImgPath, data.ImgName, data.RoomId)
	}, liveRoomRoomIdKey)
	return err
}

func (m *defaultLiveRoomModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLiveRoomRoomIdPrefix, primary)
}

func (m *defaultLiveRoomModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `roomId` = ? limit 1", liveRoomRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLiveRoomModel) tableName() string {
	return m.table
}