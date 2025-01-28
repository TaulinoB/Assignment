package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rotor represents a single rotor in the Enigma machine
type Rotor struct {
	wiring   string
	position int
}

// Reflector represents the reflector in the Enigma machine
type Reflector struct {
	wiring string
}

// NewRotor creates a new rotor with a given wiring and initial position
func NewRotor(wiring string, position int) *Rotor {
	return &Rotor{wiring: wiring, position: position}
}

// NewReflector creates a new reflector with a given wiring
func NewReflector(wiring string) *Reflector {
	return &Reflector{wiring: wiring}
}

// Rotate moves the rotor one position forward
func (r *Rotor) Rotate() {
	r.position = (r.position + 1) % 26
}

// Encode encodes a single letter using the rotor
func (r *Rotor) Encode(letter byte) byte {
	// Find the index of the letter in the alphabet
	index := int(letter-'A'+byte(r.position)) % 26
	// Get the corresponding letter from the wiring
	return r.wiring[index]
}

// Decode decodes a single letter using the rotor
func (r *Rotor) Decode(letter byte) byte {
	// Find the index of the letter in the wiring
	index := strings.IndexByte(r.wiring, letter)
	// Calculate the original letter index
	return 'A' + byte((index-int(r.position)+26)%26)
}

// EncodeReflector encodes a letter using the reflector
func (ref *Reflector) Encode(letter byte) byte {
	index := int(letter - 'A')
	return ref.wiring[index]
}

// Enigma represents the Enigma machine
type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
}

// NewEnigma creates a new Enigma machine with given rotors and reflector
func NewEnigma(rotors []*Rotor, reflector *Reflector) *Enigma {
	return &Enigma{rotors: rotors, reflector: reflector}
}

// Encrypt encrypts a message using the Enigma machine
func (e *Enigma) Encrypt(message string) string {
	var result strings.Builder

	for _, letter := range message {
		if letter < 'A' || letter > 'Z' {
			result.WriteByte(byte(letter)) // Non-alphabetic characters are unchanged
			continue
		}

		// Rotate the rotors
		for _, rotor := range e.rotors {
			rotor.Rotate()
		}

		// Encode through the rotors
		for i := len(e.rotors) - 1; i >= 0; i-- {
			letter = rune(e.rotors[i].Encode(byte(letter)))
		}

		// Reflect
		letter = rune(e.reflector.Encode(byte(letter)))

		// Encode back through the rotors
		for _, rotor := range e.rotors {
			letter = rune(rotor.Decode(byte(letter)))
		}

		result.WriteByte(byte(letter))
	}

	return result.String()
}

func main() {
	// Define rotors and reflector
	rotor1 := NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", 0)     // Rotor I
	rotor2 := NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", 0)     // Rotor II
	rotor3 := NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", 0)     // Rotor III
	reflector := NewReflector("YRUHQSLDPXNGOKIETZJWVFMCBA") // Reflector B

	// Create the Enigma machine
	enigma := NewEnigma([]*Rotor{rotor1, rotor2, rotor3}, reflector)

	// Input message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the message to encrypt (uppercase letters only): ")
	message, _ := reader.ReadString('\n')
	message = strings.ToUpper(strings.TrimSpace(message)) // Convert to uppercase and trim whitespace

	// Encrypt the message
	encrypted := enigma.Encrypt(message)
	fmt.Printf("Encrypted Message: %s\n", encrypted)
}
