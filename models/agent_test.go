package models

import (
	"testing"
)

const CONN = "nspangler:d00dsp34k@tcp(cmspapp03.twcable.com:3306)/ossi"

func TestDatastore_GetAgent(t *testing.T) {
	var ds Datastore
	ds.Init(CONN)

	agent, err := ds.GetAgent("P2158181")
	if err != nil {t.Fail()}
	if agent.Eid != "E160941" { t.Fail() }
	ds.Close()

	agent, err = ds.GetAgent("P2158181")
	if err == nil {t.Fail()}
}

