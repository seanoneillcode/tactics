package dialog

// the map is organized by npc Name
var dialogData = map[string]*NpcDialog{
	"dave": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*Dialog{
			"": {
				lines: []*Line{
					{
						Name: "Dave",
						Text: "\"I'm Dave... \nWhat do you want?\"",
					},
					{
						Name: "Player",
						Text: "\"I'm looking for some fish\"",
					},
					{
						Name: "Player",
						Text: "\"Have you seen any?\"",
					},
					{
						Name: "Dave",
						Text: "\"No...\"",
					},
					{
						Name: "Dave",
						Text: "\"...\"",
					},
					{
						Name: "Player",
						Text: "\"Great, thanks...\"",
					},
				},
			},
		},
	},
	"peter": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*Dialog{
			"": {
				lines: []*Line{
					{
						Name: "Peter",
						Text: "\"Can you fetch a fish for me?\"",
					},
				},
			},
			"got-fish": {
				lines: []*Line{
					{
						Name: "Peter",
						Text: "\"Hey it's my fish!\"",
					},
					{
						Name: "Peter",
						Text: "\"Thanks!\"",
					},
				},
			},
		},
	},
}
