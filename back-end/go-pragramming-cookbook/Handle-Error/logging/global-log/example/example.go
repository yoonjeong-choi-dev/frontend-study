package main

import (
	global_log "global-log"
)

func main() {
	if err := global_log.Init(); err != nil {
		panic(err)
	}

	global_log.Debug("before adding key")
	global_log.WithField("test-key", "test-val").Debug("added test-key")
	global_log.WithField("name", "yoonjeong").Debug("name key added")
	global_log.Debug("after adding key")
}
