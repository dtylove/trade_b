package matching

import (
	"dtyTrade/rest/models"
	"encoding/json"
	"fmt"
)

var Engines = make(map[uint]*Engine)

func GetEngine(marketId uint) *Engine {
	return Engines[marketId]
}

func InitEngineFactory() {
	fmt.Println("call InitEngineFactory ")
	pList, err := models.FindProductList()

	data, _ := json.Marshal(pList)
	fmt.Printf(string(data))
	if err != nil {
		panic(err)
	}
	for _, p := range pList {
		data, _ := json.Marshal(p)
		fmt.Println(string(data))
		Engines[p.Id] = InitEngine(p.Id)
		Engines[p.Id].Start()
	}
}
