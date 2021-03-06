package slb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeRuleAttribute invokes the slb.DescribeRuleAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describeruleattribute.html
func (client *Client) DescribeRuleAttribute(request *DescribeRuleAttributeRequest) (response *DescribeRuleAttributeResponse, err error) {
	response = CreateDescribeRuleAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRuleAttributeWithChan invokes the slb.DescribeRuleAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeruleattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRuleAttributeWithChan(request *DescribeRuleAttributeRequest) (<-chan *DescribeRuleAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeRuleAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRuleAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeRuleAttributeWithCallback invokes the slb.DescribeRuleAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeruleattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRuleAttributeWithCallback(request *DescribeRuleAttributeRequest, callback func(response *DescribeRuleAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRuleAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRuleAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeRuleAttributeRequest is the request struct for api DescribeRuleAttribute
type DescribeRuleAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	RuleId               string           `position:"Query" name:"RuleId"`
}

// DescribeRuleAttributeResponse is the response struct for api DescribeRuleAttribute
type DescribeRuleAttributeResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RuleName       string `json:"RuleName" xml:"RuleName"`
	LoadBalancerId string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	ListenerPort   string `json:"ListenerPort" xml:"ListenerPort"`
	Domain         string `json:"Domain" xml:"Domain"`
	Url            string `json:"Url" xml:"Url"`
	VServerGroupId string `json:"VServerGroupId" xml:"VServerGroupId"`
}

// CreateDescribeRuleAttributeRequest creates a request to invoke DescribeRuleAttribute API
func CreateDescribeRuleAttributeRequest() (request *DescribeRuleAttributeRequest) {
	request = &DescribeRuleAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeRuleAttribute", "slb", "openAPI")
	return
}

// CreateDescribeRuleAttributeResponse creates a response to parse from DescribeRuleAttribute response
func CreateDescribeRuleAttributeResponse() (response *DescribeRuleAttributeResponse) {
	response = &DescribeRuleAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
