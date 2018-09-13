package models

type Avaya struct {
	LoginId string           `json:"LoginId" db:"loginid"`
	Switch string            `json:"Switch" db:"switch"`
	Name string              `json:"Name" db:"name"`
	Updated string           `json:"Updated" db:"updated"`
	Skills     []Skill       `json:"Skills"`
}

type Skill struct {
	AvayaSkill
	SkillInformation
}

type AvayaSkill struct {
	Hostname string
	Acd int
	SkillId string `json:"SkillId" db:"skill"`
	Level string `json:"Level" db:"level"`
	Ordinal int `json:"Ordinal" db:"ordinal"`
}


func (ds *Datastore) AgentInfo(agent *Agent) error {
	var tmpAvaya Avaya

	for i, login := range agent.Avaya {
		agentRows, aErr := ds.agentQuery.Queryx(login.LoginId)
		if aErr != nil {
			return aErr
		}

		for agentRows.Next() {
			agentRows.StructScan(&tmpAvaya)
			agent.Avaya[i] = tmpAvaya
			sErr := ds.AgentSkills(&agent.Avaya[i])
			if sErr != nil { return sErr }
		}
	}

	return nil
}

func (ds *Datastore) AgentSkills(avaya *Avaya) error {
	var tmpSkill AvayaSkill
	skillRows, sErr := ds.agentSkillQuery.Queryx(avaya.Switch, avaya.LoginId)
	if sErr != nil {
		return sErr
	}

	for skillRows.Next() {
		skillRows.StructScan(&tmpSkill)
		avaya.Skills = append(avaya.Skills, Skill{
			AvayaSkill:       tmpSkill,
			SkillInformation: ds.SkillInformation(tmpSkill),
		})
	}
	return nil
}

func (ds *Datastore) SkillInformation(skill AvayaSkill) (info SkillInformation) {

	rows, err := ds.skillQuery.Queryx(skill.Hostname, skill.Acd, skill.SkillId)
	if err != nil { return }

	rows.Next()
	err = rows.StructScan(&info)

	return
}
