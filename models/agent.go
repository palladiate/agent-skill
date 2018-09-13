package models

type Agent struct {
	UXID
	Avaya []Avaya    `json:"Avaya"`
}

type UXID struct {
	Account string   `json:"PID" db:"SAMACCOUNTNAME"`
	Eid string       `json:"EID" db:"LTWC_EID"`
	FirstName string `json:"FirstName" db:"FIRST_NAME"`
	LastName string  `json:"LastName" db:"LAST_NAME"`
	Area string      `json:"Area" db:"MGMT_AREA"`
	Company string   `json:"Company" db:"GL_COMPANY"`
	Region string    `json:"Region" db:"REGION"`
	Division string  `json:"Division" db:"DIVISION"`
	Location string  `json:"Location" db:"LOCATION"`
	Login string     `json:"-" db:"AVAYA_LOGIN"`
}

func (ds Datastore) GetAgent(loginId string) (agent Agent, err error) {
	agent.Avaya = make([]Avaya,0)
	var tmpAgent UXID
	var tmpAvaya []Avaya
	rows, err := ds.uxidQuery.Queryx(loginId)
	if err != nil { return }

	for rows.Next() {
		rows.StructScan(&tmpAgent)
		tmpAvaya = append(tmpAvaya, Avaya{
			LoginId:     tmpAgent.Login,
		})
	}

	agent = Agent{
		UXID:  tmpAgent,
		Avaya: tmpAvaya,
	}

	return
}