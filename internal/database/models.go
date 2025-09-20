package database

import "database/sql"

// Models - это структура-контейнер, которая хранит ВСЕ модели нашей базы данных
type Models struct {
	User      UserModel
	Events    EventModel
	Attendees AttendeeModel
}

// NewModels - функция-конструктор. Она создает и возвращает экземпляр Models

func NewModels(db *sql.DB) Models {
	return Models{
		User:      UserModel{DB: db},
		Events:    EventModel{DB: db},
		Attendees: AttendeeModel{DB: db},
	}
}
