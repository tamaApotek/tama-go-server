package role

// Role represent string definition User's Role enum
type Role string

const (
	// UserRoleAdmin represent "admin" enum
	Admin Role = "admin"
	// UserRoleDoctor represent "doctor" enum
	Doctor Role = "doctor"
	// UserRolePatient represent "patient" enum
	Patient Role = "patient"
)

func (ur *Role) IsValid() bool {
	switch *ur {
	case Admin, Doctor, Patient:
		return true
	default:
		return false
	}
}
