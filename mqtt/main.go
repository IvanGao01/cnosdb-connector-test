package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ivangao01/cnosdb-connector-test/mock"
)

const broker = "localhost:1883"
const clientID = "golang-mqtt-publisher"
const topic = "oceanic_station"
const username = "root"
const password = "123456"

func main() {

	client := createMqttClient()
	Publish(*client)
}

func Publish(client mqtt.Client) {

	data := make(chan string)
	go mock.Start(time.Second*10, data)

	for {
		v, ok := <-data
		if !ok {
			fmt.Println("The chan has been close!")
		}
		token := client.Publish(topic, 0, false, v)
		token.Wait()
	}

}

func createMqttClient() *mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return &client
}
