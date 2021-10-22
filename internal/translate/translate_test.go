package translate_test

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"github.com/jonioliveira/tlt/internal/translate"
	"github.com/stretchr/testify/require"
)

func getCreds() ([]byte, error) {
	encodedCreds := os.Getenv("TRANSLATION_API_CREDS")

	creds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return nil, err
	}

	return creds, nil
}

func TestDoTranslation(t *testing.T) {
	t.Parallel()
	for scenario, fn := range map[string]func(t *testing.T, r *require.Assertions, ctx context.Context){
		"Invalid Key": testInvalidKey,
		"Valid":       testValid,
	} {
		t.Run(scenario, func(st *testing.T) {
			r := require.New(st)
			ctx := context.Background()
			fn(st, r, ctx)
		})
	}
}

func testInvalidKey(t *testing.T, r *require.Assertions, ctx context.Context) {
	_, err := translate.Do(ctx, []byte{}, "Hello World", "PT", "EN")
	r.Error(err)
}

func testValid(t *testing.T, r *require.Assertions, ctx context.Context) {
	creds, err := getCreds()
	r.NoError(err)

	result, err := translate.Do(ctx, creds, "Hello World", "PT", "EN")

	r.Equal(result, "Olá Mundo")
	r.NoError(err)
}
