package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripPriceQueryRequest() *OapiAlitripBtripPriceQueryRequest {
	return &OapiAlitripBtripPriceQueryRequest{}
}

type OapiAlitripBtripPriceQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripPriceQueryResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripPriceQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripPriceQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripPriceQueryRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiAlitripBtripPriceQueryRequest) GetReq() string {
	return this.Req
}
func (this *OapiAlitripBtripPriceQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.price.query"
}
func (this *OapiAlitripBtripPriceQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripPriceQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripPriceQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripPriceQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripPriceQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiAlitripBtripPriceQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripPriceQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripPriceQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiIntervalPriceRq struct {
	Category    string    `json:"category,omitempty"`
	Corpid      string    `json:"corpid,omitempty"`
	EndTime     time.Time `json:"end_time,omitempty"`
	FromWhere   string    `json:"from_where,omitempty"`
	ItineraryId string    `json:"itinerary_id,omitempty"`
	QueryKey    string    `json:"query_key,omitempty"`
	StartTime   time.Time `json:"start_time,omitempty"`
	ToWhere     string    `json:"to_where,omitempty"`
	Userid      string    `json:"userid,omitempty"`
}
type OapiAlitripBtripPriceQueryResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
type HotelFeeDetail struct {
	City      string `json:"city,omitempty"`
	Criterion int64  `json:"criterion,omitempty"`
}
type MostExpensive struct {
	ArrTime   string `json:"arr_time,omitempty"`
	DepTime   string `json:"dep_time,omitempty"`
	Fee       int64  `json:"fee,omitempty"`
	SeatGrade string `json:"seat_grade,omitempty"`
	VehicleNo string `json:"vehicle_no,omitempty"`
}
type Cheapest struct {
	ArrTime   string `json:"arr_time,omitempty"`
	DepTime   string `json:"dep_time,omitempty"`
	Fee       int64  `json:"fee,omitempty"`
	SeatGrade string `json:"seat_grade,omitempty"`
	VehicleNo string `json:"vehicle_no,omitempty"`
}
type BtripRoutes struct {
	Cheapest      Cheapest      `json:"cheapest,omitempty"`
	DepDate       time.Time     `json:"dep_date,omitempty"`
	DestCity      string        `json:"dest_city,omitempty"`
	ErrMsg        string        `json:"err_msg,omitempty"`
	MostExpensive MostExpensive `json:"most_expensive,omitempty"`
	OrgCity       string        `json:"org_city,omitempty"`
	Success       bool          `json:"success,omitempty"`
}
type TrafficFee struct {
	BtripRoutes []BtripRoutes `json:"btrip_routes,omitempty"`
	ErrMsg      string        `json:"err_msg,omitempty"`
	Success     bool          `json:"success,omitempty"`
}
type Module struct {
	HotelFeeDetail []HotelFeeDetail `json:"hotel_fee_detail,omitempty"`
	QueryKey       string           `json:"query_key,omitempty"`
	TrafficFee     TrafficFee       `json:"traffic_fee,omitempty"`
}
