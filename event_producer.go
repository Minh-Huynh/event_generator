package main

import (
  "fmt"
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
  "time"
  "math/rand"
  "encoding/xml"
	logrus "github.com/sirupsen/logrus"
  "strconv"
  "sync"
)

const (
  TOPIC = "eew/sys/gm-contour/data"
  QOS_CLIENT_TO_BROKER = 1
)

var BROKERS = [6]string { "tcp://localhost:1883", "tcp://localhost:1884",
                          "tcp://localhost:1885", "tcp://localhost:1886",
                          "tcp://localhost:1887", "tcp://localhost:1888",}

var ALERT_SERVER_HOSTNAMES = [6]string { "eew-ci-prod1", "eew-ci-prod2",
                                         "eew-nc-prod1","eew-nc-prod2",
                                         "eew-uw-prod1","eew-uw-prod2",}
func init() {
 logrus.SetOutput(os.Stdout)
}

//check eventMessage.Instance[0] (hostname:eew-ci-prod1/2, eew-nc-prod1/2, eew-uw-prod1/2),
//eventMessage.MessageType (should be 'new'), eventMessage.CoreInfo.ID,
//eventMessage.CoreInfo.OrigTime.Text is used to create primary key
//To generate time: time.Now().Format(time.RFC3339Nano)
//you will need to xml.Marshal the eventMessage

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
  opts := MQTT.NewClientOptions()
  for i := 0; i < len(BROKERS); i++ {
    logrus.Info("Adding Broker " + BROKERS[i])
    opts.AddBroker(BROKERS[i])
  }

  client := MQTT.NewClient(opts)

  if token := client.Connect(); token.Wait() && token.Error() != nil {
    logrus.Info(token.Error())
    panic(token.Error())
  }

  //subscribe to the topic 
  logrus.Info("Subscribing to " + TOPIC)
  if token := client.Subscribe(TOPIC,QOS_CLIENT_TO_BROKER, nil); token.Wait() && token.Error() != nil {
    os.Exit(1)
  }

  //generate eventMessage
  var wg sync.WaitGroup
  wg.Add(len(ALERT_SERVER_HOSTNAMES))
//Publish all events a second apart
  for i := 0; i < len(ALERT_SERVER_HOSTNAMES); i++ {
    go func(i int){
      defer wg.Done()
      timeString := GenerateTimeStampString(time.Now().Add(time.Duration(i * 2) * time.Second))
      logrus.Info("Index: ", i)
      logrus.Info("Setting up TimeStamp: ", timeString)
      msg := GenerateEventMsg(ALERT_SERVER_HOSTNAMES[i], "new", timeString )
      dataToSend , _ := xml.Marshal(msg) 
      logrus.Info("Generating message for " + ALERT_SERVER_HOSTNAMES[i] + "\n " )

      //Publish eventMessage
      logrus.Info("Publishing message... " + string(dataToSend) + "\n")
      token := client.Publish(TOPIC, 1, false, dataToSend)
      token.Wait()
    }(i)
  }
  wg.Wait()

  //unsubscribe from topic
  logrus.Info("Unsubscribing from " + TOPIC)
  if token := client.Unsubscribe(TOPIC); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }

  logrus.Info("Disconnecting...")
  client.Disconnect(250)
}
//check eventMessage.Instance (hostname:eew-ci-prod1/2, eew-nc-prod1/2, eew-uw-prod1/2),
//eventMessage.MessageType (should be 'new'), 
//eventMessage.CoreInfo.ID, this is unique ID for the particular alert server
//eventMessage.CoreInfo.OrigTime.Text is used to create primary key based on time
//To generate time: time.Now().Format(time.RFC3339Nano)
//you will need to xml.Marshal the eventMessage

func GenerateTimeStampString(t time.Time) string {
  return t.Format(time.RFC3339Nano)
}



func GenerateEventMsg(hostname string, msgType string, timeStamp string) EventMessage {
  uniqueID := rand.Intn(1200)

  message := EventMessage{}
  message.Instance = hostname
  message.MessageType = msgType
  message.CoreInfo.ID = strconv.Itoa(uniqueID)
  message.CoreInfo.OrigTime.Text = timeStamp

  return message
}
