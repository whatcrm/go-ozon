package models

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WarehouseTimetableHoliday struct {
	Day  string `json:"day"`
	From string `json:"from"`
	To   string `json:"to"`
}

type WarehouseTimetableWorkingDay struct {
	Day  string `json:"day"`
	From string `json:"from"`
	To   string `json:"to"`
}

type WarehouseTimetable struct {
	Holidays    []WarehouseTimetableHoliday    `json:"holidays"`
	WorkingDays []WarehouseTimetableWorkingDay `json:"working_days"`
}

type ErfbsReturnSettings struct {
	ContactDays          int32  `json:"contact_days"`
	PostOfficeZipcode    string `json:"post_office_zipcode"`
	ReturnMethod         string `json:"return_method"`
	TransportCompanyName string `json:"transport_company_name"`
}

type ErfbsAggregatorDeliveryCosts struct {
	MaxAmount float64 `json:"max_amount"`
	MinAmount float64 `json:"min_amount"`
	Percent   float64 `json:"percent"`
}

type ErfbsAggregatorDeliveryMethod struct {
	CourierComment string                       `json:"courier_comment"`
	CourierPhones  []string                     `json:"courier_phones"`
	CutIn          int64                        `json:"cut_in"`
	DeliverToPVZ   bool                         `json:"deliver_to_pvz"`
	DeliveryCosts  ErfbsAggregatorDeliveryCosts `json:"delivery_costs"`
	Name           string                       `json:"name"`
	ReturnSettings ErfbsReturnSettings          `json:"return_settings"`
}

type ErfbsAggregatorCreateRequest struct {
	AddressCoordinates Coordinates                   `json:"address_coordinates"`
	DeliveryMethod     ErfbsAggregatorDeliveryMethod `json:"delivery_method"`
	MinOrderValue      int64                         `json:"min_order_value,omitempty"`
	Name               string                        `json:"name"`
	Phone              string                        `json:"phone"`
	TimetableWarehouse WarehouseTimetable            `json:"timetable_warehouse"`
}

type ErfbsWarehouseUpdateRequest struct {
	MinOrderValue      int64               `json:"min_order_value,omitempty"`
	Name               string              `json:"name,omitempty"`
	Phone              string              `json:"phone,omitempty"`
	TimetableWarehouse *WarehouseTimetable `json:"timetable_warehouse,omitempty"`
	WarehouseID        int64               `json:"warehouse_id"`
}

type ErfbsAggregatorDeliveryMethodUpdateRequest struct {
	CourierComment   string                        `json:"courier_comment,omitempty"`
	CourierPhones    []string                      `json:"courier_phones,omitempty"`
	CutIn            int64                         `json:"cut_in,omitempty"`
	DeliverToPVZ     bool                          `json:"deliver_to_pvz,omitempty"`
	DeliveryCosts    *ErfbsAggregatorDeliveryCosts `json:"delivery_costs,omitempty"`
	DeliveryMethodID int64                         `json:"delivery_method_id"`
	Name             string                        `json:"name,omitempty"`
	ReturnSettings   *ErfbsReturnSettings          `json:"return_settings,omitempty"`
	WarehouseID      int64                         `json:"warehouse_id"`
}

type ErfbsNonIntegratedDeliveryPolygon struct {
	ID   int64 `json:"id"`
	Time int64 `json:"time"`
}

type ErfbsNonIntegratedDeliveryMethod struct {
	CourierCutoff    int64                               `json:"courier_cutoff"`
	CutIn            int64                               `json:"cut_in"`
	DeliveryPolygons []ErfbsNonIntegratedDeliveryPolygon `json:"delivery_polygons"`
	Name             string                              `json:"name"`
	ReturnSettings   ErfbsReturnSettings                 `json:"return_settings"`
}

type ErfbsNonIntegratedCreateRequest struct {
	AddressCoordinates Coordinates                      `json:"address_coordinates"`
	DeliveryMethod     ErfbsNonIntegratedDeliveryMethod `json:"delivery_method"`
	MinOrderValue      int64                            `json:"min_order_value,omitempty"`
	Name               string                           `json:"name"`
	Phone              string                           `json:"phone"`
	TimetableWarehouse WarehouseTimetable               `json:"timetable_warehouse"`
}

type ErfbsNonIntegratedDeliveryMethodUpdateRequest struct {
	CourierCutoff    int64               `json:"courier_cutoff"`
	CutIn            int64               `json:"cut_in"`
	DeliveryMethodID int64               `json:"delivery_method_id"`
	Name             string              `json:"name"`
	ReturnSettings   ErfbsReturnSettings `json:"return_settings"`
	WarehouseID      int64               `json:"warehouse_id"`
}

type WarehouseOperationResponse struct {
	OperationID string `json:"operation_id"`
}
