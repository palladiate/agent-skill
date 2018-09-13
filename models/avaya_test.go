package models

import "testing"

var agent = Agent{
	UXID{
		Account:   "P2158181",
		Eid:       "E160941",
		FirstName: "Jacinto",
		LastName:  "Gomez",
		Area:      "CC-McAllen Call Center",
		Company:   "McAllen TX",
		Region:    "Care Call Centers Grp 3",
		Division:  "CC-Customer Ops",
		Location:  "3701 N. 23rd Street McAllen TX, 78501",
		Login:     "4340043",
	},
	[]Avaya{
		{
			LoginId: "4340043",
			Switch:  "CSGIPR",
			Name:    "",
			Updated: "",
			Skills:  nil,
		},
	},
}
func TestDatastore_AgentInfo(t *testing.T) {
	var ds Datastore
	ds.Init(CONN)
	var tmpAgent = agent

	err := ds.AgentInfo(&tmpAgent)
	if err != nil {t.Fail()}
	if agent.Avaya[0].Switch == "" { t.FailNow() }

	ds.Close()
	err = ds.AgentInfo(&tmpAgent)
	if err == nil {t.Fail()}
}

func TestDatastore_AgentSkills(t *testing.T) {
	var ds Datastore
	ds.Init(CONN)
	var tmpAgent = agent

	err := ds.AgentSkills(&tmpAgent.Avaya[0])
	if err != nil { t.Fail() }
	if agent.Avaya[0].Skills == nil { t.Fail() }

	ds.Close()
	err = ds.AgentSkills(&tmpAgent.Avaya[0])
	if err == nil { t.Fail() }
}
