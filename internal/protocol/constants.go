package protocol

// HTTP Methods
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// HTTP Status Codes
const (
	StatusOK       = "200 OK"
	StatusCreated  = "201 Created"
	StatusNotFound = "404 Not Found"
)

// HTTP Header Keys
const (
	HeaderContentLength   = "Content-Length"
	HeaderContentType     = "Content-Type"
	HeaderUserAgent       = "User-Agent"
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"
)

// Common Content Types
const (
	ContentTypeTextPlain   = "text/plain"
	ContentTypeOctetStream = "application/octet-stream"
)

// Suppored Content Encodings
const (
	ContentEncodingGzip = "gzip"
)

// HTTP Version
const Version = "HTTP/1.1"
