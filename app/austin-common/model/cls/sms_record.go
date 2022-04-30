package cls

var ClsSmsRecord = struct {
	ID                string
	MessageTemplateID string
	Phone             string
	SupplierID        string
	SupplierName      string
	MsgContent        string
	SeriesID          string
	ChargingNum       string
	ReportContent     string
	Status            string
	SendDate          string
	Created           string
	Updated           string
}{
	ID:                "id",
	MessageTemplateID: "message_template_id",
	Phone:             "phone",
	SupplierID:        "supplier_id",
	SupplierName:      "supplier_name",
	MsgContent:        "msg_content",
	SeriesID:          "series_id",
	ChargingNum:       "charging_num",
	ReportContent:     "report_content",
	Status:            "status",
	SendDate:          "send_date",
	Created:           "created",
	Updated:           "updated",
}
