package app

import alexa "gitlab.com/dasjott/alexa-sdk-go"

// Locales are all the speeches
var Locales = alexa.Localisation{
	"en-US": alexa.Translation{
		"PAUSE":          "<break time='600ms'/>",
		"HELP_MESSAGE":   "You can say tell me a rajnikanth jokes, or, you can say exit... What can I help you with?",
		"LAUGH_VOICE":    "<audio src='soundbank://soundlibrary/human/amzn_sfx_laughter_giggle_02'/>",
		"CHEERS_VOICE":   "<audio src='soundbank://soundlibrary/human/amzn_sfx_crowd_cheer_med_01'/>",
		"HELP_REPROMPT":  "Do you wanna hear one more?",
		"SKILL_NAME":     "rajnikanth jokes",
		"HELLO":          "Welcome to rajnikanth jokes! ",
		"WELCOME_CONFIM": "Do you want me to tell you joke?",
		"STOP_MESSAGE": []string{
			"Bye.",
			"See you later.",
			"See you next time.",
		},
		"JOKE": []string{
			"Rajinikanth knows why this kolaveri kolaveri.\r\n",
			"Rajinikanth doesn't answer nature's call. Nature answers Rajinikanth's call.\r\n",
			"Rajnikanth`s pulse is measured in Richter scale!",
		},
	},
}
