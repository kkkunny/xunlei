package dto

import "fmt"

// TaskPhase 任务状态
type TaskPhase string

const (
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
	TaskTypeUserDownloadURL TaskType = "user#download-url"
	TaskTypeUserDownload    TaskType = "user#download"
)
