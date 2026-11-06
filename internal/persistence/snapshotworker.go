package persistence

import (
	"time"
)



type SnapshotWorker struct{
	Frequency time.Duration
	Persist *Persist
}

func NewSnapshotWorker(freq time.Duration, persist *Persist) *SnapshotWorker{
	return &SnapshotWorker{
		Frequency: freq,
		Persist: persist,
	}
}

func (s *SnapshotWorker) SnapshotWorker(){
	ticker := time.NewTicker(s.Frequency)

	for {
		<-ticker.C
		s.Persist.SaveSnapshot()

	}
}