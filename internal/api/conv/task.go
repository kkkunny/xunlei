package conv

import (
	"strconv"
	"time"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

func ConvTaskTypeToDTO(typ string) dto.TaskType {
	switch typ {
	case string(dto.TaskTypeUserDownloadURL):
		return dto.TaskTypeUserDownloadURL
	case string(dto.TaskTypeUserDownload):
		return dto.TaskTypeUserDownload
	default:
		return dto.TaskTypeUserUnknown
	}
}

func ConvTaskPhaseToDTO(phase string) dto.TaskPhase {
	switch phase {
	case string(dto.TaskPhaseTypePending):
		return dto.TaskPhaseTypePending
	case string(dto.TaskPhaseTypeRunning):
		return dto.TaskPhaseTypeRunning
	case string(dto.TaskPhaseTypePaused):
		return dto.TaskPhaseTypePaused
	case string(dto.TaskPhaseTypeError):
		return dto.TaskPhaseTypeError
	case string(dto.TaskPhaseTypeComplete):
		return dto.TaskPhaseTypeComplete
	case string(dto.TaskPhaseTypeDelete):
		return dto.TaskPhaseTypeDelete
	default:
		return dto.TaskPhaseTypeUnknown
	}
}

func ConvTaskInfoToDTO(task *api.TaskInfo) (*dto.TaskInfo, error) {
	fileSize, err := strconv.ParseInt(task.FileSize, 10, 64)
	if err != nil {
		return nil, err
	}
	speed, err := strconv.ParseInt(task.Params["speed"], 10, 64)
	if err != nil {
		return nil, err
	}

	extra := make(map[string]string)
	for k, v := range task.Params {
		extra[k] = v
	}
	delete(extra, "real_path")
	delete(extra, "speed")
	delete(extra, "url")
	if task.Kind != "" {
		extra["kind"] = task.Kind
	}
	if task.IconLink != "" {
		extra["icon_link"] = task.IconLink
	}

	createTime, err := time.Parse(time.RFC3339, task.CreatedTime)
	if err != nil {
		return nil, err
	}
	updateTime, err := time.Parse(time.RFC3339, task.UpdatedTime)
	if err != nil {
		return nil, err
	}

	return &dto.TaskInfo{
		ID:          task.ID,
		Type:        ConvTaskTypeToDTO(task.Type),
		Name:        task.Name,
		UserID:      task.UserID,
		Phase:       ConvTaskPhaseToDTO(task.Phase),
		Progress:    task.Progress,
		FileID:      task.FileID,
		FileName:    task.FileName,
		FileSize:    fileSize,
		SavePath:    task.Params["real_path"],
		Speed:       speed,
		URL:         task.Params["url"],
		Space:       task.Space,
		Extra:       extra,
		CreatedTime: createTime,
		UpdatedTime: updateTime,
	}, nil
}
