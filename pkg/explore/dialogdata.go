package explore

type DialogData struct {
	Lines []*Line
}

type Line struct {
	Name string
	Text string
}

// the map is organized by npc Name
var dialogData = map[string]*NpcDialog{
	"dave": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*DialogData{
			"": {
				Lines: []*Line{
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
	"anne": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*DialogData{
			"": {
				Lines: []*Line{
					{
						Name: "Anne",
						Text: "\"Can you fetch a fish for me?\"",
					},
				},
			},
			"got-fish": {
				Lines: []*Line{
					{
						Name: "Anne",
						Text: "\"Hey it's my fish!\"",
					},
					{
						Name: "Anne",
						Text: "\"Thanks!\"",
					},
				},
			},
		},
	},
	"wallace": {
		dialogs: map[string]*DialogData{
			"": {
				Lines: []*Line{
					{
						Name: "Wallace",
						Text: "\"How did you get back here?\"",
					},
				},
			},
		},
	},
}
