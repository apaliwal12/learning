package main

import (
	"testing"
)

type TestUser struct {
	Name  string `json:"name"`
	Email string `json:"email_address"`
	Age   int    // no tag
}

func TestExtractTags(t *testing.T) {
	u := TestUser{}
	tags := ExtractTags(u)

	if tags == nil {
		t.Skip("ExtractTags not implemented")
	}

	if tags["Name"] != "name" {
		t.Errorf("Expected Name tag to be 'name', got '%s'", tags["Name"])
	}
	if tags["Email"] != "email_address" {
		t.Errorf("Expected Email tag to be 'email_address', got '%s'", tags["Email"])
	}
	if tags["Age"] != "" {
		t.Errorf("Expected Age tag to be empty, got '%s'", tags["Age"])
	}

	// Test non-struct
	if ExtractTags("not a struct") != nil {
		t.Error("Expected nil when passing a non-struct")
	}
}

func TestSetStringField(t *testing.T) {
	u := TestUser{Name: "Old Name"}

	// Test success
	if !SetStringField(&u, "Name", "New Name") {
		t.Error("Expected SetStringField to return true on success")
	}
	if u.Name != "New Name" {
		t.Errorf("Expected Name to be 'New Name', got '%s'", u.Name)
	}

	// Test passing by value (should fail)
	u2 := TestUser{Name: "Old"}
	if SetStringField(u2, "Name", "New") {
		t.Error("Expected false when passing by value")
	}

	// Test setting non-existent field
	if SetStringField(&u, "InvalidField", "Value") {
		t.Error("Expected false when setting non-existent field")
	}

	// Test setting non-string field
	if SetStringField(&u, "Age", "30") {
		t.Error("Expected false when setting an int field with a string")
	}
}
