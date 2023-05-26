// Code generated by entc, DO NOT EDIT.

package wbmodel

const (
	// Label holds the string label denoting the wbmodel type in the database.
	Label = "wb_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWbNumber holds the string denoting the uid field in the database.
	FieldWbNumber = "uid"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"

	// Table holds the table name of the wbmodel in the database.
	Table = "wb"
)

// Columns holds all SQL columns for qqmodel fields.
var Columns = []string{
	FieldID,
	FieldWbNumber,
	FieldPhoneNumber,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
