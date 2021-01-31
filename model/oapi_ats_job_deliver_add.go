package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsJobDeliverAddRequest() *OapiAtsJobDeliverAddRequest {
	return &OapiAtsJobDeliverAddRequest{}
}

type OapiAtsJobDeliverAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsJobDeliverAddResponse
	BizCode         string
	DeliverChannel  string
	DeliverMsg      string
	DeliverOuterId  string
	DeliverStatus   string
	JobId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsJobDeliverAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsJobDeliverAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsJobDeliverAddRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsJobDeliverAddRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsJobDeliverAddRequest) SetDeliverChannel(deliverChannel2 string) {
	this.DeliverChannel = deliverChannel2
}
func (this *OapiAtsJobDeliverAddRequest) GetDeliverChannel() string {
	return this.DeliverChannel
}
func (this *OapiAtsJobDeliverAddRequest) SetDeliverMsg(deliverMsg2 string) {
	this.DeliverMsg = deliverMsg2
}
func (this *OapiAtsJobDeliverAddRequest) GetDeliverMsg() string {
	return this.DeliverMsg
}
func (this *OapiAtsJobDeliverAddRequest) SetDeliverOuterId(deliverOuterId2 string) {
	this.DeliverOuterId = deliverOuterId2
}
func (this *OapiAtsJobDeliverAddRequest) GetDeliverOuterId() string {
	return this.DeliverOuterId
}
func (this *OapiAtsJobDeliverAddRequest) SetDeliverStatus(deliverStatus2 string) {
	this.DeliverStatus = deliverStatus2
}
func (this *OapiAtsJobDeliverAddRequest) GetDeliverStatus() string {
	return this.DeliverStatus
}
func (this *OapiAtsJobDeliverAddRequest) SetJobId(jobId2 string) {
	this.JobId = jobId2
}
func (this *OapiAtsJobDeliverAddRequest) GetJobId() string {
	return this.JobId
}
func (this *OapiAtsJobDeliverAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.job.deliver.add"
}
func (this *OapiAtsJobDeliverAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsJobDeliverAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsJobDeliverAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsJobDeliverAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsJobDeliverAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["deliver_channel"] = this.DeliverChannel
	txtParams["deliver_msg"] = this.DeliverMsg
	txtParams["deliver_outer_id"] = this.DeliverOuterId
	txtParams["deliver_status"] = this.DeliverStatus
	txtParams["job_id"] = this.JobId
	return txtParams
}
func (this *OapiAtsJobDeliverAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsJobDeliverAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsJobDeliverAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsJobDeliverAddResponse struct {
	taobao.TaobaoResponse
	DeliverId string `json:"deliver_id,omitempty"`
	Errcode   int64  `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
}
