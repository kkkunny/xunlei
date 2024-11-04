package conv

import (
	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

func ConvResourceInfoToDTO(resource *api.ResourceInfo) (dto.Resource, error) {
	if resource.IsDir {
		var err error
		files := make([]dto.Resource, len(resource.Dir.Resources))
		for i, r := range resource.Dir.Resources {
			files[i], err = ConvResourceInfoToDTO(r)
			if err != nil {
				return nil, err
			}
		}

		extra := make(map[string]string)
		for k, v := range resource.Meta {
			extra[k] = v
		}
		delete(extra, "status")

		return &dto.DirResource{
			ID:           resource.ID,
			Name:         resource.Name,
			FileSize:     resource.FileSize,
			FileCount:    resource.FileCount,
			ParentID:     resource.ParentID,
			Select:       resource.Meta["status"] == "1",
			SubResources: files,
			Extra:        extra,
		}, nil
	} else {
		extra := make(map[string]string)
		for k, v := range resource.Meta {
			extra[k] = v
		}
		delete(extra, "mime_type")
		delete(extra, "status")
		delete(extra, "hash")
		if resource.ParentID != "" {
			extra["parent_id"] = resource.ParentID
		}

		var fileIndex int64
		if resource.FileIndex != nil {
			fileIndex = *resource.FileIndex
		}

		return &dto.FileResource{
			ID:        resource.ID,
			Name:      resource.Name,
			FileSize:  resource.FileSize,
			FileIndex: fileIndex,
			MIMEType:  resource.Meta["mime_type"],
			Select:    resource.Meta["status"] == "1",
			Hash:      resource.Meta["hash"],
			Extra:     extra,
		}, nil
	}
}
