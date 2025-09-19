package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file" // ★ Важно: этот импорт необходим!
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	// открываем соединение с базой данных SQLite. Файл базы будет называться data.db и лежать в текущей папке.
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Создаем специальный "драйвер" для работы библиотеки миграций с нашей SQLite-базой. По сути, это адаптер.
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// ★ Правильный способ создания мигратора
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // URL к миграциям
		"sqlite3",                       // имя драйвера БД
		instance,
	)
	if err != nil {
		log.Fatal(err)
	}

	// ★ Методы должны быть с большой буквы!
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration UP completed successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration DOWN completed successfully")
	default:
		log.Fatal("Unsupported migration direction. Use 'up' or 'down'")
	}
}
