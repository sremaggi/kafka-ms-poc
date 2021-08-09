package models


//Request POST REST
type Request struct  {
	Topic string 	`json:"topic"`
	Brokers string `json:"brokers"`
	ProducerMessage *producerMesssage `json:"producerMessagge"`


}
type producerMesssage struct {
	Headers header `json:"headers"`
	Key string `json:"key"`
	Message string`json:"messagge"`
}


type header struct {
	Key   string`json:"key"` // Header name (utf-8 string)
	Value string `json:"value"` // Header value (nil, empty, or binary)
}