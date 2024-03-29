package mqttHandler

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	IsLightOn bool
	Place     string
}

const (
	TOPIC     = "leds"
	QOS       = 1
	CLIENT_ID = "cronTrigger"
)

var SERVER_ADDRESS = os.Getenv("MQTT_SERVER_ADDRESS")

func PublishMessage(m string) {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRITICAL] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	// mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(SERVER_ADDRESS)
	opts.SetClientID(CLIENT_ID)
	opts.SetOrderMatters(false)       // Allow out of order messages (use this option unless in order delivery is essential)
	opts.ConnectTimeout = time.Second // Minimal delays on connect
	opts.WriteTimeout = time.Second   // Minimal delays on writes
	opts.KeepAlive = 10               // Keepalive every 10 seconds so we quickly detect network outages
	opts.PingTimeout = time.Second    // local broker so response should be quick
	opts.ConnectRetry = true
	opts.AutoReconnect = true
	opts.DefaultPublishHandler = func(_ mqtt.Client, msg mqtt.Message) {
		fmt.Printf("type=debug tag=mqtt msg=\"UNEXPECTED MESSAGE: %s\"\n", msg)
	}
	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		fmt.Println("type=debug tag=mqtt msg=\"connection lost\"")
	}
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("type=debug tag=mqtt msg=\"connection established\"")
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		fmt.Println("type=debug tag=mqtt msg=\"attempting to reconnect\"")
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("MQTT connection from cronTrigger is up")

	if os.Getenv("SHOULD_TRIGGER_ALEXA") == "false" {
		log.Println("Dev mode: Alexa trigger not send, but invoked, to topic: " + TOPIC)
		log.Println("Dev mode: Payload: " + m)
		client.Disconnect(1000)
		return
	}

	t := client.Publish(TOPIC, QOS, false, m)
	_ = t.Wait()
	if t.Error() != nil {
		log.Printf("type=error tag=mqtt msg=\"ERROR Publishing: %s\"\n", t.Error())
		panic(t.Error())
	} else {
		log.Printf("type=success tag=mqtt action=publish msg=\"%s\"\n", m)
	}
	client.Disconnect(1000)
	log.Println("type=debug tag=mqtt msg=\"Client disconnected\"")
}
