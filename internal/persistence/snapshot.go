package persistence

import (
	"encoding/json"
	"os"
	"redis-lite/internal/store"
	"time"
	"fmt"
	"strconv"
)

type Snapshot interface{
	Save()
	Load(filename string)
}

type Persist struct{
	Store *store.Store
}

func NewPersist(Store *store.Store) *Persist{
	return &Persist{
		Store: Store,
	}
}





func (p *Persist) SaveSnapshot(){
	timeNow := time.Now().Unix()
	fileName := fmt.Sprintf("./dump/%v_%v", p.Store.StoreName, strconv.FormatInt(timeNow, 10))
	file, err := os.Create(string(fileName))
	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(p.Store)
	if err != nil {
		fmt.Println(err)
	}

}


