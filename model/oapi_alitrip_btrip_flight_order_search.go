package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripFlightOrderSearchRequest() *OapiAlitripBtripFlightOrderSearchRequest {
	return &OapiAlitripBtripFlightOrderSearchRequest{}
}

type OapiAlitripBtripFlightOrderSearchRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripFlightOrderSearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripFlightOrderSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.flight.order.search"
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripFlightOrderSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripFlightOrderSearchResponse struct {
	taobao.TaobaoResponse

	FlightOrderList []OpenFlightOrderRs `json:"flight_order_list,omitempty"`
	Success         bool                `json:"success,omitempty"`
}
type OpenInvoiceDo struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
type OpenCostCenterDo struct {
	Corpid string `json:"corpid,omitempty"`
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}
type OpenPriceInfo struct {
	Category      string    `json:"category,omitempty"`
	GmtCreate     time.Time `json:"gmt_create,omitempty"`
	PassengerName string    `json:"passenger_name,omitempty"`
	PayType       int64     `json:"pay_type,omitempty"`
	Price         string    `json:"price,omitempty"`
	Type          int64     `json:"type,omitempty"`
}
type OpenFlightInsureInfo struct {
	InsureNo string `json:"insure_no,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   int64  `json:"status,omitempty"`
}
type OpenUserAffiliateDo struct {
	UserName string `json:"user_name,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type OpenFlightOrderRs struct {
	ApplyId              string                 `json:"apply_id,omitempty"`
	ArrAirport           string                 `json:"arr_airport,omitempty"`
	ArrCity              string                 `json:"arr_city,omitempty"`
	BtripTitle           string                 `json:"btrip_title,omitempty"`
	CabinClass           string                 `json:"cabin_class,omitempty"`
	ContactName          string                 `json:"contact_name,omitempty"`
	CorpName             string                 `json:"corp_name,omitempty"`
	Corpid               string                 `json:"corpid,omitempty"`
	CostCenter           OpenCostCenterDo       `json:"cost_center,omitempty"`
	DepAirport           string                 `json:"dep_airport,omitempty"`
	DepCity              string                 `json:"dep_city,omitempty"`
	DepDate              time.Time              `json:"dep_date,omitempty"`
	DeptName             string                 `json:"dept_name,omitempty"`
	Deptid               string                 `json:"deptid,omitempty"`
	Discount             string                 `json:"discount,omitempty"`
	FlightNo             string                 `json:"flight_no,omitempty"`
	GmtCreate            time.Time              `json:"gmt_create,omitempty"`
	GmtModified          time.Time              `json:"gmt_modified,omitempty"`
	Id                   int64                  `json:"id,omitempty"`
	InsureInfoList       []OpenFlightInsureInfo `json:"insureInfo_list,omitempty"`
	Invoice              OpenInvoiceDo          `json:"invoice,omitempty"`
	PassengerCount       int64                  `json:"passenger_count,omitempty"`
	PassengerName        string                 `json:"passenger_name,omitempty"`
	PriceInfoList        []OpenPriceInfo        `json:"price_info_list,omitempty"`
	RetDate              time.Time              `json:"ret_date,omitempty"`
	Status               int64                  `json:"status,omitempty"`
	ThirdpartApplyId     string                 `json:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId string                 `json:"thirdpart_itinerary_id,omitempty"`
	TripType             int64                  `json:"trip_type,omitempty"`
	UserAffiliateList    []OpenUserAffiliateDo  `json:"user_affiliate_list,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`
	Userid               string                 `json:"userid,omitempty"`
}
