package example

import (
	"github.com/dereulenspiegel/freeradius-go"
)

func CreateModule() freeradius.Module {
	return &exampleModule{}
}

type exampleModule struct {
	radlog freeradius.Log
}

func (e *exampleModule) Init(logger freeradius.Log) error {
	e.radlog = logger
	e.radlog.Radlog(freeradius.LogTypeInfo, "Init from go plugin")
	return nil
}
