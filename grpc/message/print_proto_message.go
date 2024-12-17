package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	print_message()
}

func print_message() {
	// Example ProtoMessage
	msg, err := anypb.New(wrapperspb.String("example value"))
	if err != nil {
		log.Fatalf("Failed to create Any message: %v", err)
	}

	// Marshal the ProtoMessage to JSON
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal proto message to JSON: %v", err)
	}

	// Print the JSON string
	fmt.Println(string(jsonBytes))

	// {"@type":"type.googleapis.com/google.protobuf.StringValue", "value":"example value"}
}
