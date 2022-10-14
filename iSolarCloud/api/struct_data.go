package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"fmt"
	datatable "go.pennock.tech/tabular/auto"
	"strings"
	"time"
)


const (
	PointTimeSpanInstant = "instant"
	PointTimeSpanBoot    = "boot"
	PointTimeSpanDaily   = "daily"
	PointTimeSpanMonthly = "monthly"
	PointTimeSpanYearly  = "yearly"
	PointTimeSpanTotal   = "total"
)

type DataPointEntries []DataEntry

func (de *DataPointEntries) Hide() {
	for range Only.Once {
		for i := range *de {
			(*de)[i].Hide = true
		}
	}
}


type DataMap struct {
	DataPoints map[string]DataPointEntries
	Order      []string
}

type DataEntry struct {
	Point *Point              `json:"point"`
	Date  valueTypes.DateTime `json:"date"`

	EndPoint   string               `json:"endpoint"`
	// FullId     valueTypes.DataPoint `json:"full_id"`
	Parent     ParentDevice         `json:"parent"`
	Value      string               `json:"value"`
	ValueFloat float64              `json:"value_float"`
	ValueBool  bool                 `json:"value_bool"`
	Index      int                  `json:"index"`
	Valid      bool                 `json:"valid"`
	Hide       bool                 `json:"hide"`
}

func (de *DataEntry) IsValid() bool {
	var ok bool
	for range Only.Once {
		if de.Point == nil {
			break
		}
		if de.Point.Valid == false {
			break
		}
		ok = true
	}
	return ok
}
func (de *DataEntry) IsNotValid() bool {
	return !de.IsValid()
}

func (dm *DataEntry) FullId() string {
	return dm.EndPoint + "." + dm.Point.Id.String()
}


func NewDataMap() DataMap {
	return DataMap {
		DataPoints: make(map[string]DataPointEntries),
	}
}

// func (dm *DataMap) Add(point string, entry DataEntry) {
// 	dm.Entries[point] = entry
// 	dm.Order = append(dm.Order, point)
// }


func (dm *DataMap) StructToPoints(ref interface{}, endpoint string, parentId string, timestamp valueTypes.DateTime) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		// Iterate over all available fields and read the tags value
		tp := apiReflect.GetPointTags(ref, endpoint)
		// fmt.Printf("TP: %v\n", tp)

		for _, f := range tp {
			if f.PointIgnore {
				// fmt.Printf("IGNORE: %s\n", f.PointId)
				continue
			}

			if f.PointName == "" {
				f.PointName = valueTypes.PointToName(f.PointId)
			}

			if f.PointDevice == "" {
				if parentId != "" {
					f.PointDevice = parentId
				} else {
					f.PointDevice = "virtual"
				}
			}
			var parents ParentDevices
			parents.Add(ParentDevice{Key: f.PointDevice})

			// fmt.Printf("DEBUG: StructToPoints(): %s / %s\n", f.Endpoint, f.PointId)
			uv, _, ok := valueTypes.AnyToUnitValue(f.Value, f.PointUnit, f.PointValueType)
			if !ok {
				continue
			}

			switch f.PointTimeSpan {
				case "PointTimeSpanInstant":
					f.PointTimeSpan = PointTimeSpanInstant
				case "PointTimeSpanBoot":
					f.PointTimeSpan = PointTimeSpanBoot
				case "PointTimeSpanDaily":
					f.PointTimeSpan = PointTimeSpanDaily
				case "PointTimeSpanMonthly":
					f.PointTimeSpan = PointTimeSpanMonthly
				case "PointTimeSpanYearly":
					f.PointTimeSpan = PointTimeSpanYearly
				case "PointTimeSpanTotal":
					f.PointTimeSpan = PointTimeSpanTotal
			}

			var now valueTypes.DateTime
			if timestamp.IsZero() {
				now = valueTypes.NewDateTime(time.Now().Round(5 * time.Minute).Format(valueTypes.DateTimeLayoutZeroSeconds))
			} else {
				now = valueTypes.NewDateTime(timestamp.String())
			}

			p := Point {
				Parents:   parents,
				Id:        valueTypes.SetPointIdString(f.PointId),
				GroupName: f.PointGroupName,
				Name:      f.PointName,
				Unit:      uv.Unit(),
				TimeSpan:  f.PointTimeSpan,
				ValueType: uv.Type(),
				Valid:     true,
				States:    nil,
			}
			p.FixUnitType()

			// Add arrays as multiple entries.
			if len(uv) > 1 {
				// @TODO - Think about adding in arrays of values OR just marshal arrays into JSON.
				res := valueTypes.SizeOfArrayLength(uv)
				for i, val := range uv {
					dm.AddEntry(JoinWithDots(res, valueTypes.DateTimeLayoutDay, f.Endpoint, i), f.PointDevice, p, now, val.String())
					// dm.AddEntry(fmt.Sprintf(res, f.Endpoint, i), f.PointDevice, p, now, val.String())
				}
				continue
			}

			dm.AddEntry(f.Endpoint, f.PointDevice, p, now, uv[0].String())
			if f.PointAlias != "" {
				// fullName = NameDevicePoint(device, PointId(alias))
				p.Id = valueTypes.SetPointIdString(f.PointAlias)
				dm.AddEntry(f.Endpoint, f.PointDevice, p, now, uv[0].String())
			}
		}

		for _, f := range tp {
			// Reference Units from another data point.
			if f.PointUnitFrom != "" {
				sdp := dm.GetEntryFromPointId(f.PointUnitFrom)
				if sdp == nil {
					continue
				}
				ddp := dm.GetEntryFromPointId(f.PointId)
				if ddp == nil {
					continue
				}
				ddp.SetUnits(sdp.GetEntry(0).Value)
				sdp.Hide()
			}
		}
	}
}

