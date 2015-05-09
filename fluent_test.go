package logrusfluent

import (
	"github.com/Sirupsen/logrus"
	"github.com/fluent/fluent-logger-golang/fluent"
	"testing"
)

func TestFluentHook(t *testing.T) {
	log := logrus.New()
	hook, err := NewFluentHook(fluent.Config{})

	if err != nil {
		t.Errorf("Unable to create hook.")
	}

	log.Hooks.Add(hook)

	for _, level := range hook.Levels() {
		if len(log.Hooks[level]) != 1 {
			t.Errorf("SyslogHook was not added. The length of log.Hooks[%v]: %v", level, len(log.Hooks[level]))
		}
	}

	log.WithFields(logrus.Fields{
		"key1": "value1",
		"key2": "value2",
	}).Info("hgoe")
}
