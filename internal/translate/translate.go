package translate

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func Do(ctx context.Context, credentials []byte, text string, source string, dest string) (string, error) {
	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		return "", err
	}
	defer client.Close()

	target, err := language.Parse(dest)
	if err != nil {
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
		return "", err
	}

	return translation.Translations[0].TranslatedText, nil
}