const LastEntry = -1
func (de *DataPointEntries) GetEntry(index int) DataEntry {
	for range Only.Once {
		l := len(*de) - 1
		if index > l {
			index = l
			break
		}
		if index < 0 {
			index = l + index + 1
			if index < 0 {
				index = 0
			}
		}
	}
	return (*de)[index]
}

func (de *DataPointEntries) GetUnits() string {
	var unit string
	for range Only.Once {
		for _, v := range *de {
			unit = v.Point.Unit
			break
		}
	}
	return unit
}

func (de *DataPointEntries) SetUnits(units string) {
	for range Only.Once {
		for i := range *de {
			(*de)[i].Point.Unit = units
		}
	}
}

func (dm *DataMap) GetEntry(entry string, index int) DataEntry {
	var ret DataEntry
	for range Only.Once {
		pe := dm.DataPoints[entry]
		if pe != nil {
			ret = pe.GetEntry(index)
			break
		}

		for k, v := range dm.DataPoints {
			if strings.HasSuffix(k, "." + entry) {
				ret = v.GetEntry(index)
				break
			}
		}
	}
	return ret
}

func (dm *DataMap) GetEntryFromPointId(pointId string) *DataPointEntries {
	var ret *DataPointEntries
	for range Only.Once {
		for _, v := range dm.DataPoints {
			pe := v.GetEntry(0)
			if pe.IsNotValid() {
				continue
			}

			if pe.Point.Id.String() == pointId {
				ret = &v
				break
			}
		}
	}
	return ret
}

// func (dm *DataMap) HideEntry(pointId valueTypes.PointId) {
// 	for range Only.Once {
// 		de := dm.GetEntryFromPointId(pointId)
// 		de.Hide()
// 	}
// }

func (dm *DataMap) GetFloatValue(entry string, index int) float64 {
	var ret float64
	for range Only.Once {
		pe := dm.GetEntry(entry, index)
		if pe.IsNotValid() {
			fmt.Printf("ERROR: GetFloatValue('%s', '%d')\n", entry, index)
			break
		}
		ret = pe.ValueFloat
	}
	return ret
}

func (dm *DataMap) LowerUpper(lowerEntry string, upperEntry string, index int) float64 {
	var ret float64
	for range Only.Once {
		l := dm.GetEntry(lowerEntry, index)
		if l.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lowerEntry, upperEntry, index)
			break
		}

		u := dm.GetEntry(upperEntry, index)
		if u.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lowerEntry, upperEntry, index)
			break
		}

		if l.ValueFloat > 0 {
			ret = 0 - l.ValueFloat
			break
		}
		ret = u.ValueFloat
	}
	return ret
}

func (dm *DataMap) GetPercent(entry string, max string, index int) float64 {
	var ret float64
	for range Only.Once {
		v := dm.GetEntry(entry, index)
		if v.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", entry, max, index)
			break
		}

		m := dm.GetEntry(max, index)
		if m.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", entry, max, index)
			break
		}

		ret = GetPercent(v.ValueFloat, m.ValueFloat)
	}
	return ret
}

func (dm *DataMap) GetValue(entry string, index int) float64 {
	var ret float64
	for range Only.Once {
		v := dm.GetEntry(entry, index)
		if v.IsNotValid() {
			fmt.Printf("ERROR: GetValue('%s', %d)\n", entry, index)
			break
		}

		ret = v.ValueFloat
	}
	return ret
}


