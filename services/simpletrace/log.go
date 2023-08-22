package simpletrace

import (
	"github.com/massimo-gollo/benchy/pkg/log"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = log.NewDefaultLogger()

}
