package core

type UI struct {
	IsActive bool
	UIType   string // inventory, shop
}

func NewUI() *UI {
	return &UI{
		IsActive: false,
		UIType:   "",
	}
}

func (i *UI) Open(uiType string) {
	i.UIType = uiType
	i.IsActive = true
}

func (i *UI) Close() {
	i.IsActive = false
	// todo set player active ?
}

func (i *UI) IsInventoryActive() bool {
	return i.IsActive && i.UIType == "inventory"
}
