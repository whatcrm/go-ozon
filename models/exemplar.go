package models

type ExemplarStatusRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ExemplarStatusResponse struct {
	PostingNumber string                       `json:"posting_number"`
	Products      []ExemplarStatusProductEntry `json:"products"`
	Status        string                       `json:"status"`
}

type ExemplarStatusProductEntry struct {
	Exemplars []ExemplarStatusItem `json:"exemplars"`
	ProductID int64                `json:"product_id"`
}

type ExemplarStatusItem struct {
	ExemplarID        int64                `json:"exemplar_id"`
	GTD               string               `json:"gtd"`
	GTDCheckStatus    string               `json:"gtd_check_status"`
	GTDErrorCodes     []string             `json:"gtd_error_codes"`
	IsGTDAbsent       bool                 `json:"is_gtd_absent"`
	IsRNPTAbsent      bool                 `json:"is_rnpt_absent"`
	Marks             []ExemplarStatusMark `json:"marks"`
	RNPT              string               `json:"rnpt"`
	RNPTCheckStatus   string               `json:"rnpt_check_status"`
	RNPTErrorCodes    []string             `json:"rnpt_error_codes"`
	Weight            float64              `json:"weight"`
	WeightCheckStatus string               `json:"weight_check_status"`
	WeightErrorCodes  []string             `json:"weight_error_codes"`
}

type ExemplarStatusMark struct {
	CheckStatus string   `json:"check_status"`
	ErrorCodes  []string `json:"error_codes"`
	Mark        string   `json:"mark"`
	MarkType    string   `json:"mark_type"`
}

type ExemplarValidateRequest struct {
	PostingNumber string                        `json:"posting_number"`
	Products      []ExemplarValidateProductItem `json:"products"`
}

type ExemplarValidateProductItem struct {
	Exemplars []ExemplarValidateItem `json:"exemplars"`
	ProductID int64                  `json:"product_id"`
}

type ExemplarValidateItem struct {
	GTD    string                 `json:"gtd,omitempty"`
	Marks  []ExemplarValidateMark `json:"marks,omitempty"`
	RNPT   string                 `json:"rnpt,omitempty"`
	Weight float64                `json:"weight,omitempty"`
}

type ExemplarValidateMark struct {
	Mark     string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type ExemplarValidateResponse struct {
	Products []ExemplarValidateResponseProduct `json:"products"`
}

type ExemplarValidateResponseProduct struct {
	Error     string                         `json:"error"`
	Exemplars []ExemplarValidateResponseItem `json:"exemplars"`
	ProductID int64                          `json:"product_id"`
	Valid     bool                           `json:"valid"`
}

type ExemplarValidateResponseItem struct {
	Errors []string                       `json:"errors"`
	GTD    string                         `json:"gtd"`
	Marks  []ExemplarValidateResponseMark `json:"marks"`
	RNPT   string                         `json:"rnpt"`
	Valid  bool                           `json:"valid"`
	Weight float64                        `json:"weight"`
}

type ExemplarValidateResponseMark struct {
	Errors   []string `json:"errors"`
	Mark     string   `json:"mark"`
	MarkType string   `json:"mark_type"`
	Valid    bool     `json:"valid"`
}

type ExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ExemplarUpdateResponse struct {
	Code    int32             `json:"code"`
	Details []GRPCErrorDetail `json:"details"`
	Message string            `json:"message"`
}

type GRPCErrorDetail struct {
	TypeURL string `json:"typeUrl"`
	Value   string `json:"value"`
}

type ExemplarSetRequest struct {
	MultiBoxQty   int32                `json:"multi_box_qty,omitempty"`
	PostingNumber string               `json:"posting_number"`
	Products      []ExemplarSetProduct `json:"products"`
}

type ExemplarSetProduct struct {
	Exemplars []ExemplarSetItem `json:"exemplars"`
	ProductID int64             `json:"product_id"`
}

type ExemplarSetItem struct {
	ExemplarID   int64                  `json:"exemplar_id,omitempty"`
	GTD          string                 `json:"gtd,omitempty"`
	IsGTDAbsent  bool                   `json:"is_gtd_absent,omitempty"`
	IsRNPTAbsent bool                   `json:"is_rnpt_absent,omitempty"`
	Marks        []ExemplarValidateMark `json:"marks,omitempty"`
	RNPT         string                 `json:"rnpt,omitempty"`
	Weight       float64                `json:"weight,omitempty"`
}

type ExemplarSetResponse struct {
	Code    int32             `json:"code"`
	Details []GRPCErrorDetail `json:"details"`
	Message string            `json:"message"`
}

type ExemplarCreateOrGetRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ExemplarCreateOrGetResponse struct {
	MultiBoxQty   int32                        `json:"multi_box_qty"`
	PostingNumber string                       `json:"posting_number"`
	Products      []ExemplarCreateOrGetProduct `json:"products"`
}

type ExemplarCreateOrGetProduct struct {
	Exemplars               []ExemplarCreateOrGetItem `json:"exemplars"`
	HasIMEI                 bool                      `json:"has_imei"`
	IsGTDNeeded             bool                      `json:"is_gtd_needed"`
	IsJwUinNeeded           bool                      `json:"is_jw_uin_needed"`
	IsMandatoryMarkNeeded   bool                      `json:"is_mandatory_mark_needed"`
	IsMandatoryMarkPossible bool                      `json:"is_mandatory_mark_possible"`
	IsRNPTNeeded            bool                      `json:"is_rnpt_needed"`
	ProductID               int64                     `json:"product_id"`
	Quantity                int32                     `json:"quantity"`
	IsWeightNeeded          bool                      `json:"is_weight_needed"`
	WeightMax               float64                   `json:"weight_max"`
	WeightMin               float64                   `json:"weight_min"`
}

type ExemplarCreateOrGetItem struct {
	ExemplarID   int64                  `json:"exemplar_id"`
	GTD          string                 `json:"gtd"`
	IsGTDAbsent  bool                   `json:"is_gtd_absent"`
	IsRNPTAbsent bool                   `json:"is_rnpt_absent"`
	Marks        []ExemplarValidateMark `json:"marks"`
	RNPT         string                 `json:"rnpt"`
	Weight       float64                `json:"weight"`
}
