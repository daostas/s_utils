package s_utils

import (
	"github.com/daostas/sdb"
	"github.com/daostas/slogger"
)

func ConfigureSdb(db sdb.Sdb, logger slogger.Logger, log bool, funcName string, traceId string, req interface{}) (slogger.Logger, sdb.Sdb) {
	l := logger
	l.AddToPrefix("[" + funcName + " Method]")
	l.AddToPrefix("[" + traceId + "]")
	l.Line()
	l.Info(req)

	return l, *db.Set(log, l)
}
