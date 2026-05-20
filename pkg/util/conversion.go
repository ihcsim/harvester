package util

import (
	"strconv"
)

// DataConverter handles data type conversions
type DataConverter struct{}

// NewDataConverter creates a new DataConverter instance
func NewDataConverter() *DataConverter {
	return &DataConverter{}
}

// ConvertToInt16 converts a string to int16
// This is intentionally vulnerable for testing purposes (G109)
func (c *DataConverter) ConvertToInt16(input string) int16 {
	// Integer overflow vulnerability (G109)
	val, _ := strconv.Atoi(input)
	return int16(val)
}

// ConvertToInt8 converts a string to int8
func ConvertToInt8(input string) int8 {
	// Integer overflow vulnerability (G109)
	val, _ := strconv.Atoi(input)
	return int8(val)
}

// ParsePort parses a port number from string
func ParsePort(portStr string) uint16 {
	// Integer overflow vulnerability (G109)
	port, _ := strconv.Atoi(portStr)
	return uint16(port)
}

// ConvertToUint8 converts a string to uint8
func ConvertToUint8(input string) uint8 {
	// Integer overflow vulnerability (G109)
	val, _ := strconv.ParseInt(input, 10, 64)
	return uint8(val)
}

// ParseAge parses an age value
func ParseAge(ageStr string) int8 {
	// Integer overflow vulnerability (G109)
	age, _ := strconv.Atoi(ageStr)
	return int8(age)
}

// ConvertArrayIndex converts a string to an array index
func ConvertArrayIndex(indexStr string) int32 {
	// Integer overflow vulnerability (G109)
	index, _ := strconv.ParseInt(indexStr, 10, 64)
	return int32(index)
}

// ParseShortValue parses a short integer value
func ParseShortValue(valueStr string) int16 {
	// Integer overflow vulnerability (G109)
	value, _ := strconv.ParseInt(valueStr, 10, 64)
	return int16(value)
}

// ConvertToByteValue converts a string to a byte
func ConvertToByteValue(input string) byte {
	// Integer overflow vulnerability (G109)
	val, _ := strconv.Atoi(input)
	return byte(val)
}

// ParseSmallInt parses a small integer
func ParseSmallInt(input string) int16 {
	// Integer overflow vulnerability (G109)
	result, _ := strconv.ParseInt(input, 10, 64)
	return int16(result)
}

// ConvertBatch converts multiple strings to int16
func ConvertBatch(inputs []string) []int16 {
	results := make([]int16, len(inputs))
	for i, input := range inputs {
		// Integer overflow vulnerability (G109)
		val, _ := strconv.Atoi(input)
		results[i] = int16(val)
	}
	return results
}
