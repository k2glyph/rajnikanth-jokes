package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/k2glyph/rajnikanth-jokes/app"
	alexa "gitlab.com/dasjott/alexa-sdk-go"
)

func main() {
	alexa.AppID = "amzn1.ask.skill.81239321-fcb3-44f5-bee7-fe2dc3118d40"
	alexa.Handlers = app.Handlers
	alexa.LocaleStrings = app.Locales
	lambda.Start(alexa.Handle)
}
