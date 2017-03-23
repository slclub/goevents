package goevents

type config struct {
	chNumber int  //parallel channel numbers
	safe     int  //events running mod
	setted   bool //every events can just set once.
}

func newConf(chm int, safe int, setted bool) *config {

	if chm <= 0 {
		chm = 5
	}
	return &config{chm, safe, setted}
}

//======================events object =========================
//setting the events object
func (ev *events) Conf(chm int, safe int) {
	if ev.config.setted {
		return
	}
	ev.config = newConf(chm, safe, true)
	ev.flush()
}

//flush events by config
func (ev *events) flush() {
	ev.concurrent.chNumber = ev.config.chNumber
}
