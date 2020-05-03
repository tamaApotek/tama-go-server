package user

// Role represent string definition User's Role enum
type Role string

const (
	// UserRoleAdmin represent "admin" enum
	RoleAdmin Role = "admin"
	// UserRoleDoctor represent "doctor" enum
	RoleDoctor Role = "doctor"
	// UserRolePatient represent "patient" enum
	RolePatient Role = "patient"
)

func (ur *Role) IsValid() bool {
	switch *ur {
	case RoleAdmin, RoleDoctor, RolePatient:
		return true
	default:
		return false
	}
}
