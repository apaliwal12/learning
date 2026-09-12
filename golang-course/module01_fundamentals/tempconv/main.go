package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Unit represents a temperature unit using iota
type Unit int

const (
	Celsius Unit = iota
	Fahrenheit
	Kelvin
)

// String makes Unit implement the fmt.Stringer interface
func (u Unit) String() string {
	switch u {
	case Celsius:
		return "Celsius"
	case Fahrenheit:
		return "Fahrenheit"
	case Kelvin:
		return "Kelvin"
	default:
		return "Unknown"
	}
}

// Convert converts a temperature from one unit to another
func Convert(value float64, from, to Unit) (float64, error) {
	if from == to {
		return value, nil
	}

	// First convert to Celsius
	var c float64
	switch from {
	case Celsius:
		c = value
	case Fahrenheit:
		c = (value - 32) * 5 / 9
	case Kelvin:
		c = value - 273.15
	default:
		return 0, fmt.Errorf("unknown source unit: %d", from)
	}

	// Then convert from Celsius to target
	switch to {
	case Celsius:
		return c, nil
	case Fahrenheit:
		return (c * 9 / 5) + 32, nil
	case Kelvin:
		return c + 273.15, nil
	default:
		return 0, fmt.Errorf("unknown target unit: %d", to)
	}
}

func parseUnit(s string) (Unit, error) {
	switch strings.ToUpper(s) {
	case "C", "CELSIUS":
		return Celsius, nil
	case "F", "FAHRENHEIT":
		return Fahrenheit, nil
	case "K", "KELVIN":
		return Kelvin, nil
	default:
		return 0, fmt.Errorf("invalid unit '%s' (use C, F, or K)", s)
	}
}

func interactiveMode() {
	fmt.Println("--- Temperature Converter (Interactive Mode) ---")
	fmt.Println("Enter 'q' to quit.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter value and unit (e.g. 100 C): ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if strings.ToLower(input) == "q" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Error: Expected format '<value> <unit>'")
			continue
		}

		valStr, unitStr := parts[0], parts[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Printf("Error: Invalid number '%s'\n", valStr)
			continue
		}

		unit, err := parseUnit(unitStr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("\n%.2f %s converts to:\n", val, unit)
		for _, target := range []Unit{Celsius, Fahrenheit, Kelvin} {
			if target != unit {
				converted, _ := Convert(val, unit, target)
				fmt.Printf(" - %.2f %s\n", converted, target)
			}
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--interactive" {
			interactiveMode()
			return
		}
		fmt.Println("Usage: run without arguments or with '--interactive'")
		os.Exit(1)
	}

	// Demo non-interactive
	fmt.Println("--- Temperature Converter Demo ---")
	val := 100.0
	unit := Celsius

	fmt.Printf("%.2f %s converts to:\n", val, unit)
	for _, target := range []Unit{Fahrenheit, Kelvin} {
		converted, err := Convert(val, unit, target)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf(" - %.2f %s\n", converted, target)
		}
	}
	fmt.Println("\nRun with '--interactive' to enter your own values.")
}
