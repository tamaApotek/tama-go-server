package constants

// UserRole represent string alias of UserRole enum
type UserRole string

const (
	// UserRoleAdmin represent "admin" enum
	UserRoleAdmin UserRole = "admin"
	// UserRoleDoctor represent "doctor" enum
	UserRoleDoctor UserRole = "doctor"
	// UserRolePatient represent "patient" enum
	UserRolePatient UserRole = "patient"
)

func (ur *UserRole) IsValid() bool {
	switch *ur {
	case UserRoleAdmin, UserRoleDoctor, UserRolePatient:
		return true
	default:
		return false
	}
}
