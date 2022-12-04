package arsenal

type Player struct {
	Weapon  string
	Value   int
	Fight_s bool
	Fight_p bool
	Fight_r bool
	Weak    string
	Strong  string
}

func (p *Player) SetName(name string) {
	p.Weapon = name
}

func (p *Player) SetPlayerStats(Weapon string, Value int, Fight_s bool, Fight_p bool, Fight_r bool) {
	p.Weapon = Weapon
	p.Value = Value
	p.Fight_s = Fight_s
	p.Fight_p = Fight_p
	p.Fight_r = Fight_r
}

func (p *Player) AgainstPaper() bool {
	return p.Fight_p
}

func (p *Player) AgainstRock() bool {
	return p.Fight_r
}

func (p *Player) AgainstScissors() bool {
	return p.Fight_s
}

func (p *Player) GetValue() int {
	return p.Value
}

func (p *Player) GetType() string {
	return p.Weapon
}

func (p *Player) GetDefeat() string {
	return p.Weak
}

func (p *Player) GetWin() string {
	return p.Strong
}
