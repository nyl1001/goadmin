// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-vue-admin/internal/app/system/dao/internal"
)

// internalSysJobDao is internal type for wrapping internal DAO implements.
type internalSysJobDao = *internal.SysJobDao

// sysJobDao is the data access object for table sys_job.
// You can define custom methods on it to extend its functionality as you wish.
type sysJobDao struct {
	internalSysJobDao
}

var (
	// SysJob is globally public accessible object for table sys_job operations.
	SysJob = sysJobDao{
		internal.NewSysJobDao(),
	}
)

// Fill with you ideas below.
