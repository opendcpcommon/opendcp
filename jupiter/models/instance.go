// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"time"
)

type InstanceStatus int

const (
	Pending     InstanceStatus = iota //正在创建
	Success                           //初始化完成
	Uninit                            //未初始化
	Initing                           //正在初始化
	InitTimeout                       //初始化超时
	Deleted                           //资源已删除
	Deleting                          //正在删除
	StatusError                       //错误状态
)

// State represents the state of a host
type InstanceState int

const (
	None InstanceState = iota
	Running
	Paused
	Saved
	Stopped
	Stopping
	Starting
	StateError
)

// Describes a tag.
type Tag struct {

	// The key of the tag.
	//
	// Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode
	// characters. May not begin with aws:
	Key string `locationName:"key" type:"string"`

	// The value of the tag.
	//
	// Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode
	// characters.
	Value string `locationName:"value" type:"string"`
}

// Describes a product code.
type ProductCode struct {

	// The product code.
	ProductCodeId string `locationName:"productCode" type:"string"`

	// The type of product code.
	ProductCodeType string `locationName:"type" type:"string" enum:"ProductCodeValues"`
}

// Contains the output of CreateSecurityGroups.
type SecurityGroupsResp struct {
	// Information about one or more security groups.
	SecurityGroups []SecurityGroup `locationName:"securityGroupInfo" locationNameList:"item" type:"list"`
}

type ListInstancesResponse struct {

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	// Zero or more reservations.
	Reservations []InstanceAllIn `locationName:"reservationSet" locationNameList:"item" type:"list"`
}

type Instance struct {
	Id                 int      `orm:"pk;auto"`
	Cluster            *Cluster `orm:"rel(fk);on_delte(do_nothing)"`
	Provider           string
	CreateTime         time.Time `orm:"auto_now_add;type(datetime)"`
	Cpu                int
	Ram                int
	InstanceId         string
	ImageId            string
	InstanceType       string
	KeyName            string
	// The ID of the VPC.
	VpcId              string
	// The ID of the subnet.
	SubnetId           string
	// The ID of the security group
	SecurityGroupId    string
	RegionId           string
	ZoneId             string
	DataDiskNum        int
	DataDiskSize       int
	DataDiskCategory   string
	SystemDiskCategory string
	CostWay            string
	PreBuyMonth        int
	// The IP address of the network interface within the subnet.
	PrivateIpAddress   string
	// The public IP address or Elastic IP address bound to the network interface.
	PublicIpAddress    string
	NatIpAddress       string
	// The status of the network interface.
	Status             InstanceStatus
	PublicKey          string `orm:"type(text);null" json:"-"`
	PrivateKey         string `orm:"type(text);null" json:"-"`
}

type StatusResp struct {
	InstanceId       string         `json:"instance_id"`
	Status           InstanceStatus `json:"status"`
	IpAddress        string         `json:"ip_address"`
}

type SecurityGroupIdSetTypeAllin struct {

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	// Zero or more reservations.
	AllSecurityGroupId []string `json:"SecurityGroupId"`
}
type IpAddressSetTypeAllin struct {
	AllIpAddress []string `json:"IpAddress"`
}
type VpcAttributesTypeAllin struct {
	VpcId            string                `json:"VpcId"`
	VSwitchId        string                `json:"VSwitchId"`
	PrivateIpAddress IpAddressSetTypeAllin `json:"PrivateIpAddress"`
	NatIpAddress     string                `json:"NatIpAddress"`
}

type EipAddressAssociateTypeAllin struct {
	AllocationId       string `json:"AllocationId"`
	IpAddress          string `json:"IpAddress"`
	Bandwidth          int    `json:"Bandwidth"`
	InternetChargeType string `json:"InternetChargeType"`
}
type InstanceAllIn struct {
	InstanceId          string                       `json:"InstanceId"`
	InstanceName        string                       `json:"InstanceName"`
	Description         string                       `json:"Description"`
	ImageId             string                       `json:"ImageId"`
	RegionId            string                       `json:"RegionId"`
	ZoneId              string                       `json:"ZoneId"`
	InstanceType        string                       `json:"InstanceType"`
	Status              string                       `json:"Status"`
	AllSecurityGroupIds SecurityGroupIdSetTypeAllin  `json:"SecurityGroupIds"`
	PublicIpAddress     IpAddressSetTypeAllin        `json:"PublicIpAddress"`
	InnerIpAddress      IpAddressSetTypeAllin        `json:"InnerIpAddress"`
	CreationTime        string                       `json:"CreationTime"`
	VpcAttributes       VpcAttributesTypeAllin       `json:"VpcAttributes"`
	EipAddress          EipAddressAssociateTypeAllin `json:"EipAddress"`
	ExpiredTime         string                       `json:"ExpiredTime"` // 过期时间，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

// Describes an IP range.
type IpRange struct {

	// The CIDR range. You can either specify a CIDR range or a source security
	// group, not both.
	CidrIp string `locationName:"cidrIp" type:"string"`
}

// The ID of the prefix.
type PrefixListId struct {

	// The ID of the prefix.
	PrefixListId string `locationName:"prefixListId" type:"string"`
}

// Describes a state change.
type StateReason struct {

	// The reason code for the state change.
	Code string `locationName:"code" type:"string"`

	// The message for the state change.
	//
	//    Server.SpotInstanceTermination: A Spot instance was terminated due to
	// an increase in the market price.
	//
	//    Server.InternalError: An internal error occurred during instance launch,
	// resulting in termination.
	//
	//    Server.InsufficientInstanceCapacity: There was insufficient instance
	// capacity to satisfy the launch request.
	//
	//    Client.InternalError: A client error caused the instance to terminate
	// on launch.
	//
	//    Client.InstanceInitiatedShutdown: The instance was shut down using the
	// shutdown -h command from the instance.
	//
	//    Client.UserInitiatedShutdown: The instance was shut down using the Amazon
	// EC2 API.
	//
	//    Client.VolumeLimitExceeded: The limit on the number of EBS volumes or
	// total storage was exceeded. Decrease usage or request an increase in your
	// limits.
	//
	//    Client.InvalidSnapshot.NotFound: The specified snapshot was not found.
	Message string `locationName:"message" type:"string"`
}