package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/consensys-hellhound/log"
)

var Logger = logrusLogger{log.New("pythia")}

type logrusLogger struct {
	*logrus.Entry
}

func (l logrusLogger) Log(keyvals ...interface{}) error {
	if len(keyvals)%2 == 0 {
		fields := logrus.Fields{}
		for i := 0; i < len(keyvals); i += 2 {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		}
		l.WithFields(fields).Info()
	} else {
		l.Info(keyvals)
	}
	return nil
}

const Method = "method"
const Took = "took"
const Err = "err"
