package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete         // We don't to actually delete the users, maybe audit if we want to hard delete them? or wait x days to purge from the table, also
	Email               string  `gorm:"not null;unique_index:idx_email"`
	UserID              *string // External user ID
	Name                *string
	NickName            *string
	FirstName           *string
	LastName            *string
	Location            *string `gorm:"size:1048"`
	Description         *string `gorm:"size:1048"`
}
