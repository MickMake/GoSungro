package energyPowerGenerationTrends

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/orgService/energyPowerGenerationTrends"
const Disabled = false
const EndPointName = "MttvScreenService.energyPowerGenerationTrends"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	KpiList []interface{} `json:"kpiList"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
