package explore

type Control struct {
	Command string
}

func (r *Control) ExitGame() {
	r.Command = "exit"
}
