package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupMemberListRequest() *OapiSmartdeviceFacegroupMemberListRequest {
	return &OapiSmartdeviceFacegroupMemberListRequest{}
}

type OapiSmartdeviceFacegroupMemberListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupMemberListResponse
	BizId           string
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.member.list"
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupMemberListResponse struct {
	taobao.TaobaoResponse
	Result  PagedList `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
