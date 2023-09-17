package document

import "fmt"

// TextDocument is a specific type of document that can be read and written.
type TextDocument struct{}

// Single Responsibility Principle (SRP)

func (t TextDocument) Open() {
	fmt.Println("Opening text document...")
}

func (t TextDocument) Close() {
	fmt.Println("Closing text document...")
}

func (t TextDocument) Read() {
	fmt.Println("Reading text document...")
}

func (t *TextDocument) Write(data string) {
	fmt.Println("Writing to text document:", data)
}

func (t TextDocument) Print() {
	fmt.Println("Printing text document...")
}

func (t TextDocument) Scan() {
	fmt.Println("Scanning text document...")
}

// Bad practice for Single Responsibility Principle (SRP)

func (t TextDocument) ReadWriteAndPrint(data string) {
	fmt.Println("Reading text document...")
	fmt.Println("Writing to text document:", data)
	fmt.Println("Printing text document...")
}
