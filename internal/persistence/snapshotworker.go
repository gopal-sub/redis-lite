package persistence

import (
	"time"
	"fmt"
	"os"
	"encoding/json"
	"strconv"
	"redis-lite/internal/snapshot"
)



type SnapshotWorker struct{
	//SnapshotWorker maintains Persist for each store => live map
	Frequency time.Duration
	StoresPersist []*Persist
	LastSnapshot time.Time
}

func NewSnapshotWorker(freq time.Duration) *SnapshotWorker{
	return &SnapshotWorker{
		Frequency: freq,
	}
}

func (s *SnapshotWorker)AppendStorePersistenceToSnapShotWorker(persist *Persist){
	s.StoresPersist = append(s.StoresPersist, persist)
}

func (s *SnapshotWorker) SnapshotWorker(){
	ticker := time.NewTicker(s.Frequency)

	for {
		<-ticker.C
		databaseSnapshot := snapshot.NewDatabaseSnapshot()
		for _, persist := range s.StoresPersist{
			storeSnap := persist.CopyLiveStoreDataToBuffer()
			databaseSnapshot.AddStoreSnapToDatabaseSnapshot(storeSnap)

		}

		timeNow := time.Now().Unix()
		fileName := fmt.Sprintf("./dump/%v.dump", strconv.FormatInt(timeNow, 10))
		file, err := os.Create(string(fileName))
		if err != nil {
			fmt.Println(err)
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(databaseSnapshot)
		if err != nil {
			fmt.Println(err)
		}

		}
}