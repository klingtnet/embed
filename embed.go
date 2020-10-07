package embed

// Embed defines methods for retrieving embedded content.
type Embed interface {
	// Files returns an alphabetically sorted list of the embedded files.
	Files() []string
	// File returns the content of the file embedded as path.
	File(path string) []byte
	// FileString is a convenience function
	// that works like File but returns a string
	// instead of a byte slice.
	FileString(path string) string
}
