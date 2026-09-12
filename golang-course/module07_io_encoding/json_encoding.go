package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ─────── 3. ENCODING (JSON) ───────
// Go's standard library `encoding/json` uses reflection to serialize and deserialize data.
// 
// Two main approaches:
// 1. Marshal / Unmarshal: Operates on []byte (in-memory).
// 2. Encoder / Decoder: Operates on io.Writer / io.Reader (streams).

type Config struct {
	ServerName string `json:"server"`
	Port       int    `json:"port"`
	// Use pointers for fields that might be absent in the JSON,
	// so you can differentiate between "0" and "null/missing".
	Timeout *int `json:"timeout,omitempty"`
}

func demonstrateJSON() {
	fmt.Println("\n--- JSON Encoding ---")

	// 1. Marshal (Struct to []byte)
	cfg := Config{
		ServerName: "api.example.com",
		Port:       8080,
	}

	// json.MarshalIndent makes it pretty-printed. For production, use json.Marshal.
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Printf("Marshaled JSON:\n%s\n", data)

	// 2. Unmarshal ([]byte to Struct)
	jsonStr := `{"server": "db.local", "port": 5432, "timeout": 30}`
	var newCfg Config
	
	// WARNING: You must pass a POINTER to the struct so Unmarshal can modify it!
	err = json.Unmarshal([]byte(jsonStr), &newCfg)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	
	fmt.Printf("Unmarshaled Struct: %+v\n", newCfg)
	if newCfg.Timeout != nil {
		fmt.Printf("Timeout value: %d\n", *newCfg.Timeout)
	}

	// 3. Encoder / Decoder (Streaming)
	// This is highly efficient for reading from/writing to files or network sockets
	// because it doesn't need to load the entire JSON string into memory first.
	
	fmt.Println("\nStreaming JSON:")
	var buf bytes.Buffer
	
	// Create an Encoder that writes to the buffer (could be a file or http.ResponseWriter!)
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(cfg) // Encodes and writes in one step
	
	fmt.Printf("Stream encoded: %s", buf.String())
	
	// Create a Decoder that reads from the buffer
	decoder := json.NewDecoder(&buf)
	var streamCfg Config
	err = decoder.Decode(&streamCfg) // Reads and decodes in one step
	
	fmt.Printf("Stream decoded: %+v\n", streamCfg)
}
