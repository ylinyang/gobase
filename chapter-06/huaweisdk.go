package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
	"log"
)

var (
	ak = "9V7YUT1CKGYOCY9N8KGS"
	sk = "sq6m5wpC5RX6ecJrDwNSgipjOeqWphsEDZk2auKj"
)

func base() *vpc.VpcClient {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())
	return client
}

func ListAddressGroup() []string {
	client := base()
	request := &model.ListAddressGroupRequest{}
	response, err := client.ListAddressGroup(request)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var g = []string{}
	for _, group := range *response.AddressGroups {
		g = append(g, group.Id)
	}
	return g
}

func ShowAddressGroup(id string) map[string][]string {
	client := base()
	request := &model.ShowAddressGroupRequest{}
	request.AddressGroupId = id
	response, err := client.ShowAddressGroup(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	m := make(map[string][]string)
	m[response.AddressGroup.Id] = response.AddressGroup.IpSet
	//fmt.Println(m)
	return m
}

func UpdateAddressGroup(add []string, id string, act string) error {
	client := base()
	request := &model.UpdateAddressGroupRequest{}
	request.AddressGroupId = id

	address := ShowAddressGroup(id)
	var listIpSetAddressGroup = address[id]
	if act == "add" {
		listIpSetAddressGroup = append(listIpSetAddressGroup, add...)

	} else if act == "del" {
		listIpSetAddressGroup = del(listIpSetAddressGroup, add)
	} else if act == "del all" {
		listIpSetAddressGroup = add
	}
	descriptionAddressGroup := "信任地址"
	addressGroupbody := &model.UpdateAddressGroupOption{
		Description: &descriptionAddressGroup,
		IpSet:       &listIpSetAddressGroup,
	}
	request.Body = &model.UpdateAddressGroupRequestBody{
		AddressGroup: addressGroupbody,
	}
	_, err := client.UpdateAddressGroup(request)
	if err != nil {
		return err
	}
	return nil
}

func del(src, dst []string) (m []string) {
	fmt.Println(src)
	fmt.Println(dst)
	for _, s := range src {
		for _, d := range dst {
			if s != d {
				m = append(m, s)
			}
		}
	}
	fmt.Println(m)
	return
}
