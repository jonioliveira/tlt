package credentials

import (
	"encoding/base64"
	"os"
)

func GetEnvCreds(env string) ([]byte, error) {
	encodedCreds := os.Getenv(env)

	creds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
