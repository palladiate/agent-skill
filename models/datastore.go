package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Datastore struct {
	db *sqlx.DB
	uxidQuery *sqlx.Stmt
	agentQuery *sqlx.Stmt
	skillQuery *sqlx.Stmt
	agentSkillQuery *sqlx.Stmt
}

func (ds *Datastore) Init(conn string) (err error) {
	// TODO: Set as env
	ds.db = sqlx.MustConnect("mysql","nspangler:d00dsp34k@tcp(cmspapp03.twcable.com:3306)/ossi")

	ds.skillQuery, err = ds.db.Preparex(`
	  select
		*
	  from ossi.IVR_CII_RDM_SKILL_SPLIT a
	  where a.switchnm = ? 
	  and a.ACD = ? 
	  and a.QUEUEID = ?

	`)
	if err != nil { panic(err) }

	ds.agentQuery, err = ds.db.Preparex(`
		select 
			switch,
			loginid,
			name,
			updated
		from ossi.agents
		where
			switch in ('CSGBILLING','CSGIPR', 'CSGVIDEO', 'ICOMSBILLING',
			           'ICOMSIPR','ICOMSVIDEO','OUTSOURCER','TANDEM','COMMERCIAL')
		and loginid = ?
	`)
	if err != nil { panic(err) }

	ds.agentSkillQuery, err = ds.db.Preparex(`
		select
			h.hostname,
			h.acd,
			s.skill,
			s.level,
			s.ordinal
		from
			ossi.agentskills s
		join 
			ossi.skillsources h
		on s.switch = h.name	
		where 
			s.switch = ?
			and s.agentid = ?
	`)
	if err != nil {panic(err)}

	ds.uxidQuery, err = ds.db.Preparex(`
		select 
			SAMACCOUNTNAME,
			LTWC_EID,
			FIRST_NAME,
			LAST_NAME,
			MGMT_AREA,
			GL_COMPANY,
			REGION,
			DIVISION,
			LOCATION,
			AVAYA_LOGIN
		FROM
			ossi.UXID_INFO
		where SAMACCOUNTNAME = ?
	`)
	if err != nil {panic(err)}

	return
}

func (ds *Datastore) Close() {

	ds.skillQuery.Close()
	ds.agentQuery.Close()
	ds.agentSkillQuery.Close()

	ds.db.Close()
}
