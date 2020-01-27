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
			"Sachin: The God of cricket...... \r\n Sehwag: The Rajinikanth of cricket.\r\n",
			"The only time Rajini needed a body double was crying scene.\r\n",
			"Rajinikanth knows why this kolaveri kolaveri.\r\n",
			"Rajinikanth threw a grenade and killed 50 people, then grenade exploded.\r\n",
			"Rajinikanth doesn't answer nature's call. Nature answers Rajinikanth's\r\n",
			"Einstein Said: Everything is relative. Karunanidhi said: Relative is everthing. Rajinikanth said: I am everything.",
			"Rajnikanth was shot today... Tomorrow is the bullet`s funeral!",
			"Rajnikanth killed a terrorist in Pakistan 'via Bluetooth'!",
			"Rajnikanth can draw a straight line with a compass!0",
			"Rajnikanth knows who let the dogs out!!",
			"Rajnikanth`s pulse is measured in Richter scale!",
			"The new Rupee symbol is actually Rajnikanth`s signature!!!",
			"Rajnikanth has a statue of Madame Tussauds at his house!!",
			"Rajnikanth is the secret of Boost`s energy; and Complan is a Rajnikanth boy!",
			"Rajnikanth participated in 100m race, obviously he came first, but Einstein died watching that, since Light came second!!",
			"Intel's new ad: Rajnikanth Inside",
			"When Alexander Graham Bell first used his telephone, he realized that he already had two missed calls from Rajinikanth.",
			"The apple which fell on Newton was actually thrown by Rajinikanth!",
			"An email was sent from Mysore to Bangalore, Rajinikanth stopped it in mandya!",
			"Rajinikanth can whistle in 5 different languages!",
			"Only Rajinikanth knows why Mona Lisa is smiling.",
			"Rajinikanth is the person in the world who can make his girlfriend admit her mistake!",
			"Rajinikanth went to the world cooking championship...of course Rajini won. But guess what did he make in final??? Lal mirchi ki meethi kheer",
			"Roger Federer: I know everything about tennis. Ask me anything and I will answer. Rajnikanth: Kanna.. How many holes are there in the net ?",
			"When Rajnikant switches on his AC, then winter starts in India.",
			"Facebook founder Mark Zukerburg is hospitalized with serious injury…… Reason- Rajnikant poked him on facebook…",
			"Rajnikant once wrote his autobiography ………. Today that book is known as Guiness book of world records….",
			"WhY ‘EARTHQUAKE’ OCCURS ON EARTH??? BCOZ at that Time RAJNIKANT‘S MOBILE IS ON VIBRATE MODE!!!",
			"Rajnikant can make calls from his iPod to his iPad…!!!",
			"Once Rajnikant Participated In Bike Race Dont Even Try To Guess What",
			"Happened Rajnikant Won The Race On Neutral Gear..",
			"Rajnikant creats his new mail i.d. gmail@rajnikant.com",
			"Rajnikant was practising for spelling test. The rough sheet he used is today known as the oxford dictionary!!",
			"Rajnikant opened his refrigerator and forgot to shut the door, Winters  arrive in India!",
			"Rajnikant was once told to choose 3 subjects when he got admission in jr.college. He chose science,arts and commerce!",
			"CAT is outdated. Now the students have 2 prepare for RAT= RAJNIKANT  APTITUDE TEST.",
			"Hrithik tried to participate in a dance competition with Rajnikanth. Result: He is in a wheel chair in Gujarish.",
			"A basket ball player spun a basketball on his finger and asked Rajnikant can u do it? Rajani-how do u think the earth spins?",
			"There is no theory of evolution. Just a list of creatures Rajnikant has allowed to live.",
			"Outer space exists because it’s afraid to be on the same planet with Rajnikant.",
			"When Rajnikant does a pushup, he isn’t lifting himself up, he’s pushing the Earth down.",
			"Rajnikant doesn’t wear a watch, HE decides what time it is.",
			"Rajnikant gave Mona Lisa that smile.",
			"Rajnikant does not get frostbite. Rajnikant bites frost.",
			"Once upon a time Dinosaurs borrowed money from RAJNIKANT and dint pay him back.That was the last day when anyone saw Dinosaurs.",
			"Once Rajnikant raced with time.Result?Time is still Running.",
			"When Rajnikant was studyin in 3rd std. some1 stole his rough note & Now they call it as Wikipedia!",
			"ONCE RAJNI TAUGHT A KID HOW TO PLAY Counter Strike NOW HE IS NOW KNOWN AS OSAMA BIN LADEN.",
			"Rajnikanth  joined CID. Guess what happened?? Rajnikanth found the criminal before the crime.",
			"EK BAAR RAJNIKANT NE EK BACCHE KO KIS KAR LEYA AAJ VO BACCHA IMRAN HASMI KEHLATA HAI.",
			"Rajnikant mom- beta rajni ye solar heater pani garam nhi kr raha kuch kar. rajnikanta :- yanna rascala tum ruko maa mein suraj ko thik karke",
			"Rajnikant got selected in Roadies. next day during voteout.Rajnikant: I m sorry Raghu. apka safar yhi khatam hota hai. aata hu.",
			"A farmer instead of scarecrow kept a photo of rajnikant in his field.u know wht happened-birds takes the grain back taken last year!",
			"Once Rajnikant was traveling on a train suddenly he saw a dark black boy. Rkant said ‘ae kon” after that the boy is known as popular singer AKON.",
			"once rajnikant sufferer from loose motion and now his loose motion known as RAJNIGANDHA.",
			"Rajinikanth can delete the Recycle Bin.",
			"Rajinikanth once ordered a plate of idli in McDonald’s, and got it.",
			"The Bermuda Triangle used to be the Bermuda Square, until Rajinikanth kicked one of the corners off.",
			"Rajnikanth can make any lady pregnant via Bluetooth…",
			"Scientific community’s question “Which liquid turns solid on heating?” , answered by Rajnikanth – DOSA",
			"Obama request to INDIA: plz u take all power of America against of  RAJNIKANT……..(America want to adopt RAJNIKANT)",
			"Indian Space Research Organization ISRO…. closed up……Why Coz Rajinikant  took all the Rockets for Diwali…",
			"WORLD famous SUPER HEROES like he man, super man, spider man,bat man etc. get’s coaching on RAJNI home.",
			"Rajni can run bycycle without tyres at the speed of 10000 km/milli second.",
			"When Rajinikanth enters a room, he doesn’t turn the lights on.. he turns the dark off.",
			"Rajinikant can give pain to Painkillers and headache to Anacin.",
			"Rajnikanth once told a lady – “Tu stage pe jake gaa”. That lady is known as lady gaga.",
			"When Rajnikanth stares at the sun in anger, the sun hides behind the moon, and this phenomena is knows as a Solar Eclipse!",
			"Rajnikant does not install an anti-virus on his PC. All computer virus are looking for an Anti-Rajanikant software to save themselves from hands of Rajanikant.",
			"When Graham Bell invented telephone, he already had 2 missed calls from Rajnikanth.",
			"Girl (romantically) to Rajnikanth: 1 chutki sindoor ki keemat tum kya jano? Rajnikanth: 0.00078924576 Rs. per gram. Don`t mess with Rajni!",
			"Once Rajnikanth threw an ignited cigarette up in the sky. It fell on a planet, which is now known as `SUN`.",
			"Q: Why Rajnikanth doesn`t play cricket?  A: Bcoz Sachin Tendulkar requested him to keep his world records intact.",
			"Part of apple`s logo that is missing was eaten by Rajnikanth.",
			"Headlines of Today: Ek train cycle ki chapet mein aayi, train mein sawar sabhi log mare gaye.",
			"Ram and Ravan seriously yudh kar rahe the ki Achanak!  Ravan: Chal bye. Ram: Kyu!!??.. Ravan: Nahi yaar, Bye. Ram: Abey par hua kya?!.Ravan : Piche dekh!. Rajnikant!!!",
			"RAJNIKANTH purchased a 40×40 acre land with 4 wells at each corner of that square site. Guess y? Simple. TO PLAY CARROM….",
			"Once a mosquito bitten rajni so guess waht happened???….. Mosquito suffered from dengue.",
			"When ranjnikanth is Driving Car on the road , all signal’s on road get automatic Green for him , and Red for all other’s.",
			"GABAR Say’s: Kitne ADMI Thai . Kaliya Say’s: “Rajnikant” Sarkar. GABAR left Ramgad with his army..",
			"bill getes: mera ghar itna bada hai ki ander local train chalti hai. rajnikanth: mera ghar itna bada hai ki ghar me call karne k leya kone me chale jao to roaming lagta hai.",
			"Rajnikant lives vicariously through himself.",
			"Rajnikant's enemies list him as their emergency contact number.",
			"Rajnikant can speak Russian… in French.",
			"Rajnikant's business cards just say: \"I'll call you.\"",
			"Rajnikant's mother has a tattoo that says \"son\"",
			"If Rajnikant were to mail a letter without postage, it would still get there.",
			"Rajnikant has won the lifetime achievement award, twice.",
			"Mosquitos refuse to bite Rajnikant, purely out of respect.",
			"Google uses Rajnikant as a search engine.",
			"Rajnikant once won the world series of poker using UNO cards.",
			"Rajnikant once turned a vampire into a vegetarian.",
			"Rajnikant's signature won a Pulitzer.",
			"Rajnikant's moustache alone has experienced more than a lesser man’s entire body.",
			"When it is raining, it is because Rajnikant is thinking of something sad.",
			"Rajnikant's blood smells like cologne.",
			"Cuba imports cigars from Rajnikant.",
			"In museums, only Rajnikant is allowed to touch the art.",
			"If Rajnikant were to punch you in the face, you would have to fight off the urge to thank him.",
			"In the bowling alley, Rajnikant bowls overhand.",
			"Rajnikant is fluent in all languages, including three that he only speaks.",
			"Once while sailing around the world, Rajnikant discovered a shortcut.",
			"Rajnikant's passport requires no photograph.",
			"Rajnikanth threw the apple on Newton's head from above! Newton thought it was gravity",
			"Rajni can hear & understand each word of Manmohan Singh clearly!",
			"Rajni could make PAPPU dance! Now, Pappu CAN dance saala",
			"Rajni doesn't like the sound of letter \"D\". So, DJANGO decided to keep D silent!",
			"Once Rajinikanth created magnet now scientist call them black hole.",
			"One day Rajinikanth killed Prime Minister of USA . Now USA don’t have Prime Minister.",
			"Once Rajinikanth accidently dropped wallet in bank now it is called as Swiss Bank.",
			"USA is ready to give its all power to India but they are demanding Rajinikanth.",
			"No one is perfect except Rajinikanth.",
			"Once Rajinikanth was a guest contestant in KBC. Amitabh says Computerji,Rajinikanthji ko phele sawal pooochiye. Computer : Main life line use karna chahta hu.",
			"WhY ‘EARTHQUAKE’ OCCURS ON EARTH??? BCOZ at that Time RAJNIKANT‘S MOBILE IS ON VIBRATE MODE!!!",
		},
	},
}
