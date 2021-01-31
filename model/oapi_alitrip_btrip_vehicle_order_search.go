package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripVehicleOrderSearchRequest() *OapiAlitripBtripVehicleOrderSearchRequest {
	return &OapiAlitripBtripVehicleOrderSearchRequest{}
}

type OapiAlitripBtripVehicleOrderSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripVehicleOrderSearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.vehicle.order.search"
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripVehicleOrderSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripVehicleOrderSearchResponse struct {
	taobao.TaobaoResponse
	Errcode          int64                `json:"errcode,omitempty"`
	Errmsg           string               `json:"errmsg,omitempty"`
	Success          bool                 `json:"success,omitempty"`
	VehicleOrderList []OpenVehicleOrderRs `json:"vehicle_order_list,omitempty"`
}
type OpenVehicleOrderRs struct {
	ApplyId              int64                 `json:"apply_id,omitempty"`
	ApplyShowId          string                `json:"apply_show_id,omitempty"`
	BtripTitle           string                `json:"btrip_title,omitempty"`
	BusinessCategory     string                `json:"business_category,omitempty"`
	CancelTime           time.Time             `json:"cancel_time,omitempty"`
	CarInfo              string                `json:"car_info,omitempty"`
	CarLevel             string                `json:"car_level,omitempty"`
	CorpName             string                `json:"corp_name,omitempty"`
	Corpid               string                `json:"corpid,omitempty"`
	CostCenterId         int64                 `json:"cost_center_id,omitempty"`
	CostCenterName       string                `json:"cost_center_name,omitempty"`
	CostCenterNumber     string                `json:"cost_center_number,omitempty"`
	DeptName             string                `json:"dept_name,omitempty"`
	Deptid               string                `json:"deptid,omitempty"`
	DriverConfirmTime    time.Time             `json:"driver_confirm_time,omitempty"`
	EstimatePrice        string                `json:"estimate_price,omitempty"`
	FromAddress          string                `json:"from_address,omitempty"`
	FromCityName         string                `json:"from_city_name,omitempty"`
	GmtCreate            time.Time             `json:"gmt_create,omitempty"`
	GmtModified          time.Time             `json:"gmt_modified,omitempty"`
	Id                   int64                 `json:"id,omitempty"`
	InvoiceId            int64                 `json:"invoice_id,omitempty"`
	InvoiceTitle         string                `json:"invoice_title,omitempty"`
	IsSpecial            bool                  `json:"is_special,omitempty"`
	Memo                 string                `json:"memo,omitempty"`
	OrderStatus          int64                 `json:"order_status,omitempty"`
	PassengerName        string                `json:"passenger_name,omitempty"`
	PayTime              time.Time             `json:"pay_time,omitempty"`
	PriceInfoList        []OpenPriceInfo       `json:"price_info_list,omitempty"`
	ProjectCode          string                `json:"project_code,omitempty"`
	ProjectTitle         string                `json:"project_title,omitempty"`
	Provider             int64                 `json:"provider,omitempty"`
	PublishTime          time.Time             `json:"publish_time,omitempty"`
	RealFromAddress      string                `json:"real_from_address,omitempty"`
	RealFromCityName     string                `json:"real_from_city_name,omitempty"`
	RealToAddress        string                `json:"real_to_address,omitempty"`
	RealToCityName       string                `json:"real_to_city_name,omitempty"`
	ServiceType          int64                 `json:"service_type,omitempty"`
	SpecialTypes         []string              `json:"special_types,omitempty"`
	TakenTime            time.Time             `json:"taken_time,omitempty"`
	ThirdpartApplyId     string                `json:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId string                `json:"thirdpart_itinerary_id,omitempty"`
	ToAddress            string                `json:"to_address,omitempty"`
	ToCityName           string                `json:"to_city_name,omitempty"`
	TravelDistance       string                `json:"travel_distance,omitempty"`
	UserAffiliateList    []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty"`
	UserConfirm          int64                 `json:"user_confirm,omitempty"`
	UserName             string                `json:"user_name,omitempty"`
	Userid               string                `json:"userid,omitempty"`
}
