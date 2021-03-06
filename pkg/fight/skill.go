package fight

type Skill struct {
	TargetPattern Pattern
	EffectPattern Pattern
	Effects       []Effect
}

func SlashSkill() *Skill {
	return &Skill{
		TargetPattern: &CrossPattern{},
		EffectPattern: &SinglePattern{},
		Effects: []Effect{
			&DamageActorEffect{Amount: 2},
		},
	}
}

func ArrowSkill() *Skill {
	return &Skill{
		TargetPattern: &LinePattern{},
		EffectPattern: &SinglePattern{},
		Effects: []Effect{
			&DamageActorEffect{Amount: 1},
		},
	}
}

func FireBallSkill() *Skill {
	return &Skill{
		TargetPattern: &CrossDistancePattern{
			Distance: 4,
		},
		EffectPattern: &BallPattern{},
		Effects: []Effect{
			&DamageActorEffect{Amount: 1},
		},
	}
}

func HealSkill() *Skill {
	return &Skill{
		TargetPattern: &SinglePattern{},
		EffectPattern: &SinglePattern{},
		Effects: []Effect{
			&HealActorEffect{Amount: 3},
		},
	}
}
