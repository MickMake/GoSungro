package getPowerDevicePointInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/reportService/getPowerDevicePointInfo"
const Disabled = false

type RequestData struct {
	Id valueTypes.Integer `json:"id" required:"true"`
	// Id string `json:"id"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	DeviceType    valueTypes.Integer  `json:"device_type" PointId:"device_type" PointValueType:"" PointTimeSpan:""`
	ID            valueTypes.Integer  `json:"id" PointId:"id" PointValueType:"" PointTimeSpan:""`
	Period        valueTypes.Integer  `json:"period" PointId:"period" PointValueType:"" PointTimeSpan:""`
	PointID       valueTypes.Integer  `json:"point_id" PointId:"point_id" PointValueType:"" PointTimeSpan:""`
	PointName     string `json:"point_name" PointId:"point_name" PointValueType:"" PointTimeSpan:""`
	ShowPointName string `json:"show_point_name" PointId:"show_point_name" PointValueType:"" PointTimeSpan:""`
	TranslationID valueTypes.Integer  `json:"translation_id" PointId:"translation_id" PointValueType:"" PointTimeSpan:""`
}

func (e *ResultData) IsValid() error {
	var err error
		// switch {
		// 	case e.Dummy == "":
		// 		break
		// 	default:
		// 		err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
		// }
	return err
}

func (e *EndPoint) GetData2() ResultData {
	return e.Response.ResultData
}

// func (e *EndPoint) GetData() api.DataMap {
// 	return e.Response.ResultData.GetData()
// }
//
// func (e *ResultData) GetData() api.DataMap {
// 	entries := api.NewDataMap()
//
// 	for range Only.Once {
// 		entries.StructToPoints("", *e)
// 	}
// 	return entries
// }

func (e *EndPoint) AddDataTable(table output.Table) output.Table {

	for range Only.Once {
		rd := e.Response.ResultData

		if rd.ID.Value() == 0 {
			break
		}
		_ = table.AddRow(rd.DeviceType, rd.ID, rd.Period, rd.PointID,rd.PointName, rd.ShowPointName, rd.TranslationID)
	}
	return table
}

func (e *EndPoint) GetPointDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"DeviceType",
			"Id",
			"Period",
			"Point Id",
			"Point Name",
			"Show Point Name",
			"Translation Id",
		)
		rd := e.Response.ResultData
		_ = table.AddRow(rd.DeviceType, rd.ID, rd.Period, rd.PointID,rd.PointName, rd.ShowPointName, rd.TranslationID)
	}
	return table
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
