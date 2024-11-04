package dto

import (
	"fmt"
	"time"
)

// TaskPhase 任务状态
type TaskPhase string

const (
	TaskPhaseTypeUnknown  TaskPhase = ""
	TaskPhaseTypePending  TaskPhase = "PHASE_TYPE_PENDING"
	TaskPhaseTypeRunning  TaskPhase = "PHASE_TYPE_RUNNING"
	TaskPhaseTypePaused   TaskPhase = "PHASE_TYPE_PAUSED"
	TaskPhaseTypeError    TaskPhase = "PHASE_TYPE_ERROR"
	TaskPhaseTypeComplete TaskPhase = "PHASE_TYPE_COMPLETE"
	TaskPhaseTypeDelete   TaskPhase = "PHASE_TYPE_DELETE" // 包括文件
)

func (phase TaskPhase) Spec() string {
	switch phase {
	case TaskPhaseTypeRunning:
		return "running"
	case TaskPhaseTypePaused:
		return "pause"
	case TaskPhaseTypeDelete:
		return "delete"
	default:
		panic(fmt.Sprintf("unknown TaskPhase `%s`", phase))
	}
}

// TaskType 任务类型
type TaskType string

const (
	TaskTypeUserUnknown     TaskType = ""
	TaskTypeUserDownloadURL TaskType = "user#download-url"
	TaskTypeUserDownload    TaskType = "user#download"
)

type TaskInfo struct {
	ID       string    // 任务ID
	Type     TaskType  // 任务类型
	Name     string    // 任务名
	UserID   string    // 用户ID
	Phase    TaskPhase // 任务状态
	Progress int64     // 下载进度百分比

	FileID   string // 文件ID
	FileName string // 文件名
	FileSize int64  // 文件大小 单位：byte

	SavePath string // 保存地址
	Speed    int64  // 下载速度 单位：byte
	URL      string // 下载链接

	Space string // 所属空间

	Extra       map[string]string
	CreatedTime time.Time // 创建时间
	UpdatedTime time.Time // 更新时间
}
