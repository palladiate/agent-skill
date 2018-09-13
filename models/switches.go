package models

type Switches []string

func (ds *Datastore) Switches() (switches Switches) {
	rows, err := ds.db.Queryx(`
		select distinct name from ossi.skillsources
	`)
	if err != nil { return }

	for rows.Next() {
		var region string
		rows.Scan(&region)
		switches = append(switches,region)
	}
	return
}
