package common

type Mode string

const (
	NoneMode     Mode = "none"
	ExploreMode  Mode = "explore"
	FightMode    Mode = "fight"
	GameOverMode Mode = "game-over"
	MainMenuMode Mode = "main-menu"
)
