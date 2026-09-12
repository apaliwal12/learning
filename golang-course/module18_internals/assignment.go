package main

// TODO 1: Implement ExtractTags
// It should take ANY struct instance (passed as an interface{}).
// It should return a map[string]string where the key is the field name,
// and the value is the "json" struct tag.
// If the input is not a struct, return nil.
// If a field doesn't have a json tag, its value in the map should be an empty string.
func ExtractTags(obj interface{}) map[string]string {
	return nil // Fix me
}

// TODO 2: Implement SetStringField
// It should take a POINTER to a struct (passed as interface{}),
// a field name (string), and a new value (string).
// It should use reflection to update the struct field if possible.
// Return true if it successfully updated, false otherwise.
func SetStringField(obj interface{}, fieldName string, value string) bool {
	return false // Fix me
}
