package security

import (
	"encoding/base64"
)

func EncodeFile(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
