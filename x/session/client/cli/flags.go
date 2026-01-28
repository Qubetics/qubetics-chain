package cli

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/pflag"
)

const (
	flagSignature = "signature"
)

func GetSignature(flags *pflag.FlagSet) ([]byte, error) {
	s, err := flags.GetString(flagSignature)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, fmt.Errorf("signature cannot be empty")
	}

	return base64.StdEncoding.DecodeString(s)
}
