package asset

import "io"

// SourceType represents the format or state of an asset source.
type SourceType uint8

// DestType represents the format or state of an asset destination.
type DestType uint8

const (
	// SourceTypeRaw indicates that the source contains unprocessed, raw data.
	SourceTypeRaw SourceType = iota
	// SourceTypeBaked indicates that the source contains preprocessed, baked data.
	SourceTypeBaked
)

const (
	// DestTypeRaw indicates that the destination expects raw data.
	DestTypeRaw DestType = iota
	// DestTypeBaked indicates that the destination expects baked data.
	DestTypeBaked
)

// Source represents a readable asset stream with a specific SourceType.
type Source interface {
	// Type returns the SourceType of the data stream.
	Type() SourceType
	io.ReadCloser
}

// Dest represents a writable asset stream with a specific DestType.
type Dest interface {
	// Type returns the DestType of the data stream.
	Type() DestType
	io.WriteCloser
}
