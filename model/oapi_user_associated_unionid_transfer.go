package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserAssociatedUnionidTransferRequest() *OapiUserAssociatedUnionidTransferRequest {
	return &OapiUserAssociatedUnionidTransferRequest{}
}

type OapiUserAssociatedUnionidTransferRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiUserAssociatedUnionidTransferResponse
	AssociatedUnionid string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiUserAssociatedUnionidTransferRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserAssociatedUnionidTransferRequest) SetAssociatedUnionid(associatedUnionid2 string) {
	this.AssociatedUnionid = associatedUnionid2
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetAssociatedUnionid() string {
	return this.AssociatedUnionid
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.associated_unionid.transfer"
}
func (this *OapiUserAssociatedUnionidTransferRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserAssociatedUnionidTransferRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserAssociatedUnionidTransferRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["associated_unionid"] = this.AssociatedUnionid
	return txtParams
}
func (this *OapiUserAssociatedUnionidTransferRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AssociatedUnionid, "associatedUnionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserAssociatedUnionidTransferRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserAssociatedUnionidTransferResponse struct {
	taobao.TaobaoResponse
	Result UseridVo `json:"result,omitempty"`
}
type UseridVo struct {
	ContactType int64  `json:"contact_type,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
