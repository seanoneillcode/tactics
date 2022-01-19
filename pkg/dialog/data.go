package dialog

// the map is organized by npc name
var dialogData = map[string]*NpcDialog{
	"dave": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*Dialog{
			"": {
				lines: []*Line{
					{
						name: "Dave",
						text: "\"I'm Dave... \nWhat do you want?\"",
					},
					{
						name: "Player",
						text: "\"I'm looking for some fish\"",
					},
					{
						name: "Player",
						text: "\"Have you seen any?\"",
					},
					{
						name: "Dave",
						text: "\"No.\"",
					},
					{
						name: "Dave",
						text: "\"...\"",
					},
					{
						name: "Player",
						text: "\"Great, thanks...\"",
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
						name: "peter",
						text: "Can you fetch a fish for me?",
					},
				},
			},
			"got-fish": {
				lines: []*Line{
					{
						name: "peter",
						text: "Hey it's my fish!",
					},
					{
						name: "peter",
						text: "Thanks!",
					},
				},
			},
		},
	},
}
