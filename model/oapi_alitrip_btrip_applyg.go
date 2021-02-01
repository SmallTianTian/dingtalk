package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripApplyGetRequest() *OapiAlitripBtripApplyGetRequest {
	return &OapiAlitripBtripApplyGetRequest{}
}

type OapiAlitripBtripApplyGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripApplyGetResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripApplyGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripApplyGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripApplyGetRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripApplyGetRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripApplyGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.apply.get"
}
func (this *OapiAlitripBtripApplyGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripApplyGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripApplyGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripApplyGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripApplyGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripApplyGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripApplyGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripApplyGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenSearchRq struct {
	ApplyId          int64  `json:"apply_id,omitempty"`
	ApplyShowId      string `json:"apply_show_id,omitempty"`
	Corpid           string `json:"corpid,omitempty"`
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty"`
}
type OapiAlitripBtripApplyGetResponse struct {
	taobao.TaobaoResponse
	Module  OpenApplyRs `json:"module,omitempty"`
	Success bool        `json:"success,omitempty"`
}
type OpenItineraryInfo struct {
	ArrCity        string    `json:"arr_city,omitempty"`
	ArrDate        time.Time `json:"arr_date,omitempty"`
	CostCenterName string    `json:"cost_center_name,omitempty"`
	DepCity        string    `json:"dep_city,omitempty"`
	DepDate        time.Time `json:"dep_date,omitempty"`
	InvoiceName    string    `json:"invoice_name,omitempty"`
	ItineraryId    string    `json:"itinerary_id,omitempty"`
	ProjectCode    string    `json:"project_code,omitempty"`
	ProjectTitle   string    `json:"project_title,omitempty"`
	TrafficType    int64     `json:"traffic_type,omitempty"`
	TripWay        int64     `json:"trip_way,omitempty"`
}
type OpenUserInfo struct {
	UserName string `json:"user_name,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type OpenApproverInfo struct {
	Note        string    `json:"note,omitempty"`
	OperateTime time.Time `json:"operate_time,omitempty"`
	Order       int64     `json:"order,omitempty"`
	Status      int64     `json:"status,omitempty"`
	StatusDesc  string    `json:"status_desc,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Userid      string    `json:"userid,omitempty"`
}
type OpenApplyRs struct {
	ApplyShowId   string              `json:"apply_show_id,omitempty"`
	ApproverList  []OpenApproverInfo  `json:"approver_list,omitempty"`
	CorpName      string              `json:"corp_name,omitempty"`
	Corpid        string              `json:"corpid,omitempty"`
	DeptName      string              `json:"dept_name,omitempty"`
	Deptid        string              `json:"deptid,omitempty"`
	GmtCreate     time.Time           `json:"gmt_create,omitempty"`
	GmtModified   time.Time           `json:"gmt_modified,omitempty"`
	Id            int64               `json:"id,omitempty"`
	ItineraryList []OpenItineraryInfo `json:"itinerary_list,omitempty"`
	Status        int64               `json:"status,omitempty"`
	StatusDesc    string              `json:"status_desc,omitempty"`
	ThirdpartId   string              `json:"thirdpart_id,omitempty"`
	TravelerList  []OpenUserInfo      `json:"traveler_list,omitempty"`
	TripCause     string              `json:"trip_cause,omitempty"`
	TripDay       int64               `json:"trip_day,omitempty"`
	TripTitle     string              `json:"trip_title,omitempty"`
	UserName      string              `json:"user_name,omitempty"`
	Userid        string              `json:"userid,omitempty"`
}
