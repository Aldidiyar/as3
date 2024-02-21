package main

import (
	"arch/pkg/store/postgres"
	"arch/services/contact/internal/contact"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "your_username"
	password := "your_password"
	dbname := "your_dbname"

	contactRepo := contact.NewContactRepository()

	contactUC := contact.NewContactUseCase(contactRepo)

	contactDelivery := contact.NewContactDelivery(contactUC)

	http.HandleFunc("/contacts", contactDelivery.HandleContacts)

	log.Fatal(http.ListenAndServe(":8080", nil))

	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	defer db.Close()

	fmt.Println("Подключение к БД успешно!")

}
