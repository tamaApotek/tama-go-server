package constants

// UserRole represent string alias of UserRole enum
type UserRole string

const (
	// UserRoleAdmin represent "admin" enum
	UserRoleAdmin UserRole = "admin"
	// UserRoleDoctor represent "doctor" enum
	UserRoleDoctor = "doctor"
	// UserRolePatient represent "patient" enum
	UserRolePatient = "patient"
)
