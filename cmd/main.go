package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/andreggpereira/go-kafka/infra/kafka"
	repository2 "github.com/andreggpereira/go-kafka/infra/repository"
	usercase2 "github.com/andreggpereira/go-kafka/usercase"
	_ "github.com/go-sql-driver/mysql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")
	if err != nil {
		log.Fatal(err)
	}
	repository := repository2.CourseMySQLRepository{Db: db}
	usercase := usercase2.CreateCouse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id":          "appgo",
	}

	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	for msg := range msgChan {

		var input usercase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usercase.Excute(input)
		if err != nil {
			fmt.Println("a...: Error", err)
		} else {
			fmt.Println(output)
		}

	}

}
