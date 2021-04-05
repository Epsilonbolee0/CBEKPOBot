package users

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	ID              string    `gorm:"primaryKey"`
	Name            string    `gorm:"default:Дорогуша"`
	MessageFromTime time.Time `gorm:"default:'2001-10-11 12:00+00'"`
	MessageToTime   time.Time `gorm:"default:'2001-10-11 18:00+00'"`
}

func (user *User) String() string {
	builder := strings.Builder{}
	dateFormat := "15:04:05"

	builder.WriteString("Id => " + user.ID + "; ")
	builder.WriteString("Name => " + user.Name + "; ")
	builder.WriteString("FromTime => " + user.MessageFromTime.Format(dateFormat) + "; ")
	builder.WriteString("ToTime => " + user.MessageToTime.Format(dateFormat))

	return builder.String()
}
