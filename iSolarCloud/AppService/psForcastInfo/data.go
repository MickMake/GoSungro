package psForcastInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)


const Url = "/v1/powerStationService/psForcastInfo"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	AreaForcastList []struct {
		Chill             valueTypes.Float   `json:"chill"`
		City              valueTypes.String   `json:"city"`
		Code              valueTypes.Integer  `json:"code"`
		CodeName          valueTypes.String   `json:"code_name"`
		DateTime          valueTypes.DateTime `json:"date_time" PointIgnore:"true" PointNameDateFormat:"2006/01/02 15:04:05"`
		Direction         valueTypes.Float    `json:"direction"`
		High              valueTypes.Float    `json:"high" PointUnit:"F"`
		Highc             valueTypes.Float    `json:"highc" PointUnit:"C"`
		Humidity          valueTypes.Float    `json:"humidity"`
		Low               valueTypes.Float    `json:"low PointUnit:"F""`
		Lowc              valueTypes.Float    `json:"lowc" PointUnit:"C"`
		Pressure          valueTypes.Float    `json:"pressure" PointUnit:"hPa"`
		PsKnowledge       valueTypes.String   `json:"ps_knowledge"`
		Rising            valueTypes.Bool     `json:"rising"`

		Speed             valueTypes.Float    `json:"speed" PointUnitFrom:"SpeedUnit"`
		SpeedUnit         valueTypes.String   `json:"speed_unit"`

		SpeedOriginal     valueTypes.Float    `json:"speed_original" PointUnitFrom:"SpeedOriginalUnit"`
		SpeedOriginalUnit valueTypes.String   `json:"speed_original_unit" PointIgnore:"true"`

		Sunrise           valueTypes.Time     `json:"sunrise"`
		Sunset            valueTypes.Time     `json:"sunset"`
		Visibility        valueTypes.Float    `json:"visibility"`
		WeatherDesc       valueTypes.String   `json:"weather_desc"`
		WeatherURL        valueTypes.String   `json:"weather_url"`
	} `json:"areaForcastList" PointIdFromChild:"DateTime" PointNameDateFormat:"20060102" DataTable:"true"`
	StationsCityCode []struct {
		City   valueTypes.String  `json:"city"`
		PsId   valueTypes.PsId `json:"ps_id"`
		PsName valueTypes.String  `json:"ps_name"`
	} `json:"stationsCityCode" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

		// for _, v := range e.Response.ResultData.AreaForcastList {
		// 	entries.StructToDataMap(v, api.JoinWithDots(0, valueTypes.DateTimeLayoutDay, pkg, v.DateTime), e.Request.PsId.String(), valueTypes.NewDateTime(""))
		// }
	}

	return entries
}
