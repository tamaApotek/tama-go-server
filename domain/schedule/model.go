package schedule

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ScheduleDay represent doctor schedule in each week day
type ScheduleDay struct {
	Day time.Weekday `json:"day" bson:"day" binding:"required,gte=0,lt=7"`
	// StartTime "minutes" + 00:00
	StartTime int `json:"start_time" bson:"start_time" binding:"gte=0"`
	// EndTime "minutes" + 00:00
	EndTime int `json:"end_time" bson:"end_time" binding:"gtfield=start_time"`
}

type Schedule struct {
	DoctorID     primitive.ObjectID `json:"doctor_id" bson:"doctor_id"`
	ScheduleDays []ScheduleDay      `json:"schedule_days" bson:"schedule_days"`
}
