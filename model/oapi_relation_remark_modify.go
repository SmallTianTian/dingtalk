package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRelationRemarkModifyRequest() *OapiRelationRemarkModifyRequest {
	return &OapiRelationRemarkModifyRequest{}
}

type OapiRelationRemarkModifyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRelationRemarkModifyResponse
	Markees         string
	Markers         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRelationRemarkModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRelationRemarkModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRelationRemarkModifyRequest) SetMarkees(markees2 string) {
	this.Markees = markees2
}
func (this *OapiRelationRemarkModifyRequest) GetMarkees() string {
	return this.Markees
}
func (this *OapiRelationRemarkModifyRequest) SetMarkers(markers2 string) {
	this.Markers = markers2
}
func (this *OapiRelationRemarkModifyRequest) GetMarkers() string {
	return this.Markers
}
func (this *OapiRelationRemarkModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.relation.remark.modify"
}
func (this *OapiRelationRemarkModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRelationRemarkModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRelationRemarkModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRelationRemarkModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRelationRemarkModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["markees"] = this.Markees
	txtParams["markers"] = this.Markers
	return txtParams
}
func (this *OapiRelationRemarkModifyRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Markees, 100, "markees"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRelationRemarkModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRelationRemarkModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type RemarkModel struct {
	RemarkName string `json:"remark_name,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type OapiRelationRemarkModifyResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
