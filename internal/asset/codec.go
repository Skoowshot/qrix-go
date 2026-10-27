package asset

import (
	"io"
)

// CodecID is a runtime identifier of Codec, used for hot-path access.
type CodecID uint64

// CodecName is a persistent identifier of Codec, used for saving and loading.
type CodecName string

// Codec handles decoding and encoding of asset data.
type Codec interface {
	// Decode reads and decodes data from the provided source.
	Decode(source Source) (Data, error)

	// Encode encodes and writes data to the provided destination.
	Encode(data Data, dest Dest) error
}

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
