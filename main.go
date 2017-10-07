package main

import (
  "sync"
  "time"
)

type (
  Push struct {
  }
  Stats struct {
	  sync.RWMutex

	  //ElapsedTime ...
	  //Methods ...
	  //AppID ...
	  //...
  }
)

func send(push Push) {
	if !doSmth(push) {
		resendChan <- push
	}
}

func main() {
  for {
	  select {
	  case push := <-mainChan:
		  send(push)
	  case push := <-resendChan:
		  send(push)
	  default:
		  // ...
	  }
  }
  addStatsTicker := time.Tick(5 * time.Second)
  for {
	  select {
	  case <-addStatsTicker:
		  globalStats.Lock()
		  gcm.stats.Lock()
		  mergeStatsToGlobal(&gcm.stats)
		  cleanStats(&gcm.stats)
		  gcm.stats.Unlock()
		  globalStats.Unlock()
		
	  case push := <-mainChan:
		  // таких статистик много, это пример одной из них
		  gcm.stats.Lock()
		  statsMethodIncr(&gcm.stats, push.Method)
		  statsAppIDIncr(&gcm.stats, push.AppID)
		  gcm.stats.Unlock()

		  send(push)
    	// ...
	  }
  }
}
