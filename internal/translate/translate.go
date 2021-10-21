package translate

import (
	"context"
	"fmt"
	"log"

	translate "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func Do(ctx context.Context, credentials []byte, text string, source string, dest string) (string, error) {
	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	target, err := language.Parse(dest)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
		return "", err
	}

	translation, err := client.TranslateText(ctx, &translatepb.TranslateTextRequest{
		Contents:           []string{text},
		SourceLanguageCode: source,
		TargetLanguageCode: target.String(),
		MimeType:           "text/plain",
		Parent:             fmt.Sprintf("projects/%s/locations/global", "translate-329709"),
	})
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
		return "", err
	}

	return translation.String(), nil
}
