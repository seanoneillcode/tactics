package fight

type Skill struct {
	TargetPattern Pattern
	EffectPattern Pattern
	Effects       Effect
}

func SlashSkill() *Skill {
	return &Skill{
		TargetPattern: &BasicPattern{},
		EffectPattern: &SinglePattern{},
		Effects:       &DamageActorEffect{Amount: 2},
	}
}

func FireBallSkill() *Skill {
	return &Skill{
		TargetPattern: &CrossDistancePattern{
			Distance: 4,
		},
		EffectPattern: &BallPattern{},
		Effects:       &DamageActorEffect{Amount: 1},
	}
}

func HealSkill() *Skill {
	return &Skill{
		TargetPattern: &SinglePattern{},
		EffectPattern: &SinglePattern{},
		Effects:       &HealActorEffect{Amount: 3},
	}
}
