package asset

// Data represents generic asset content as a key-value map.
type Data map[string]any

// Asset is a high-level wrapper around Data with defined FormatID.
type Asset struct {
	Format FormatID
	Data   Data
}
