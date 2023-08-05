package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Junkes887/queues/internal/infra/kafkaConsumer"
	"github.com/Junkes887/queues/internal/infra/repository"
	"github.com/Junkes887/queues/internal/usecase"
	"github.com/Junkes887/queues/internal/utils"
	"github.com/Junkes887/queues/internal/web"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	productUseCase := usecase.NewProductUseCase(repository)

	productHandler := web.NewProductHandler(productUseCase)

	routes := chi.NewRouter()
	routes.Post("/products", productHandler.Create)
	routes.Get("/products", productHandler.List)

	go http.ListenAndServe(":8000", routes)

	msgChan := make(chan *kafka.Message)
	go kafkaConsumer.Run([]string{"products"}, "localhost:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.ProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		utils.ErrorMessage(err)

		productUseCase.Create(dto)
	}
}
