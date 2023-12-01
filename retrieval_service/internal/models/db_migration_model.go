package models

// User [...]
type SchemaUser struct {
	ID       string `gorm:"column:id;type:varchar(100);primary_key" json:"id"`
	Username string `gorm:"column:username;type:varchar(100)" json:"username"`
	Password string `gorm:"column:password;type:varchar(500)" json:"password"`
}

// TableName get sql table name
func (m *SchemaUser) TableName() string {
	return "user"
}

// UserColumns get sql column name
var UserColumns = struct {
	ID       string
	Username string
	Password string
}{
	ID:       "id",
	Username: "username",
	Password: "password",
}
