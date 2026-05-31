package constants

import "net/textproto"

// Keempat variable ini digunakan untuk keamanan (security), jika ada FE yang ingin mengakses BE (service), maka harus ada keempat header ini
var (
	XServiceName 	= textproto.CanonicalMIMEHeaderKey("x-service-name")
	XApiKey 		= textproto.CanonicalMIMEHeaderKey("x-api-key")
	XRequestAt 		= textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization 	= textproto.CanonicalMIMEHeaderKey("authorization")
)