func (dm *DataMap) AppendMap(add DataMap) {
	for range Only.Once {
		if dm.DataPoints == nil {
			dm.DataPoints = make(map[string]DataPointEntries)
		}

		for point, de := range add.DataPoints {
			if dd, ok := dm.DataPoints[point]; ok {
				jde, _ := json.Marshal(de)
				jdd, _ := json.Marshal(dd)
				if string(jdd) != string(jde) {
					fmt.Printf("DIFF ")
				}
				fmt.Printf("Duplicate[%s]:\n%s\n%s\n", point, jde, jdd)
				continue
			}
			dm.DataPoints[point] = de
			dm.Order = append(dm.Order, point)

			if Points.Exists(point) {
				fmt.Printf("EXISTS: %s\n", point)
			}
			Points.Add(*de[len(de)-1].Point)
		}
	}
}

func (dm *DataMap) Add(de DataEntry) {
	for range Only.Once {
		// if !strings.Contains(endpoint, ".") {
		// 	// endpoint = valueTypes.PointId(de.EndPoint + "." + string(endpoint))
		// 	endpoint = JoinWithDots(0, "", de.EndPoint, endpoint)
		// }

		// DataEntry {
		// 	Point:      Point{
		// 		Parents:   ParentDevices{},
		// 		Id:        valueTypes.PointId{},
		// 		GroupName: "",
		// 		Name:      "",
		// 		Unit:      "",
		// 		TimeSpan:  "",
		// 		ValueType: "",
		// 		Valid:     false,
		// 		States:    nil,
		// 	},
		// 	Date:       valueTypes.DateTime{},
		// 	EndPoint:   "",
		// 	FullId:     valueTypes.DataPoint{},
		// 	Parent:     ParentDevice{},
		// 	Value:      "",
		// 	ValueFloat: 0,
		// 	ValueBool:  false,
		// 	Index:      0,
		// 	Valid:      false,
		// 	Hide:       false,
		// }

		// fmt.Printf("DEBUG: dm.Add(): %s ?= %s.%s\n", de.FullId(), de.EndPoint, de.Point.Id)
		endpoint := de.FullId()
		de.Index = len(dm.Order)
		dm.DataPoints[endpoint] = append(dm.DataPoints[endpoint], de)
		dm.Order = append(dm.Order, endpoint)

		if Points.Exists(endpoint) {
			fmt.Printf("EXISTS: %s\n", endpoint)
		}
		Points.Add(*de.Point)
	}
}

func (dm *DataMap) AddEntry(endpoint string, parentId string, point Point, date valueTypes.DateTime, value string) {
	for range Only.Once {
		unit := point.Unit	// Save unit.
		vType := point.ValueType	// Save type.

		// Match to a previously defined point.
		p := GetPoint(point.Id.String())
		if p == nil {
			point = *p
		}

		// var parents ParentDevices
		// parents.Add(ParentDevice{Key: device})
		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		if point.Name == "" {
			point.Name = point.Id.PointToName()
		}
		// fid := JoinDevicePoint(parent.Key, point.Id)
		ref := valueTypes.SetUnitValueString(value, unit, vType)
		point.Unit = ref.Unit()
		point.Valid = true

		if _, ok := dm.DataPoints[point.Id.String()]; ok {
			fmt.Printf("BARF: %s\n", point.Id)
		}

		// dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
		dm.Add(DataEntry {
			EndPoint:   endpoint,
			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
			// FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Point:      &point,
			Date:       date,
			Value:      ref.String(),
			ValueFloat: ref.Value(),
		})
	}
}

func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, ref valueTypes.UnitValue) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		ref = ref.UnitValueFix()

		if name == "" {
			name = pid.PointToName()
		}

		point := GetPoint(pid.String())
		if point == nil {
			// No UV found. Create one.
			dm.Add(CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, ref))
			break
		}

		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		if point.Unit == "" {
			point.Unit = ref.Unit()
		}
		if point.Name == "" {
			point.Name = name
		}
		if point.Name == "" {
			point.Name = pid.PointToName()
		}
		if point.GroupName == "" {
			point.GroupName = groupName
		}

		dm.Add(DataEntry {
			EndPoint:   endpoint,
			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
			// FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Point:      point,
			Date:       date,
			Value:      ref.String(),
			ValueFloat: ref.Value(),
		})
	}
}

