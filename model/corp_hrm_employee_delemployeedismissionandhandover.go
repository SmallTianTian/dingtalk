package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHrmEmployeeDelemployeedismissionandhandoverRequest() *CorpHrmEmployeeDelemployeedismissionandhandoverRequest {
	return &CorpHrmEmployeeDelemployeedismissionandhandoverRequest{}
}

type CorpHrmEmployeeDelemployeedismissionandhandoverRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                       CorpHrmEmployeeDelemployeedismissionandhandoverResponse
	DismissionInfoWithHandOver string
	OpUserid                   string
	TopHttpMethod              string
	TopResponseType            string
}

func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) SetDismissionInfoWithHandOver(dismissionInfoWithHandOver2 string) {
	this.DismissionInfoWithHandOver = dismissionInfoWithHandOver2
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetDismissionInfoWithHandOver() string {
	return this.DismissionInfoWithHandOver
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.delemployeedismissionandhandover"
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dismission_info_with_hand_over"] = this.DismissionInfoWithHandOver
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeDelemployeedismissionandhandoverRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpDismissionInfoWithHandOverVo struct {
	DismissionMemo   string    `json:"dismission_memo,omitempty"`
	DismissionReason string    `json:"dismission_reason,omitempty"`
	DismissionUserid string    `json:"dismission_userid,omitempty"`
	HandOverUserid   string    `json:"hand_over_userid,omitempty"`
	LastWorkDate     time.Time `json:"last_work_date,omitempty"`
}
type CorpHrmEmployeeDelemployeedismissionandhandoverResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
