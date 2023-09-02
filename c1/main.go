// package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"

// 	kingpin "github.com/alecthomas/kingpin/v2"

// 	"github.com/IBM/sarama"
// )

// var (
// 	brokerList        = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
// 	topic             = kingpin.Flag("topic", "Topic name").Default("important").String()
// 	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
// 	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
// 	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
// )

// func main() {
// 	kingpin.Parse()
// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true
// 	brokers := *brokerList
// 	master, err := sarama.NewConsumer(brokers, config)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer func() {
// 		if err := master.Close(); err != nil {
// 			log.Panic(err)
// 		}
// 	}()
// 	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, os.Interrupt)
// 	doneCh := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case err := <-consumer.Errors():
// 				log.Println(err)
// 			case msg := <-consumer.Messages():
// 				*messageCountStart++
// 				log.Println("Received messages", string(msg.Key), string(msg.Value))
// 			case <-signals:
// 				log.Println("Interrupt is detected")
// 				doneCh <- struct{}{}
// 			}
// 		}
// 	}()
// 	<-doneCh
// 	log.Println("Processed", *messageCountStart, "messages")
// }



package main


import(
	"context"
	"log"
	"time"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main(){
	
topic := "my-topic"
partition := 0

conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
if err != nil {
    log.Fatal("failed to dial leader:", err)
}

conn.SetReadDeadline(time.Now().Add(10*time.Second))
batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

b := make([]byte, 10e3) // 10KB max per message
for {
    n, err := batch.Read(b)
    if err != nil {
        break
    }
    fmt.Println(string(b[:n]))
}

if err := batch.Close(); err != nil {
    log.Fatal("failed to close batch:", err)
}

if err := conn.Close(); err != nil {
    log.Fatal("failed to close connection:", err)
}
}
