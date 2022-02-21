// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebIscmAppService

import (
	"GoSungrow/iSolarCloud/api"
	"fmt"
)


var endpoints = [][]string {
	{"WebIscmAppService","addPowerDeviceModel","/v1/devService/getDevSysTypeAndFactoryList"},
	{"WebIscmAppService","addPowerPointManage","/v1/devService/addPowerPointManage"},
	{"WebIscmAppService","addSubTypeDevice","/v1/devService/addSubTypeDevice"},
	{"WebIscmAppService","batchAddDevicesPropertis","/v1/devService/batchAddDevicesPropertis"},
	{"WebIscmAppService","batchDelDevice","/v1/devService/batchDelDevice"},
	{"WebIscmAppService","batchSavePowerDeviceTechnical","/v1/devService/batchSavePowerDeviceTechnical"},
	{"WebIscmAppService","checkDeviceModel","/v1/devService/checkDeviceModel"},
	{"WebIscmAppService","contactMessageOpera","/v1/powerStationService/contactMessageOpera"},
	{"WebIscmAppService","delDevice","/v1/devService/delDevice"},
	{"WebIscmAppService","deleteDeviceFactory","/v1/devService/deleteDeviceFactory"},
	{"WebIscmAppService","deleteDeviceType","/v1/devService/deleteDeviceType"},
	{"WebIscmAppService","deleteMenu","/v1/userService/deleteMenu"},
	{"WebIscmAppService","deleteOneNotice","/v1/otherService/deleteOneNotice"},
	{"WebIscmAppService","deleteOrgNodeInfo","/v1/orgService/deleteOrgNodeInfo"},
	{"WebIscmAppService","deletePicture","/v1/powerStationService/deletePicture"},
	{"WebIscmAppService","deletePointInfo","/v1/devService/deletePointInfo"},
	{"WebIscmAppService","deletePowerDeviceChannl","/v1/devService/deletePowerDeviceChannl"},
	{"WebIscmAppService","deletePowerDeviceModel","/v1/devService/deletePowerDeviceModel"},
	{"WebIscmAppService","deletePowerDeviceParameterPage","/v1/devService/deletePowerDeviceParameterPage"},
	{"WebIscmAppService","deletePowerDeviceSubType","/v1/devService/deletePowerDeviceSubType"},
	{"WebIscmAppService","deletePowerDeviceTechnical","/v1/devService/deletePowerDeviceTechnical"},
	{"WebIscmAppService","deletePowerStore","/v1/otherService/deletePowerStore"},
	{"WebIscmAppService","deleteProcessDefinition","/v1/faultService/deleteProcessDefinition"},
	{"WebIscmAppService","deleteReport","/v1/reportService/deleteReport"},
	{"WebIscmAppService","deleteUserNode","/v1/orgService/deleteUserNode"},
	{"WebIscmAppService","deployProcess","/v1/faultService/deployProcess"},
	{"WebIscmAppService","editProcessManageAction","/v1/userService/editProcessManageAction"},
	{"WebIscmAppService","findImageInputStreamString","/v1/faultService/findImageInputStreamString"},
	{"WebIscmAppService","getAllDevTypeList","/v1/devService/getAllDevTypeList"},
	{"WebIscmAppService","getAllNodeByType","/v1/faultService/getAllNodeByType"},
	{"WebIscmAppService","getAuthKey","/v1/powerStationService/getAuthKey"},
	{"WebIscmAppService","getAuthKeyList","/v1/devService/getAuthKeyList"},
	{"WebIscmAppService","getCodeByType","/v1/powerStationService/getCodeByType"},
	{"WebIscmAppService","getContactMessage","/v1/powerStationService/getContactMessage"},
	{"WebIscmAppService","getCountryNew","/v1/commonService/getCountryNew"},
	{"WebIscmAppService","getDefinitionIdByKey","/v1/faultService/getDefinitionIdByKey"},
	{"WebIscmAppService","getDeploymentList","/v1/faultService/getDeploymentList"},
	{"WebIscmAppService","getDeviceFactoryListByIds","/v1/devService/getDeviceFactoryListByIds"},
	{"WebIscmAppService","getDeviceModel","/v1/devService/getDeviceModel"},
	{"WebIscmAppService","getDevicePro","/v1/devService/getDevicePro"},
	{"WebIscmAppService","getDeviceSubType","/v1/devService/getDeviceSubType"},
	{"WebIscmAppService","getDeviceTechnical","/v1/devService/getDeviceTechnical"},
	{"WebIscmAppService","getDeviceType","/v1/devService/getDeviceType"},
	{"WebIscmAppService","getDeviceTypeInfoById","/v1/devService/getDeviceTypeInfoById"},
	{"WebIscmAppService","getDutyUserList","/v1/userService/getDutyUserList"},
	{"WebIscmAppService","getFatherPrivileges","/v1/userService/getFatherPrivileges"},
	{"WebIscmAppService","getGroupManSettings","/v1/faultService/getGroupManSettings"},
	{"WebIscmAppService","getGroupManSettingsMembers","/v1/faultService/getGroupManSettingsMembers"},
	{"WebIscmAppService","getMaterialByListId","/v1/devService/getMaterialByListId"},
	{"WebIscmAppService","getMaterialByType","/v1/devService/getMaterialByType"},
	{"WebIscmAppService","getMaterialList","/v1/devService/getMaterialList"},
	{"WebIscmAppService","getMaxDeviceIdByPsId","/v1/devService/getMaxDeviceIdByPsId"},
	{"WebIscmAppService","getModelPoints","/v1/devService/getModelPoints"},
	{"WebIscmAppService","getMoneyUnitList","/v1/commonService/getMoneyUnitList"},
	{"WebIscmAppService","getNamecnNew","/v1/commonService/getNamecnNew"},
	{"WebIscmAppService","getNationList","/v1/commonService/getNationList"},
	{"WebIscmAppService","getOperationRecord","/v1/commonService/getOperationRecord"},
	{"WebIscmAppService","getOrgAndChildBasicInfoOptions","/v1/orgService/getOrgAndChildBasicInfoOptions"},
	{"WebIscmAppService","getOrgAndStateAndCode","/v1/userService/getOrgAndStateAndCode"},
	{"WebIscmAppService","getOrgForPs","/v1/powerStationService/getOrgForPs"},
	{"WebIscmAppService","getOrgList","/v1/orgService/getOrgList"},
	{"WebIscmAppService","getOrgListForUser","/v1/orgService/getOrgListForUser"},
	{"WebIscmAppService","getOrgNodeInfo","/v1/orgService/getOrgNodeInfo"},
	{"WebIscmAppService","getOrgStationList","/v1/orgService/getOrgStationList"},
	{"WebIscmAppService","getOrgStationListByPage","/v1/orgService/getOrgStationListByPage"},
	{"WebIscmAppService","getOrgUserList","/v1/userService/getOrgUserList"},
	{"WebIscmAppService","getOrgUserMapData","/v1/faultService/getOrgUserMapData"},
	{"WebIscmAppService","getOrgZtree","/v1/orgService/getOrgZtree"},
	{"WebIscmAppService","getOrgZtree4User","/v1/orgService/getOrgZtree4User"},
	{"WebIscmAppService","getOrgZtreeAsync","/v1/orgService/getOrgZtreeAsync"},
	{"WebIscmAppService","getOrgZtreeForUser","/v1/orgService/getOrgZtreeForUser"},
	{"WebIscmAppService","getPictureList","/v1/powerStationService/getPictureList"},
	{"WebIscmAppService","getPointInfo","/v1/devService/getPointInfo"},
	{"WebIscmAppService","getPointInfoPage","/v1/devService/getPointInfoPage"},
	{"WebIscmAppService","getPowerDevice","/v1/devService/getPowerDevice"},
	{"WebIscmAppService","getPowerDeviceChannl","/v1/devService/getPowerDeviceChannl"},
	{"WebIscmAppService","getPowerDeviceFactory","/v1/devService/getPowerDeviceFactory"},
	{"WebIscmAppService","getPowerDeviceFactoryListCount","/v1/devService/getPowerDeviceFactoryListCount"},
	{"WebIscmAppService","getPowerDeviceInfo","/v1/devService/getPowerDeviceInfo"},
	{"WebIscmAppService","getPowerDeviceModelList","/v1/devService/getPowerDeviceModelList"},
	{"WebIscmAppService","getPowerDeviceModelTechList","/v1/devService/getPowerDeviceModelTechList"},
	{"WebIscmAppService","getPowerDeviceTypeList","/v1/devService/getPowerDeviceTypeList"},
	{"WebIscmAppService","getPowerPlanList","/v1/powerStationService/getPowerPlanList"},
	{"WebIscmAppService","getPowerStation","/v1/otherService/getPowerStationStore"},
	{"WebIscmAppService","getPowerStationInfo","/v1/powerStationService/getPowerStationInfoForBackSys"},
	{"WebIscmAppService","getPowerStationList","/v1/powerStationService/getPowerStationList"},
	{"WebIscmAppService","getPowerStore","/v1/otherService/getPowerStore"},
	{"WebIscmAppService","getProvcnNew","/v1/commonService/getProvcnNew"},
	{"WebIscmAppService","getPsTreeMenu","/v1/devService/getPsTreeMenu"},
	{"WebIscmAppService","getRoleByUserIds","/v1/userService/getRoleByUserIds"},
	{"WebIscmAppService","getRootOrgInfoByUserId","/v1/userService/getRootOrgInfoByUserId"},
	{"WebIscmAppService","getSettingUserMapData","/v1/faultService/getSettingUserMapData"},
	{"WebIscmAppService","getStateNew","/v1/commonService/getStateNew"},
	{"WebIscmAppService","getSubTypeDevice","/v1/devService/getSubTypeDevice"},
	{"WebIscmAppService","getSysHomeList2","/v1/userService/getSysHomeList2"},
	{"WebIscmAppService","getSysMenu","/v1/userService/getSysMenu"},
	{"WebIscmAppService","getSysOrgPro","/v1/orgService/getSysOrgPro"},
	{"WebIscmAppService","getSysUser","/v1/userService/getSysUser"},
	{"WebIscmAppService","getSystemOrgInfo","/v1/orgService/getSystemOrgInfo"},
	{"WebIscmAppService","getSystemRoleInfo","/v1/userService/getSystemRoleInfo"},
	{"WebIscmAppService","getSystemRoleList2","/v1/userService/getSystemRoleList2"},
	{"WebIscmAppService","getTownValueNew","/v1/commonService/getTownValueNew"},
	{"WebIscmAppService","getUserMenuLs","/v1/userService/getUserMenuLs"},
	{"WebIscmAppService","getUserOrgPage","/v1/orgService/getUserOrgPage"},
	{"WebIscmAppService","getVillageList","/v1/commonService/getVillageList"},
	{"WebIscmAppService","getVillageListNew","/v1/commonService/getVillageListNew"},
	{"WebIscmAppService","getZtreeAsyncSysMenu","/v1/userService/getZtreeAsyncSysMenu"},
	{"WebIscmAppService","getZtreeChildMenu","/v1/devService/getZtreeChildMenu"},
	{"WebIscmAppService","getZtreeMenu","/v1/devService/getZtreeMenu"},
	{"WebIscmAppService","getZtreeSysMenu","/v1/userService/getZtreeSysMenu"},
	{"WebIscmAppService","getZtreeSysMenu2","/v1/userService/getZtreeSysMenu2"},
	{"WebIscmAppService","goToDevicePropertyPage","/v1/devService/goToDevicePropertyPage"},
	{"WebIscmAppService","isCanAddUser","/v1/userService/isCanAddUser"},
	{"WebIscmAppService","isHasIrradiationData","/v1/powerStationService/isHasIrradiationData"},
	{"WebIscmAppService","isHasPlan","/v1/powerStationService/isHasPlan"},
	{"WebIscmAppService","loadDevice","/v1/devService/loadDevice"},
	{"WebIscmAppService","modelPointsPage","/v1/devService/modelPointsPage"},
	{"WebIscmAppService","modifyDevice","/v1/devService/modifyDevicePre"},
	{"WebIscmAppService","modifyPowerDeviceChannl","/v1/devService/modifyPowerDeviceChannl"},
	{"WebIscmAppService","modifySysOrg","/v1/orgService/modifySysOrg"},
	{"WebIscmAppService","modifySystemMenu","/v1/userService/modifySystemMenu"},
	{"WebIscmAppService","modifySystemOrgNode","/v1/orgService/modifySystemOrgNode"},
	{"WebIscmAppService","modifySystemRole","/v1/userService/modifySystemRole"},
	{"WebIscmAppService","modifySystemUser","/v1/userService/modifySystemUser"},
	{"WebIscmAppService","publishNotice","/v1/otherService/publishNotice"},
	{"WebIscmAppService","queryDeviceList","/v1/devService/queryDeviceListForBackSys"},
	{"WebIscmAppService","queryDutyType","/v1/otherService/queryDutyType"},
	{"WebIscmAppService","queryReportDataById","/v1/reportService/queryReportDataById"},
	{"WebIscmAppService","resetPasW","/v1/userService/resetPasW"},
	{"WebIscmAppService","saveAuthKey","/v1/devService/saveAuthKey"},
	{"WebIscmAppService","saveDevice","/v1/devService/saveDevice"},
	{"WebIscmAppService","saveDeviceFactory","/v1/devService/saveDeviceFactory"},
	{"WebIscmAppService","saveDeviceType","/v1/devService/saveDeviceType"},
	{"WebIscmAppService","saveIrradiationData","/v1/powerStationService/saveIrradiationData"},
	{"WebIscmAppService","saveModelPoints","/v1/devService/saveModelPoints"},
	{"WebIscmAppService","saveNewNotice","/v1/otherService/saveNewNotice"},
	{"WebIscmAppService","saveOrUpdateReport","/v1/reportService/saveOrUpdateReport"},
	{"WebIscmAppService","saveOrgNode","/v1/orgService/saveOrgNode"},
	{"WebIscmAppService","saveOrgUsers","/v1/faultService/saveOrgUsers"},
	{"WebIscmAppService","savePicture","/v1/powerStationService/savePicture"},
	{"WebIscmAppService","savePointManage","/v1/devService/savePointManage"},
	{"WebIscmAppService","savePowerDeviceChannl","/v1/devService/savePowerDeviceChannl"},
	{"WebIscmAppService","savePowerDeviceModel","/v1/devService/savePowerDeviceModel"},
	{"WebIscmAppService","savePowerDeviceParameterPage","/v1/devService/savePowerDeviceParameterPage"},
	{"WebIscmAppService","savePowerDeviceSubType","/v1/devService/savePowerDeviceSubType"},
	{"WebIscmAppService","savePowerDeviceTechnical","/v1/devService/savePowerDeviceTechnical"},
	{"WebIscmAppService","savePowerPlan","/v1/powerStationService/savePowerPlan"},
	{"WebIscmAppService","savePowerStationByPowerStore","/v1/otherService/savePowerStationByPowerStore"},
	{"WebIscmAppService","savePowerStore","/v1/otherService/savePowerStore"},
	{"WebIscmAppService","savePsOrg","/v1/powerStationService/savePsOrg"},
	{"WebIscmAppService","saveRelDevice","/v1/devService/saveRelDevice"},
	{"WebIscmAppService","saveRoleAssign","/v1/userService/saveRoleAssign"},
	{"WebIscmAppService","saveSysMenu","/v1/userService/saveSysMenu"},
	{"WebIscmAppService","saveSysOrg","/v1/orgService/saveSysOrg"},
	{"WebIscmAppService","saveSysRole","/v1/userService/saveSysRole"},
	{"WebIscmAppService","saveSysUser","/v1/userService/saveSysUser"},
	{"WebIscmAppService","saveUserNode","/v1/orgService/saveUserNode"},
	{"WebIscmAppService","saveUserRole","/v1/userService/saveUserRole"},
	{"WebIscmAppService","searchIrradiationData","/v1/powerStationService/searchIrradiationData"},
	{"WebIscmAppService","searchTechnicalNums","/v1/devService/searchTechnicalNums"},
	{"WebIscmAppService","selectDeviceTypeByPsId","/v1/devService/selectDeviceTypeByPsId"},
	{"WebIscmAppService","selectPowerDeviceTechnicals","/v1/devService/selectPowerDeviceTechnicals"},
	{"WebIscmAppService","selectPowerDeviceType","/v1/devService/selectPowerDeviceType"},
	{"WebIscmAppService","setupUserRole4AddUser","/v1/userService/setupUserRole4AddUser"},
	{"WebIscmAppService","startWorkFlow","/v1/faultService/startWorkFlow"},
	{"WebIscmAppService","updateDevice","/v1/devService/updateDevice"},
	{"WebIscmAppService","updateDeviceType","/v1/devService/updateDeviceType"},
	{"WebIscmAppService","updateFaultLevel","/v1/faultService/updateFaultLevel"},
	{"WebIscmAppService","updateNotice","/v1/otherService/updateNotice"},
	{"WebIscmAppService","updatePointInfo","/v1/devService/updatePointInfo"},
	{"WebIscmAppService","updatePowerDeviceModel","/v1/devService/updatePowerDeviceModel"},
	{"WebIscmAppService","updatePowerDeviceParameterPage","/v1/devService/updatePowerDeviceParameterPage"},
	{"WebIscmAppService","updatePowerDeviceSubType","/v1/devService/updatePowerDeviceSubType"},
	{"WebIscmAppService","updatePowerDeviceTechnical","/v1/devService/updatePowerDeviceTechnical"},
	{"WebIscmAppService","updateProcessManage","/v1/userService/updateProcessManage"},
	{"WebIscmAppService","updateSysOrgPro","/v1/orgService/updateSysOrgPro"},
	{"WebIscmAppService","updateSysRoleValidFlag","/v1/userService/updateSysRoleValidFlag"},
	{"WebIscmAppService","updateUserValidFlag","/v1/userService/updateUserValidFlag"},
	{"WebIscmAppService","updateValidFlag","/v1/devService/updateValidFlag"},
	{"WebIscmAppService","viewDeviceModel","/v1/devService/viewDeviceModel"},
	{"WebIscmAppService","viewDeviceParameter","/v1/devService/viewDeviceParameter"},
	{"WebIscmAppService","workFlowImplementStep","/v1/faultService/workFlowImplementStep"},
	{"WebIscmAppService","workFlowIsStart","/v1/faultService/workFlowIsStart"},
	{"WebIscmAppService","workFlowTransferStep","/v1/faultService/workFlowTransferStep"},
}


var _ api.Area = (*Area)(nil)

type Area api.AreaStruct


func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot *api.Web) Area {
	area := Area {
		ApiRoot:   apiRoot,
		Name:      api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints {},
	}

	return area
}


// ****************************************
// Methods not scoped by api.EndPoint interface type

func GetAreaName() string {
	return string(api.GetArea(Area{}))
}

func (a Area) GetEndPoint(name api.EndPointName) api.EndPoint {
	var ret api.EndPoint
	for _, e := range a.EndPoints {
		// fmt.Printf("endpoint: %v\n", e)
		if e.GetName() == name {
			ret = e
			break
		}
	}
	return ret
}


// ****************************************
// Methods scoped by api.Area interface type

func (a Area) Init(apiRoot *api.Web) api.AreaStruct {
	panic("implement me")
}

func (a Area) GetAreaName() api.AreaName {
	return a.Name
}

func (a Area) GetEndPoints() api.TypeEndPoints {
	for _, endpoint := range a.EndPoints {
		fmt.Printf("endpoint: %v\n", endpoint)
	}
	return a.EndPoints
}

func (a Area) Call(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

func (a Area) GetRequest(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) GetResponse(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) GetData(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
