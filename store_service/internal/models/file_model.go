package models

type FileGorm struct {
	ID        string `gorm:"id"`
	Name      string `gorm:"name"`
	Type      string `gorm:"type"`
	FilePath  string `gorm:"file_path"`
	CreatedAt string `gorm:"created_at"`
}

type TagGorm struct {
	ID    string `gorm:"id"`
	Value string `gorm:"value"`
}

type FileTagGorm struct {
	ID     string `gorm:"id"`
	TagId  string `gorm:"tag_id"`
	FileId string `gorm:"file_id"`
}

type RequestStore struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

type RequestRetreive struct {
	Name string   `json:"name"`
	Tag  []string `json:"tag"`
}
