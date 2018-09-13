package main

const AGENTINFO = `
	select 
		SWITCH,
		LOGINID,
		NAME,
		updated
	from ossi.agents
	where
		switch = ?
	and loginid = ?
`
