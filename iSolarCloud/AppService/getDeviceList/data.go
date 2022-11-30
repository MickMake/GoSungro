package getDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceList"
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
	PageList []struct {
		GoStructParent          GoStruct.GoStructParent   `json:"-" PointDeviceFrom:"PsKey"`

		AttrId                  valueTypes.Integer  `json:"attr_id"`
		ChannelId               valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		CommandStatus           valueTypes.Integer  `json:"command_status"`
		ConnectState            valueTypes.Integer  `json:"connect_state"`
		DataFlag                valueTypes.Integer  `json:"data_flag"`
		DataFlagDetail          valueTypes.Integer  `json:"data_flag_detail"`
		DevFaultStatus          valueTypes.Integer  `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer  `json:"dev_status"`
		DeviceArea              valueTypes.Integer  `json:"device_area"`
		DeviceCode              valueTypes.Integer  `json:"device_code"`
		DeviceFactoryDate       valueTypes.DateTime `json:"device_factory_date" PointNameDateFormat:"2006/01/02 15:04:05"`
		DeviceId                valueTypes.Integer  `json:"device_id"`
		DeviceModel             valueTypes.String   `json:"device_model"`
		DeviceModelCode         valueTypes.String   `json:"device_model_code"`
		DeviceModelId           valueTypes.Integer  `json:"device_model_id"`
		DeviceName              valueTypes.String   `json:"device_name"`
		DeviceProSn             valueTypes.String   `json:"device_pro_sn" PointName:"Device Serial Number"`
		DeviceState             valueTypes.Integer  `json:"device_state"`
		DeviceSubType           interface{}         `json:"device_sub_type"`
		DeviceSubTypeName       interface{}         `json:"device_sub_type_name"`
		DeviceType              valueTypes.Integer  `json:"device_type"`
		FactoryName             valueTypes.String   `json:"factory_name"`
		InstallerDevFaultStatus valueTypes.Integer  `json:"installer_dev_fault_status"`
		InverterModelType       valueTypes.Integer  `json:"inverter_model_type"`
		IsCountryCheck          valueTypes.Bool     `json:"is_country_check"`
		IsHasFunctionEnum       valueTypes.Bool     `json:"is_has_function_enum"`
		IsHasTheAbility         valueTypes.Bool     `json:"is_has_the_ability"`
		IsInit                  valueTypes.Bool     `json:"is_init"`
		IsReadSet               valueTypes.Bool     `json:"is_read_set"`
		IsReplacing             valueTypes.Bool     `json:"is_replacing"`
		IsReset                 valueTypes.Bool     `json:"is_reset"`
		IsSecond                valueTypes.Bool     `json:"is_second"`
		IsThirdParty            valueTypes.Bool     `json:"is_third_party"`
		ModuleUUID              valueTypes.Integer  `json:"module_uuid"`
		OwnerDevFaultStatus     valueTypes.Integer  `json:"owner_dev_fault_status"`
		P24                     interface{}         `json:"p24"`
		Posx                    interface{}         `json:"posx"`
		Posy                    interface{}         `json:"posy"`
		PsId                    valueTypes.PsId     `json:"ps_id"`
		PsKey                   valueTypes.PsKey    `json:"ps_key"`
		RelState                valueTypes.Integer  `json:"rel_state"`
		Sn                      valueTypes.String   `json:"sn" PointName:"Serial Number"`
		TypeName                valueTypes.String   `json:"type_name"`
		UUID                    valueTypes.Integer  `json:"uuid"`
	} `json:"pageList" PointId:"page_list" PointIdFromChild:"PsKey" PointIdReplace:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

type Device struct {
	Vendor        valueTypes.String
	PsId          valueTypes.PsId
	PsKey         valueTypes.PsKey
	DeviceName    valueTypes.String
	DeviceProSn   valueTypes.String
	DeviceModel   valueTypes.String
	DeviceType    valueTypes.Integer
	DeviceCode    valueTypes.Integer
	ChannelId     valueTypes.Integer
	DeviceModelId valueTypes.Integer
	TypeName      valueTypes.String
	DeviceState   valueTypes.Integer
	DevStatus     valueTypes.Integer
	Uuid          valueTypes.Integer
}
type Devices []Device

func (e *EndPoint) GetDevices() Devices {
	var ret Devices
	for _, d := range e.Response.ResultData.PageList {
		ret = append(ret, Device{
			Vendor:        d.FactoryName,
			PsKey:         d.PsKey,
			PsId:          d.PsId,
			DeviceType:    d.DeviceType,
			DeviceCode:    d.DeviceCode,
			ChannelId:     d.ChannelId,
			TypeName:      d.TypeName,
			DeviceProSn:   d.DeviceProSn,
			DeviceModel:   d.DeviceModel,
			DeviceModelId: d.DeviceModelId,
			DeviceName:    d.DeviceName,
			DeviceState:   d.DeviceState,
			DevStatus:     d.DevStatus,
			Uuid:          d.ModuleUUID,
		})
	}
	return ret
}

func (e *EndPoint) GetDevicesTable() output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsId.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelId.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }

		data := e.GetDevices()
		table = GetDevicesTable(data)
	}
	return table
}

func GetDevicesTable(data Devices) output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsId.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelId.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }

		table = output.NewTable(
			"Vendor",
			"Ps Key",
			"Ps Id",
			"Type",
			"Code",
			"Id",
			"Type Name",
			"Serial Number",
			"Model",
			"Model Id",
			"Name",
			"State",
			"Status",
			"UUID",
		)
		table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Vendor",
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	"UUID",
		// )

		for _, d := range data {
			_ = table.AddRow(d.Vendor.String(),
				d.PsKey.String(),
				d.PsId.String(),
				d.DeviceType.String(),
				d.DeviceCode.String(),
				d.ChannelId.String(),
				d.TypeName.String(),
				d.DeviceProSn.String(),
				d.DeviceModel.String(),
				d.DeviceModelId.String(),
				d.DeviceName.String(),
				d.DeviceState.String(),
				d.DevStatus.String(),
				d.Uuid.String(),
			)
		}
	}
	return table
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