// func (dm *DataMap) AddFloat(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value float64) {
// 	for range Only.Once {
// 		// fvs := Float64ToString(value)
// 		point := GetPoint(parentId, pid)
// 		if point == nil {
// 			// No UV found. Create one.
// 			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
// 			break
// 		}
//
// 		ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
// 		if ref.Unit() != point.Unit {
// 			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
// 			point.Unit = ref.Unit()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
//
// 		dm.Add(pid, DataEntry {
// 			EndPoint:   endpoint,
// 			FullId:     JoinDevicePoint(endpoint, point.Id),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Date:       date,
// 			Point:      point,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 		})
// 	}
//
// 	uv := valueTypes.SetUnitValueFloat(value, "", "float")
// 	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 	// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 	// 	Unit:       "float",
// 	// 	Value:      fmt.Sprintf("%f", value),
// 	// 	ValueFloat: 0,
// 	// })
// 	dm.Add(pid, de)
// }
//
// func (dm *DataMap) AddString(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value string) {
// 	dm.Add(pid, CreateDataEntryString(date, endpoint, parentId, pid, name, value))
// }
//
// func (dm *DataMap) AddInt(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value int64) {
//
// 	for range Only.Once {
// 		uvs, ok := valueTypes.AnyToUnitValue(value, "", "")
// 		if !ok {
// 			fmt.Printf("ERROR: AddInt(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value %d)",
// 				endpoint, parentId, pid, name, date, value)
// 			break
// 		}
// 		for _, uv := range uvs {
// 			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 			dm.Add(pid, de)
// 		}
//
// 		// uv := valueTypes.SetUnitValueInteger(value, "", "int")
// 		// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 		// // de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 		// // 	Unit:       "int",
// 		// // 	Value:      fmt.Sprintf("%d", value),
// 		// // 	ValueFloat: float64(value),
// 		// // })
// 		// dm.Add(pid, de)
// 	}
// }

func (dm *DataMap) AddAny(endpoint string, parentId string, pid valueTypes.PointId, name string, date valueTypes.DateTime, value interface{}) {

	for range Only.Once {
		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, "", "")
		if !ok {
			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
				endpoint, parentId, pid, name, date, value)
			break
		}

		point := GetPoint(parentId + "." + pid.String())
		if point == nil {
			// No UV found. Create one.
			for _, uv := range uvs {
				de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
				if isNil {
					de.Point.ValueType += "(NIL)"
				}
				dm.Add(de)
			}
			// dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name,
			// 	valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
			break
		}

		// ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
		// if ref.Unit() != point.Unit {
		// 	fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
		// 	point.Unit = ref.Unit()
		// }

		if isNil {
			point.ValueType += "(NIL)"
		}

		for _, uv := range uvs {
			if uv.Unit() != point.Unit {
				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.Unit())
				point.Unit = uv.Unit()
			}

			var parent ParentDevice
			parent.Set(parentId)
			point.Parents.Add(parent)

			dm.Add(DataEntry{
				EndPoint: endpoint,
				// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
				Parent:   parent,

				Date:       date,
				Point:      point,
				Value:      uv.String(),
				ValueFloat: uv.Value(),
				ValueBool:  uv.ValueBool(),
			})
		}

		for _, uv := range uvs {
			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
			dm.Add(de)
		}
	}
}

func (dm *DataMap) FromRefAddAlias(entry string, parentId string, pid string, name string) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddAlias('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pe.CreateAlias(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name))
}

func (dm *DataMap) FromRefAddState(entry string, parentId string, pid string, name string) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddState('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pe.CreateState(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name))
}

func (dm *DataMap) FromRefAddFloat(entry string, parentId string, pid string, name string, value float64) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddFloat('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pe.CreateFloat(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, value))
}


func (dm *DataMap) Print() {
	for range Only.Once {
		table := datatable.New("utf8-heavy")
		table.AddHeaders(
			"Index",
			"EndPoint",

			"Id",
			"Name",
			"Unit",
			"Type",
			"Value",
			"Valid",

			"GroupName",
			"Parent Ids",
			"Parent Types",
			"Parent Codes",
		)

		for i, k := range dm.Order {
			for _, v := range dm.DataPoints[k] {
				table.AddRowItems(
					i,
					v.EndPoint,

					v.Point.Id,
					v.Point.Name,
					v.Point.Unit,
					v.Point.TimeSpan,
					v.Value,
					v.Point.Valid,
					// fmt.Sprintf("%s\n%s\n", v.FullId, v.Value),

					v.Point.GroupName,
					v.Point.Parents.PsIds(),
					v.Point.Parents.Types(),
					v.Point.Parents.Codes(),
				)
			}
		}

		ret, _ := table.Render()
		fmt.Println(ret)
	}
}


