package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserRelationGetRequest() *OapiEduUserRelationGetRequest {
	return &OapiEduUserRelationGetRequest{}
}

type OapiEduUserRelationGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserRelationGetResponse
	ClassId         int64
	FromUserid      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduUserRelationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserRelationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserRelationGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduUserRelationGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduUserRelationGetRequest) SetFromUserid(fromUserid2 string) {
	this.FromUserid = fromUserid2
}
func (this *OapiEduUserRelationGetRequest) GetFromUserid() string {
	return this.FromUserid
}
func (this *OapiEduUserRelationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.relation.get"
}
func (this *OapiEduUserRelationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserRelationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserRelationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserRelationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserRelationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["from_userid"] = this.FromUserid
	return txtParams
}
func (this *OapiEduUserRelationGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserRelationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserRelationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserRelationGetResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type OpenEduUserRelationDetail struct {
	ClassId      int64  `json:"class_id,omitempty"`
	FromUserid   string `json:"from_userid,omitempty"`
	RelationCode string `json:"relation_code,omitempty"`
	RelationName string `json:"relation_name,omitempty"`
	ToUserid     string `json:"to_userid,omitempty"`
}
