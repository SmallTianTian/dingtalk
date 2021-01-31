package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripHotelOrderSearchRequest() *OapiAlitripBtripHotelOrderSearchRequest {
	return &OapiAlitripBtripHotelOrderSearchRequest{}
}

type OapiAlitripBtripHotelOrderSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripHotelOrderSearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripHotelOrderSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.hotel.order.search"
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripHotelOrderSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripHotelOrderSearchResponse struct {
	taobao.TaobaoResponse
	Module  []OpenHotelOrderRs `json:"module,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type OpenHotelOrderRs struct {
	ApplyId              int64                 `json:"apply_id,omitempty"`
	BtripTitle           string                `json:"btrip_title,omitempty"`
	CheckIn              time.Time             `json:"check_in,omitempty"`
	CheckOut             time.Time             `json:"check_out,omitempty"`
	City                 string                `json:"city,omitempty"`
	ContactName          string                `json:"contact_name,omitempty"`
	CorpName             string                `json:"corp_name,omitempty"`
	Corpid               string                `json:"corpid,omitempty"`
	CostCenter           OpenCostCenterDo      `json:"cost_center,omitempty"`
	DeptName             string                `json:"dept_name,omitempty"`
	Deptid               string                `json:"deptid,omitempty"`
	GmtCreate            time.Time             `json:"gmt_create,omitempty"`
	GmtModified          time.Time             `json:"gmt_modified,omitempty"`
	Guest                string                `json:"guest,omitempty"`
	HotelName            string                `json:"hotel_name,omitempty"`
	Id                   int64                 `json:"id,omitempty"`
	Invoice              OpenInvoiceDo         `json:"invoice,omitempty"`
	Night                int64                 `json:"night,omitempty"`
	OrderStatus          int64                 `json:"order_status,omitempty"`
	OrderStatusDesc      string                `json:"order_status_desc,omitempty"`
	OrderType            int64                 `json:"order_type,omitempty"`
	OrderTypeDesc        string                `json:"order_type_desc,omitempty"`
	PriceInfoList        []OpenPriceInfo       `json:"price_info_list,omitempty"`
	RoomNum              int64                 `json:"room_num,omitempty"`
	RoomType             string                `json:"room_type,omitempty"`
	ThirdpartApplyId     string                `json:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId string                `json:"thirdpart_itinerary_id,omitempty"`
	UserAffiliateList    []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty"`
	UserName             string                `json:"user_name,omitempty"`
	Userid               string                `json:"userid,omitempty"`
}
