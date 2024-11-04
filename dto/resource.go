package dto

type Resource interface {
	IsDir() bool
	GetID() string
	GetName() string
	GetFileSize() int64
	IsSelected() bool
	GetFileCount() int64
	GetFiles() []*FileResource
	GetExtra() map[string]string
}

type DirResource struct {
	ID           string
	Name         string
	FileSize     int64
	FileCount    int64
	ParentID     string
	Select       bool
	SubResources []Resource
	Extra        map[string]string
}

func (r *DirResource) IsDir() bool         { return true }
func (r *DirResource) GetID() string       { return r.ID }
func (r *DirResource) GetName() string     { return r.Name }
func (r *DirResource) GetFileSize() int64  { return r.FileSize }
func (r *DirResource) IsSelected() bool    { return r.Select }
func (r *DirResource) GetFileCount() int64 { return r.FileCount }
func (r *DirResource) GetFiles() []*FileResource {
	files := make([]*FileResource, 0, r.GetFileCount())
	for _, subRes := range r.SubResources {
		files = append(files, subRes.GetFiles()...)
	}
	return files
}
func (r *DirResource) GetExtra() map[string]string { return r.Extra }

type FileResource struct {
	ID        string
	Name      string
	FileSize  int64
	FileIndex int64
	MIMEType  string
	Select    bool
	Hash      string
	Extra     map[string]string
}

func (r *FileResource) IsDir() bool                 { return false }
func (r *FileResource) GetID() string               { return r.ID }
func (r *FileResource) GetName() string             { return r.Name }
func (r *FileResource) GetFileSize() int64          { return r.FileSize }
func (r *FileResource) IsSelected() bool            { return r.Select }
func (r *FileResource) GetFileCount() int64         { return 1 }
func (r *FileResource) GetFiles() []*FileResource   { return []*FileResource{r} }
func (r *FileResource) GetExtra() map[string]string { return r.Extra }
