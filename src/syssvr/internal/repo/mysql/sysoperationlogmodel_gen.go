// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysOperationLogFieldNames          = builder.RawFieldNames(&SysOperationLog{})
	sysOperationLogRows                = strings.Join(sysOperationLogFieldNames, ",")
	sysOperationLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysOperationLogFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	sysOperationLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysOperationLogFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	sysOperationLogModel interface {
		Insert(ctx context.Context, data *SysOperationLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysOperationLog, error)
		Update(ctx context.Context, data *SysOperationLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysOperationLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysOperationLog struct {
		Id           int64          `db:"id"`           // 编号
		OptUid       int64          `db:"optUid"`       // 用户id
		OptUserName  string         `db:"optUserName"`  // 操作人员名称
		OptName      string         `db:"optName"`      // 操作名称
		BusinessType int64          `db:"businessType"` // 业务类型（1新增 2修改 3删除 4查询）
		Uri          string         `db:"uri"`          // 请求地址
		OptIpAddr    string         `db:"optIpAddr"`    // 主机地址
		OptLocation  string         `db:"optLocation"`  // 操作地点
		Req          sql.NullString `db:"req"`          // 请求参数
		Resp         sql.NullString `db:"resp"`         // 返回参数
		Code         int64          `db:"code"`         // 登录状态（200成功 其它失败）
		Msg          string         `db:"msg"`          // 提示消息
		CreatedTime  time.Time      `db:"createdTime"`  // 操作时间
	}
)

func newSysOperationLogModel(conn sqlx.SqlConn) *defaultSysOperationLogModel {
	return &defaultSysOperationLogModel{
		conn:  conn,
		table: "`sys_operation_log`",
	}
}

func (m *defaultSysOperationLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysOperationLogModel) FindOne(ctx context.Context, id int64) (*SysOperationLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysOperationLogRows, m.table)
	var resp SysOperationLog
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysOperationLogModel) Insert(ctx context.Context, data *SysOperationLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysOperationLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OptUid, data.OptUserName, data.OptName, data.BusinessType, data.Uri, data.OptIpAddr, data.OptLocation, data.Req, data.Resp, data.Code, data.Msg)
	return ret, err
}

func (m *defaultSysOperationLogModel) Update(ctx context.Context, data *SysOperationLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysOperationLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OptUid, data.OptUserName, data.OptName, data.BusinessType, data.Uri, data.OptIpAddr, data.OptLocation, data.Req, data.Resp, data.Code, data.Msg, data.Id)
	return err
}

func (m *defaultSysOperationLogModel) tableName() string {
	return m.table
}
