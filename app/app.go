package app

import alexa "gitlab.com/dasjott/alexa-sdk-go"

func launch(c *alexa.Context) {
	c.Ask(c.T("HELLO", "PAUSE", "WELCOME_CONFIM"))
}

func getJoke(c *alexa.Context) {
	joke := c.T("JOKE")
	c.Ask(joke+c.T("PAUSE", "LAUGH_VOICE", "HELP_REPROMPT")).SimpleCard(c.T("SKILL_NAME"), joke)
}

func help(c *alexa.Context) {
	c.Ask(c.T("HELP_MESSAGE"), c.T("HELP_REPROMPT"))
}

func bye(c *alexa.Context) {
	c.Tell(c.T("STOP_MESSAGE"))
}

// Handlers maps intent functions to intent names
var Handlers = alexa.IntentHandlers{
	"LaunchRequest":       launch,
	"GetNewJokeIntent":    getJoke,
	"AMAZON.HelpIntent":   help,
	"AMAZON.CancelIntent": bye,
	"AMAZON.NoIntent":     bye,
	"AMAZON.YesIntent":    getJoke,
	"Unhandled":           bye,
}
