package main

// User represents an item from the "users" table, column
// names are mapped to Go values.
type User struct {
	// Map the "id" column to the ID field. Only exported
	// fields can be mapped to database columns.
	ID uint `db:"id"`
	// The "created_at" column is a VARCHAR type, upper-db converts
	// Go types into database-specific types and vice versa.
	CreatedAt string `db:"created_at"`
	// The "author_id" column is a VARCHAR type.
	UpdatedAt string `db:"update_at"`
	// The "subject_id" column is a VARCHAR type.
	DeletedAt string `db:"deleted_at"`
	// The "name" column is a VARCHAR type.
	Name string `db:"name"`
	// The "email" column is a VARCHAR type.
	Email string `db:"email"`
	// The "mobile" column is a VARCHAR type.
	Mobile string `db:"mobile"`
	// The "password" column is a VARCHAR type.
	Password string `db:"password"`
	// The "status" column is a VARCHAR type.
	Status string `db:"status"`
	// The "status" column is a VARCHAR type.
	Type string `db:"type"`
	// The "remember_token" column is a VARCHAR type.
	RememberToken string `db:"remember_token"`
	// The "sms_token" column is an integer type.
	SmsToken uint `db:"sms_token"`
}
