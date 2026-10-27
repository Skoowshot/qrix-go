package asset

// FormatID is a runtime identifier of Format, used for hot-path access.
type FormatID uint64

// FormatName is a persistent identifier of Format, used for saving and reading.
type FormatName string

// Format is a hot-path definition of Asset processing.
type Format struct {
	Codec CodecID
}

// FormatInfo is a cold-path information about the Format.
type FormatInfo struct {
	Name  FormatName
	Codec CodecName
}
