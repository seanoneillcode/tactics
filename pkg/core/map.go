package core

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"log"
)

type Map struct {
	Level           *Level
	transitionTimer int64
	link            *Link
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) LoadLevel(name string) {
	newLevel := NewLevel(name + ".json")
	m.Level = newLevel
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
	wx, wy := common.WorldToTile(state.Player.character.pos)
	// offset of position between tile position while character is moving
	offset := state.Player.character.pos.Sub(common.VectorFromInt(wx*common.TileSize, wy*common.TileSize))
	if state.Player.character.velocity.X < 0 {
		// magic
		offset.X = offset.X - common.TileSize
	}
	if state.Player.character.velocity.Y < 0 {
		// magic
		offset.Y = offset.Y - common.TileSize
	}

	state.Player.SetPosition(offset.Add(common.VectorFromInt(toLink.x, toLink.y)))
	m.link = nil
}
