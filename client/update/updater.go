package update

import (
	"fmt"
	"live-config/logger"
)

type Updater interface {

	//OnPropertyUpdate function that will be called once a property gets created/updated
	OnPropertyUpdate(p interface{})
}

type DummyUpdater struct {}

func (d *DummyUpdater) OnPropertyUpdate(p interface{}) {
	logger.Instance.Info(fmt.Sprintf("RECEIVED %v", p))
}
