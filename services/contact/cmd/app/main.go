package main

import (
	"arch/pkg/store/postgres" // Импортируем пакет с функцией ConnectDB
	"arch/services/contact/internal/contact"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Параметры подключения к БД
	host := "localhost"
	port := "5432"
	user := "your_username"
	password := "your_password"
	dbname := "your_dbname"

	contactRepo := contact.NewContactRepository()

	// Создание UseCase контактов
	contactUC := contact.NewContactUseCase(contactRepo)

	// Создание доставки контактов
	contactDelivery := contact.NewContactDelivery(contactUC)

	// Настройка обработчиков маршрутов
	http.HandleFunc("/contacts", contactDelivery.HandleContacts)

	// Запуск HTTP-сервера для прослушивания REST-запросов
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Подключение к БД
	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	defer db.Close()

	fmt.Println("Подключение к БД успешно!")

}
