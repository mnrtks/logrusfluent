# logrus-fluent-hook

Fluent Hooks for [Logrus](https://github.com/Sirupsen/logrus)

This is a hook to send the log to fluentd.

## Usage

```go
import (
	log "github.com/Sirupsen/logrus"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/mnrtks/logrusfluent"
)

func main() {
	fluentConf := fluent.Config{}
	hook, err := logrusfluent.NewFluentHook(fluentConf)
	if err == nil {
		hook.DefaultTag = "logrusfluent"
		log.AddHook(hook)
	}

	log.WithFields(log.Fields{
		"tag":    "hoge",
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// The default tag is used if you do not specify a tag.
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}
```
