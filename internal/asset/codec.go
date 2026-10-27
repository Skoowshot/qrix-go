package asset

// CodecID is a runtime identifier of Codec, used for hot-path access.
type CodecID uint64

// CodecName is a persistent identifier of Codec, used for saving and loading.
type CodecName string

// CodecInfo is a cold-path information about the Codec.
type CodecInfo struct {
	Name CodecName
}

// Codec handles decoding and encoding of asset data.
type Codec interface {
	// Decode reads and decodes data from the provided source.
	Decode(source Source) (Data, error)

	// Encode encodes and writes data to the provided destination.
	Encode(data Data, dest Dest) error
}
