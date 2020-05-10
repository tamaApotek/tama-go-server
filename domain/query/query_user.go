package query

type User struct {
	FullName    *string `json:"full_name" bson:"full_name"`
	Email       *string `json:"email" bson:"email"`
	PhoneNumber *string `json:"phone_number" bson:"phone_number"`
}
