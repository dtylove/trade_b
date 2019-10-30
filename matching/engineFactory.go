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
	var pList []models.Product
	err := models.FindList(&pList)

	if err != nil {
		panic(err)
	}
	for _, p := range pList {
		data, _ := json.Marshal(p)
		fmt.Println(string(data))
		Engines[p.Id] = InitEngine(p.Id)
		Engines[p.Id].Start()
		fmt.Println(Engines[p.Id])
	}
}

func AddProduct(p *models.Product) {
	Engines[p.Id] = InitEngine(p.Id)
}
