package document

// Interface Segregation Principle (ISP)

// Could be bad practice to have a large, generic interface.

// LargeGenericDocument represents a document with many methods.
type LargeGenericDocument interface {
	Open()
	Close()
	Read()
	Write()
	Print()
	Scan()
}

// Client is forced to depend on a large, generic interface.
func ClientUsingLargeInterface(doc LargeGenericDocument) {
	doc.Open()
	doc.Read()
	doc.Write()
	doc.Close()
	// Client can also call Print() and Scan() methods, which it may not need.
	doc.Print()
	doc.Scan()
}

// Good practice to have smaller, focused interfaces.

// Document represents a generic document.
type Document interface {
	Open()
	Close()
}

// ReadableDocument represents a document that can be read.
type ReadableDocument interface {
	Read()
}

// WritableDocument represents a document that can be written to.
type WritableDocument interface {
	Write(data string)
}

// Client focuses on the interfaces it needs.
func ClientViewEvidence(reader ReadableDocument) {
	reader.Read()
	// Cannot accidentally access methods like Write().
}

// Client focuses on the interfaces it needs.
func ClientWriteSensorData(writer WritableDocument, data string) {
	writer.Write("Sensor data: " + data)
	// Cannot accidentally access methods like Print().
}