func (de *DataEntry) CreateAlias(endpoint string, parentId string, pid valueTypes.PointId, name string) DataEntry {
	if name == "" {
		name = pid.PointToName()
	}

	ret := DataEntry {
		Point:      &Point {
			Parents:   de.Point.Parents,
			Id:        pid,
			GroupName: "alias",
			Name:      name,
			Unit:      de.Point.Unit,
			TimeSpan:  de.Point.TimeSpan,
			Valid:     true,
			States:    nil,
		},
		Date:       de.Date,
		EndPoint:   endpoint,
		// FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
		// FullId:     JoinDevicePoint(parentId, pid),
		Parent:     de.Parent,		// ParentDevice{},
		Value:      de.Value,
		ValueFloat: de.ValueFloat,
		ValueBool:  de.ValueBool,
		Index:      de.Index,
	}
	ret.Point.FixUnitType()

	ret.Parent.Set(parentId)
	de.Point.Parents.Add(ret.Parent)

	// de.FullId = JoinDevicePoint(endpoint, pid)
	// de.FullId = NameDevicePoint(ret.Parent.Key, pid)
	// de.Point.Id = pid
	// de.Point.Name = name
	// de.Point.GroupName = parentId
	// de.Point.Valid = true
	// de.EndPoint = endpoint
	// de.Index = 0

	return ret
}

func (de *DataEntry) CreateFloat(endpoint string, parentId string, pid valueTypes.PointId, name string, value float64) DataEntry {
	if name == "" {
		name = pid.PointToName()
	}

	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	uv := valueTypes.SetUnitValueFloat(value, de2.Point.Unit, de2.Point.ValueType)
	de2.Value = uv.String()
	de2.ValueFloat = uv.Value()

	return de2
}

func (de *DataEntry) CreateState(endpoint string, parentId string, pid valueTypes.PointId, name string) DataEntry {
	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	if de2.ValueFloat == 0 {
		de2.Value = "false"
		de2.ValueBool = false
		de2.ValueFloat = 0
	} else {
		de2.Value = "true"
		de2.ValueBool = true
		de2.ValueFloat = 1
	}
	de2.Point.Unit = "binary"

	return de2
}


// func CreateDataEntryActive(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value float64) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "state")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      fmt.Sprintf("%v", IsActive(value)),
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }
//
// func CreateDataEntryString(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value string) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "string")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, pid),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      value,
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }

func CreateDataEntryUnitValue(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value valueTypes.UnitValue) DataEntry {
	value = value.UnitValueFix()

	point := GetPoint(parentId + "." + pid.String())
	if point == nil {
		if name == "" {
			name = pid.PointToName()
		}
		point = CreatePoint(parentId, pid, name, value.Unit())
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		// FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
		// FullId:     JoinDevicePoint(parent.Key, pid),
		Parent:     parent,

		Point:      point,
		Date:       date,
		Value:      value.String(),
		ValueFloat: value.Value(),
		Index:      0,
	}
}

func CreatePoint(parentId string, pid valueTypes.PointId, name string, unit string) *Point {
	if name == "" {
		name = pid.PointToName()
	}

	ret := &Point {
		Id:        pid,
		GroupName: parentId,
		Name:      name,
		Unit:      unit,
		TimeSpan:  "",
		Valid:     true,
	}
	ret.FixUnitType()
	return ret
}

// func IsActive(value float64) bool {
// 	if (value > 0.01) || (value < -0.01) {
// 		return true
// 	}
// 	return false
// }

func GetPercent(value float64, max float64) float64 {
	if max == 0 {
		return 0
	}
	return (value / max) * 100
}


// func JoinDevicePoint(endpoint string, pid valueTypes.PointId) valueTypes.PointId {
// 	var ret valueTypes.PointId
// 	for range Only.Once {
// 		if endpoint == "" {
// 			endpoint = "virtual"
// 		}
// 		ret = valueTypes.PointId(JoinWithDots(0, "", endpoint, pid))
// 	}
// 	return ret
// }

func JoinWithDots(intSize int, dateFormat string, args ...interface{}) string {
	var ret string
	for range Only.Once {
		var a []string
		for _, e := range args {
			v := valueTypes.TypeToString(intSize, dateFormat, e)
			if v == "" {
				continue
			}
			a = append(a, v)
		}
		ret = strings.Join(a, ".")
	}
	return ret
}

// func JoinStringsWithDots(args ...string) string {
// 	return strings.Join(args, ".")
// }
