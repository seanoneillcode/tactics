package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type State struct {
	NextMode         common.Mode
	ActiveTeam       *Team
	PlayerController *PlayerController
	PlayerTeam       *Team
	AiController     AiController
	AiTeam           *Team
	Camera           *Camera
	Scene            *Scene
}

func (s *State) Update(delta int64) {
	s.AiController.Update(delta, s)
	s.PlayerController.Update(delta, s)
	s.Camera.Update(delta, s)
	s.PlayerTeam.Update(delta, s)
	s.AiTeam.Update(delta, s)
}

func (s *State) Draw(camera *Camera) {
	s.Scene.Draw(camera)
	s.PlayerTeam.Draw(camera)
	s.AiTeam.Draw(camera)
}

func (s *State) ChangeMode(mode common.Mode) {
	s.NextMode = mode
}

func (s *State) StartFight(playerActors []*Actor, enemyActors []*Actor, sceneName string) {
	s.Scene = NewScene(sceneName, playerActors, enemyActors)
	s.PlayerTeam = NewTeam(playerActors)
	s.AiTeam = NewTeam(enemyActors)
	s.PlayerController.StartTurn(s)
}
