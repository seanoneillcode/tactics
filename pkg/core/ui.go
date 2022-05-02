package core

type UI struct {
	IsActive bool
	UIType   UIName
}

type UIName string

const (
	NoUI        UIName = "none"
	MenuUI      UIName = "menu"
	ItemsUI     UIName = "items"
	EquipmentUI UIName = "equipment"
	MagicUI     UIName = "magic"
	SettingsUI  UIName = "settings"
	FileUI      UIName = "file"
)

func NewUI() *UI {
	return &UI{
		IsActive: false,
		UIType:   NoUI,
	}
}

func (i *UI) Open(uiType UIName) {
	if uiType == ItemsUI || uiType == MenuUI || uiType == EquipmentUI {
		i.UIType = uiType
		i.IsActive = true
	}
}

func (i *UI) Close() {
	i.IsActive = false
	// todo set player active ?
}

func (i *UI) IsInventoryActive() bool {
	return i.IsActive && i.UIType == "items"
}

func (i *UI) IsEquipmentActive() bool {
	return i.IsActive && i.UIType == "equipment"
}

func (i *UI) IsMenuActive() bool {
	return i.IsActive && i.UIType == "menu"
}
