package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaReader *kafka.Reader
var kafkaWriter *kafka.Writer

func PublishToKafkaNew(topic string, body []byte, retry int) {

	if kafkaWriter == nil {
		fmt.Println(" kafkawrier is nil")
	} else {
		fmt.Println("kafka writer is \n", kafkaWriter)
	}

	messageLog := make(map[string]interface{})
	if retry < 5 {
		messageLog["retries"] = 5 - retry
		time.Sleep(100 * time.Millisecond)
	}
	messageLog["topic"] = topic
	messageLog["data"] = body

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := kafkaWriter.WriteMessages(ctx, kafka.Message{Value: body})

	fmt.Println(err)
	if err != nil {
		fmt.Println("error in publising message", err)
		if retry == 1 {
			fmt.Println("error in publising message at all")

			return
		} else {
			PublishToKafkaNew(topic, body, retry-1)
			return
		}
	} else {
		fmt.Println("##### successfully published messages")
	}
}

func InitializeKafka(topic string, conGroup string, read bool, write bool) {
	brokers := []string{
		"dev-kafka-india-1a-inst:9092",
		"dev-kafka-idnia-2a-inst:9092",
		"dev-kafka-india-3a:9092",
	}
	if read {
		readerConfig := kafka.ReaderConfig{
			Brokers:        brokers,
			GroupID:        conGroup,
			Topic:          topic,
			RetentionTime:  time.Duration(time.Hour * 168),
			CommitInterval: 0,
		}

		readerConfig.StartOffset = kafka.LastOffset

		kafkaReader = kafka.NewReader(readerConfig)

		if kafkaReader == nil {
			fmt.Println("Error: Kafka Reader initialization failed")
		} else {
			fmt.Println("Kafka Reader initialized successfully")
		}
	}

	if write {
		kafkaWriter = &kafka.Writer{
			Addr:  kafka.TCP(brokers...), // Connect to the Kafka brokers
			Topic: topic,
			Async: true,
		}
		if kafkaWriter == nil {
			fmt.Println("Error: Kafka Writer initialization failed")
		} else {
			fmt.Println("Kafka Writer initialized successfully")
		}
	}
}

type Counter struct {
	Mu    sync.Mutex
	Read  int
	Name  string
	Write int
	Error int
}

func (c *Counter) Add(typeCounter string) {
	c.Mu.Lock()
	if typeCounter == "read" {
		c.Read++
	} else if typeCounter == "write" {
		c.Write++
	} else {
		c.Error++
	}

	c.Mu.Unlock()
}

var RateStruct = Counter{Mu: sync.Mutex{}, Read: 0, Write: 0, Name: ""}
var CounterStruct = Counter{Mu: sync.Mutex{}, Read: 0, Write: 0, Name: ""}

func CompletionHandler(messages []kafka.Message, err error) {

	if err != nil {
		fmt.Println("error")

	} else {
		for range messages {
			//This is used to measure the successful publish rate
			RateStruct.Add("write")
			CounterStruct.Add("write")
		}
	}
}

func main() {
	// Define Kafka broker addresses
	// brokers := []string{
	// 	"dev-kafka-india-1a-inst:9092",
	// 	"dev-kafka-idnia-2a-inst:9092",
	// 	"dev-kafka-india-3a:9092",
	// }

	// Kafka topic
	topic := "UserUpdate_Glids"

	// // Create a Kafka writer (producer)
	// writer := kafka.Writer{
	// 	Addr:         kafka.TCP(brokers...), // Connect to the Kafka brokers
	// 	Topic:        topic,                 // Kafka topic to write to
	// 	Balancer:     &kafka.LeastBytes{},   // Use LeastBytes for load balancing
	// 	RequiredAcks: kafka.RequireAll,      // Ensure all brokers acknowledge the message
	// 	Async:        false,                 // Synchronous writing
	// }

	// defer func() {
	// 	if err := writer.Close(); err != nil {
	// 		log.Fatalf("Failed to close Kafka writer: %v", err)
	// 	}
	// }()

	// // Create a context with timeout for the write operation
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // Message to publish
	message := "Hello, Kafka using Segmentio!"
	InitializeKafka(topic, "", false, true)
	PublishToKafkaNew(topic, []byte(message), 0)
	// // // Write message to Kafka
	// err := kafkaWriter.WriteMessages(ctx, kafka.Message{
	// 	Value: []byte(message),
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to write message to Kafka: %v", err)
	// }

	fmt.Printf("Message successfully published to topic '%s': %s\n", topic, message)
	// kafkaWriter.WriteMessages()

	defer func() {
		// kafkaReader.Close()
		kafkaWriter.Close()
	}()
}
