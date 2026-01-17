package dbgivent

import (
	"database-repository/config"
	"database-repository/model/dbGivent"
)

func GetEvent() ([]dbgivent.Event, error, string) {
	var event []dbgivent.Event

	nama := "Event Table"

	result := config.DBGivent.Table("event").
		Find(&event)

	return event, result.Error, nama
}
