package models

// Storage [...]
type SchemaStorage struct {
	ID        string `gorm:"column:id;type:varchar(100);primary_key" json:"id"`
	Name      string `gorm:"column:name;type:varchar(100)" json:"name"`
	Type      string `gorm:"column:type;type:varchar(100)" json:"type"`
	FilePath  string `gorm:"column:file_path;type:varchar(100)" json:"file_path"`
	CreatedAt string `gorm:"column:created_at;type:varchar(100)" json:"created_at"`
}

// TableName get sql table name
func (m *SchemaStorage) TableName() string {
	return "storage"
}

// StorageColumns get sql column name
var StorageColumns = struct {
	ID        string
	Name      string
	Type      string
	FilePath  string
	CreatedAt string
}{
	ID:        "id",
	Name:      "name",
	Type:      "type",
	FilePath:  "file_path",
	CreatedAt: "created_at",
}

// Tag [...]
type SchemaTag struct {
	ID    string `gorm:"column:id;type:varchar(100);primary_key" json:"id"`
	Value string `gorm:"column:value;type:varchar(100)" json:"value"`
}

// TableName get sql table name
func (m *SchemaTag) TableName() string {
	return "tag"
}

// TagColumns get sql column name
var TagColumns = struct {
	ID    string
	Value string
}{
	ID:    "id",
	Value: "value",
}

// StoreTag [...]
type SchemaStoreTag struct {
	ID      string        `gorm:"column:id;type:varchar(100);primary_key" json:"id"`
	TagId   string        `gorm:"column:tag_id;type:varchar(100)" json:"tag_id"`
	FileId  string        `gorm:"column:file_id;type:varchar(100)" json:"file_id"`
	Tag     SchemaTag     `gorm:"joinForeignKey:tag_id;foreignKey:TagId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tag"`
	Storage SchemaStorage `gorm:"joinForeignKey:file_id;foreignKey:FileId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"storage"`
}

// TableName get sql table name
func (m *SchemaStoreTag) TableName() string {
	return "store_tag"
}

// StoreTagColumns get sql column name
var StoreTagColumns = struct {
	ID     string
	TagId  string
	FileId string
}{
	ID:     "id",
	TagId:  "tag_id",
	FileId: "file_id",
}
