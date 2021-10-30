package version

import (
	"io/ioutil"
	"os"
)

const version = "TLT_VERSION"

func GetVersion() (string, error) {
	version := os.Getenv(version)
	if version == "" {
		versionByte, err := ioutil.ReadFile("VERSION")
		if err != nil {
			return "", err
		}
		version = string(versionByte)
	}
	return version, nil
}
