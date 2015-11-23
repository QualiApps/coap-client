package main

import (
	"github.com/dustin/go-coap"
	"log"
)

const (
	host = "localhost"
	port = "5683"
)

func main() {

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 12345,
		//Payload:   []byte("{\"title\": \"ABCD\", \"desc\": \"DESCRIPTION\"}"),
	}

	req.SetOption(coap.ETag, "weetag")
	req.SetOption(coap.MaxAge, 3)
	req.SetPathString("/api-db")

	c, err := coap.Dial("udp", "localhost:5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		log.Printf("Response payload: %#v", rv)
		log.Printf("Payload is: %s", rv.Payload)
	}

}
