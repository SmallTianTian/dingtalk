package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataFollowrecordListRequest() *OapiCrmObjectdataFollowrecordListRequest {
	return &OapiCrmObjectdataFollowrecordListRequest{}
}

type OapiCrmObjectdataFollowrecordListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataFollowrecordListResponse
	CurrentOperatorUserid string
	DataIdList            string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataFollowrecordListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataFollowrecordListRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataFollowrecordListRequest) SetDataIdList(dataIdList2 string) {
	this.DataIdList = dataIdList2
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetDataIdList() string {
	return this.DataIdList
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.followrecord.list"
}
func (this *OapiCrmObjectdataFollowrecordListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataFollowrecordListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataFollowrecordListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["data_id_list"] = this.DataIdList
	return txtParams
}
func (this *OapiCrmObjectdataFollowrecordListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataIdList, "dataIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataFollowrecordListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataFollowrecordListResponse struct {
	taobao.TaobaoResponse

	ResultList []ObjectDataInstanceVo `json:"result_list,omitempty"`
}
