package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"poc1/internal/db"

	kafka "github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

const (
	topic         = "purchase-data"
	brokerAddress = "localhost:9092"
)

type Tabler interface {
	TablePurchase() string
}

func (purchase) TablePurchase() string {
	return "purchase"
}

type purchase struct {
	ID           int `gorm:"primaryKey"`
	CompanyName  string
	CompanyCNPJ  string
	Value        float64
	CustomerName string
	CustomerCPF  string
}

func main() {

	dbConn = db.OpenConnection()

	ctx := context.Background()

	go produce(ctx)
	consume(ctx)

}

func produce(ctx context.Context) {

	//var dbconn *gorm.DB
	//conn = db.OpenConnection()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokerAddress},
		Topic:        topic,
		BatchSize:    10,
		BatchTimeout: time.Millisecond,
	})

	for i := 1; i <= 100; i++ {

		p := purchase{
			CompanyName:  "Empresa ABC",
			CompanyCNPJ:  "43.387.863/0001-76",
			Value:        55.0,
			CustomerName: "José da Silva",
			CustomerCPF:  "981.064.590-20",
		}

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			Value: []byte(strconv.Itoa(i) + "ª mensagem recebida : " +
				p.CompanyCNPJ + " | " + p.CompanyName + " | " + p.CustomerCPF + " | " + p.CustomerName),
		})
		if err != nil {
			panic("Não pôde escrever a mensagem " + err.Error())
		}

		fmt.Printf("%dª mensagem enviada: %s - %s - %s - %s\n", i, p.CompanyCNPJ, p.CompanyName, p.CustomerCPF, p.CustomerName)
		dbConn.Create(&p)

	}

	time.Sleep(time.Second)
}

func consume(ctx context.Context) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "group-purchase-data",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Não pôde ler a mensagem " + err.Error())
		}
		//dbConn.Create(&c)
		fmt.Println(string(msg.Value))
	}
}
