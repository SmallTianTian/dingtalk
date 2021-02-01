package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripTrainOrderSearchRequest() *OapiAlitripBtripTrainOrderSearchRequest {
	return &OapiAlitripBtripTrainOrderSearchRequest{}
}

type OapiAlitripBtripTrainOrderSearchRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripTrainOrderSearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripTrainOrderSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.train.order.search"
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripTrainOrderSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripTrainOrderSearchResponse struct {
	taobao.TaobaoResponse

	Success        bool               `json:"success,omitempty"`
	TrainOrderList []OpenTrainOrderRs `json:"train_order_list,omitempty"`
}
type OpenTrainOrderRs struct {
	ApplyId              int64                 `json:"apply_id,omitempty"`
	ArrCity              string                `json:"arr_city,omitempty"`
	ArrStation           string                `json:"arr_station,omitempty"`
	ArrTime              time.Time             `json:"arr_time,omitempty"`
	BtripTitle           string                `json:"btrip_title,omitempty"`
	ContactName          string                `json:"contact_name,omitempty"`
	CorpName             string                `json:"corp_name,omitempty"`
	Corpid               string                `json:"corpid,omitempty"`
	CostCenter           OpenCostCenterDo      `json:"cost_center,omitempty"`
	DepCity              string                `json:"dep_city,omitempty"`
	DepStation           string                `json:"dep_station,omitempty"`
	DepTime              time.Time             `json:"dep_time,omitempty"`
	DeptName             string                `json:"dept_name,omitempty"`
	Deptid               string                `json:"deptid,omitempty"`
	GmtCreate            time.Time             `json:"gmt_create,omitempty"`
	GmtModified          time.Time             `json:"gmt_modified,omitempty"`
	Id                   int64                 `json:"id,omitempty"`
	Invoice              OpenInvoiceDo         `json:"invoice,omitempty"`
	PriceInfoList        []OpenPriceInfo       `json:"price_info_list,omitempty"`
	RiderName            string                `json:"rider_name,omitempty"`
	RunTime              string                `json:"run_time,omitempty"`
	SeatType             string                `json:"seat_type,omitempty"`
	Status               int64                 `json:"status,omitempty"`
	ThirdpartApplyId     string                `json:"thirdpart_apply_id,omitempty"`
	ThirdpartItineraryId string                `json:"thirdpart_itinerary_id,omitempty"`
	TicketCount          int64                 `json:"ticket_count,omitempty"`
	TicketNo12306        string                `json:"ticket_no_12306,omitempty"`
	TrainNumber          string                `json:"train_number,omitempty"`
	TrainType            string                `json:"train_type,omitempty"`
	UserAffiliateList    []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty"`
	UserName             string                `json:"user_name,omitempty"`
	Userid               string                `json:"userid,omitempty"`
}
