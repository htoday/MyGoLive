package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LiveRoomModel = (*customLiveRoomModel)(nil)

type (
	// LiveRoomModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLiveRoomModel.
	LiveRoomModel interface {
		liveRoomModel
	}

	customLiveRoomModel struct {
		*defaultLiveRoomModel
	}
)

// NewLiveRoomModel returns a model for the database table.
func NewLiveRoomModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LiveRoomModel {
	return &customLiveRoomModel{
		defaultLiveRoomModel: newLiveRoomModel(conn, c, opts...),
	}
}
