package models

// User defined a user for the app
type User struct {
	// We don't need to delete users,
	// maybe audit if we want to delete
	// them, or wait x days to purge from
	// the table, also
	BaseModelSoftdelete
	Email       string  `gorm:"not null;unique_index:idx_email"`
	UserID      *string // External user ID
	Name        *string
	NickName    *string
	FirstName   *string
	LastName    *string
	Location    *string `gorm:"size:1048"`
	Description *string `gorm:"size:1048"`
}
