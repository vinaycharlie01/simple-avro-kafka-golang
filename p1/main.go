// package main

// import (
// 	"log"

// 	kingpin "github.com/alecthomas/kingpin/v2"

// 	"github.com/IBM/sarama"
// )

// var (
// 	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
// 	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
// 	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
// )

// func main() {
// 	kingpin.Parse()
// 	config := sarama.NewConfig()
// 	config.Producer.RequiredAcks = sarama.WaitForAll
// 	config.Producer.Retry.Max = *maxRetry
// 	config.Producer.Return.Successes = true
// 	producer, err := sarama.NewSyncProducer(*brokerList, config)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer func() {
// 		if err := producer.Close(); err != nil {
// 			log.Panic(err)
// 		}
// 	}()
// 	msg := &sarama.ProducerMessage{
// 		Topic: *topic,
// 		Value: sarama.StringEncoder("Something Cool"),
// 	}
// 	partition, offset, err := producer.SendMessage(msg)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
// }


// to produce messages
package main


import (
	"context"
	"log"
	"time"
	"github.com/segmentio/kafka-go"
)

func main(){
		// to produce messages
		topic := "my-topic"
		partition := 0
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
                 log.Fatal("failed to dial leader:", err)
				}

				conn.SetWriteDeadline(time.Now().Add(10*time.Second))
_, err = conn.WriteMessages(
    kafka.Message{Value: []byte("one!")},
    kafka.Message{Value: []byte("two!")},
    kafka.Message{Value: []byte("three!")},
)
if err != nil {
    log.Fatal("failed to write messages:", err)
}

if err := conn.Close(); err != nil {
    log.Fatal("failed to close writer:", err)
}

}