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
						name: "player",
						text: "\"Great thanks. Do you know where I can get some gameplay around here?\"",
					},
					{
						name: "player",
						text: "\"Who are you?\"",
					},
					{
						name: "",
						text: "\"My name is dave.\"",
					},
					{
						name: "dave",
						text: "\"Welcome to the game!\"",
					},
					{
						name: "player",
						text: "\"Great thanks. Do you know where I can get some gameplay around here?\"",
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
						text: "Can you fetch my fish for me?",
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
