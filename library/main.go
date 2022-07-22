package library

import "kafka-exception-iterator/pkg/config"

func main() {
	//config load
	configInstance := config.CreateConfigInstance()
	applicationConfig, err := configInstance.GetConfig()
	if err != nil {
		panic("application config read failed: " + err.Error())
	}

	var produceFn produceExceptionFn = func(message Message) error {
		return nil
	}
	var consumeFn consumeExceptionFn = func() (Message, error) {

		return sampleConsumer.ReadMessage
	}

	manager := newExceptionManager(produceFn, consumeFn, applicationConfig.Kafka.Consumer.ExceptionTopic)

}
