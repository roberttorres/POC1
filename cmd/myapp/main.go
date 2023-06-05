package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "purchase-data"
	brokerAddress = "localhost:9092"
)

type purchaseData struct {
	companyName  string
	companyCNPJ  string
	value        float64
	customerName string
	customerCPF  string
}

func main() {

	ctx := context.Background()

	go produce(ctx)
	consume(ctx)

}

func produce(ctx context.Context) {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokerAddress},
		Topic:        topic,
		BatchSize:    10,
		BatchTimeout: time.Millisecond,
	})

	p := purchaseData{
		companyName:  "Empresa ABC",
		companyCNPJ:  "43.387.863/0001-76",
		value:        55.0,
		customerName: "José da Silva",
		customerCPF:  "981.064.590-20",
	}

	for i := 1; i <= 100; i++ {

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			Value: []byte(strconv.Itoa(i) + "ª mensagem recebida : " +
				p.companyCNPJ + " | " + p.companyName + " | " + p.customerCPF + " | " + p.customerName),
		})
		if err != nil {
			panic("Não pôde escrever a mensagem " + err.Error())
		}

		fmt.Printf("%dª mensagem enviada: %s - %s - %s - %s\n", i, p.companyCNPJ, p.companyName, p.customerCPF, p.customerName)

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
		fmt.Println(string(msg.Value))
	}
}
