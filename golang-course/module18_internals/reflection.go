package main

import (
	"fmt"
	"reflect"
)

// ─────── 3. REFLECTION ───────
// Reflection allows a program to inspect its own variables and types AT RUNTIME.
// It is incredibly powerful but DANGEROUS. It bypasses compile-time type checking,
// is significantly slower than standard code, and can cause panics if used incorrectly.
//
// Rule of Thumb: "Clear is better than clever. Reflection is rarely clear."
// Use cases: JSON marshalling/unmarshalling, ORMs (like GORM), deeply generic utilities.

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func demonstrateReflection() {
	fmt.Println("\n--- Reflection ---")
	
	cfg := Config{Host: "localhost", Port: 8080}
	
	// 1. reflect.TypeOf
	// Gets the Type information of an interface{}
	t := reflect.TypeOf(cfg)
	fmt.Println("Type:", t.Name()) // "Config"
	
	// 2. reflect.ValueOf
	// Gets the actual runtime Value of an interface{}
	v := reflect.ValueOf(cfg)
	
	// Iterate over the fields of the struct!
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)      // StructField (type info, tags, etc.)
		val := v.Field(i)        // Value (the actual data)
		
		fmt.Printf("Field: %s, Type: %s, Value: %v, Tag(json): %s\n", 
			field.Name, field.Type, val, field.Tag.Get("json"))
	}
	
	// 3. Modifying values via Reflection
	// To modify a value, you MUST pass a pointer to reflect.ValueOf,
	// and then call .Elem() to dereference the pointer and get the settable value.
	fmt.Println("\nModifying values...")
	
	// We MUST pass &cfg, not cfg!
	vPtr := reflect.ValueOf(&cfg)
	vSettable := vPtr.Elem()
	
	// Get the "Port" field
	portField := vSettable.FieldByName("Port")
	
	// Verify we can set it and it's an int
	if portField.CanSet() && portField.Kind() == reflect.Int {
		portField.SetInt(9090)
		fmt.Println("Successfully changed Port via reflection!")
	}
	
	fmt.Println("Updated Config:", cfg)
}
