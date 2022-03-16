package core

import (
	"log"

	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Map struct {
	Level           *Level
	transitionTimer int64
	link            *Link
	levels          []*Level
}

func NewMap() *Map {
	return &Map{
		levels: []*Level{},
	}
}

func (m *Map) LoadLevel(name string) {
	for _, level := range m.levels {
		if level.name == name {
			log.Printf("using existing loaded level")
			m.Level = level
			return
		}
	}
	log.Printf("loading new level")
	newLevel := NewLevel(name)
	m.Level = newLevel
	m.levels = append(m.levels, newLevel)
}

func (m *Map) Update(delta int64, state *State) {
	if m.transitionTimer > 0 {
		m.transitionTimer = m.transitionTimer - delta
		if m.transitionTimer <= 0 {
			m.transitionToLevel(state)
		}
	}
	m.Level.Update(delta, state)
}

func (m *Map) StartTransition(link *Link) {
	m.transitionTimer = 100
	m.link = link
}

func (m *Map) transitionToLevel(state *State) {
	m.LoadLevel(m.link.toLevel)
	var toLink *Link
	for _, link := range m.Level.links {
		if link.name == m.link.name {
			toLink = link
		}
	}
	if toLink == nil {
		log.Fatalf("failed to transition to level with no link")
	}
	wx, wy := common.WorldToTile(state.Player.Character.pos)
	// offset of position between tile position while Character is moving
	offset := state.Player.Character.pos.Sub(common.PositionFromInt(wx*common.TileSize, wy*common.TileSize))
	if state.Player.Character.velocity.X < 0 {
		// magic
		offset.X = offset.X - common.TileSize
	}
	if state.Player.Character.velocity.Y < 0 {
		// magic
		offset.Y = offset.Y - common.TileSize
	}

	state.Player.SetPosition(offset.Add(common.PositionFromInt(toLink.x, toLink.y)))
	m.link = nil
}
