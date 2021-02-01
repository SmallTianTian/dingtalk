package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDingReceiverstatusListRequest() *CorpDingReceiverstatusListRequest {
	return &CorpDingReceiverstatusListRequest{}
}

type CorpDingReceiverstatusListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDingReceiverstatusListResponse
	ConfirmedStatus int64
	DingId          string
	PageNo          int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDingReceiverstatusListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDingReceiverstatusListRequest) SetConfirmedStatus(confirmedStatus2 int64) {
	this.ConfirmedStatus = confirmedStatus2
}
func (this *CorpDingReceiverstatusListRequest) GetConfirmedStatus() int64 {
	return this.ConfirmedStatus
}
func (this *CorpDingReceiverstatusListRequest) SetDingId(dingId2 string) {
	this.DingId = dingId2
}
func (this *CorpDingReceiverstatusListRequest) GetDingId() string {
	return this.DingId
}
func (this *CorpDingReceiverstatusListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *CorpDingReceiverstatusListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *CorpDingReceiverstatusListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *CorpDingReceiverstatusListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *CorpDingReceiverstatusListRequest) GetApiMethodName() string {
	return "dingtalk.corp.ding.receiverstatus.list"
}
func (this *CorpDingReceiverstatusListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDingReceiverstatusListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDingReceiverstatusListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDingReceiverstatusListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDingReceiverstatusListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDingReceiverstatusListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["confirmed_status"] = this.ConfirmedStatus
	txtParams["ding_id"] = this.DingId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *CorpDingReceiverstatusListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DingId, "dingId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDingReceiverstatusListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDingReceiverstatusListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDingReceiverstatusListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type DingReceiverStatusResult struct {
	ConfirmedStatus int64  `json:"confirmed_status,omitempty"`
	ConfirmedTime   int64  `json:"confirmed_time,omitempty"`
	DingId          string `json:"ding_id,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
