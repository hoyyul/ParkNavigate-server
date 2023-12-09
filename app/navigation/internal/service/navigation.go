package service

import (
	"ParkNavigate/global"
	"ParkNavigate/pkg/pb/navigation_pb"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type NavigationService struct {
	navigation_pb.UnimplementedNavigationServiceServer
}

func NewNavigationService() *NavigationService {
	return &NavigationService{}
}

func (s *NavigationService) ShowParkings(ctx context.Context, req *navigation_pb.NavigationRequest) (*navigation_pb.NavigationResponse, error) {
	url := "https://piomatixlbs.stg.pioneerapis.com/lbsapi/searchfe/search/v1/freeword"
	searchReq := makeSearchRequest(req)
	b, _ := json.Marshal(searchReq)

	// 创建一个新的 HTTP 请求
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		global.Logger.Panic(err)
		return nil, err
	}

	// 设置 header
	httpReq.Header.Set("Content-Type", "application/json;charset=utf-8")
	httpReq.Header.Set("Authorization", "PAK1 Credential=zPsJKOJCxNbkpddLbLoZ Signature=6a75i625tzfqok3mchxe79831o2t5u0gccl9np0eyf3mq61e29megt8ml3f2htky")
	httpReq.Header.Set("PEC-Traffic-ProviderKey", "none")
	httpReq.Header.Set("PEC-Traffic-ProviderUserID", "testUser")

	// 发送请求
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		global.Logger.Panic(err)
		return nil, err
	}
	defer httpResp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		global.Logger.Panic(err)
		return nil, err
	}

	global.Logger.Println("Response:", string(body))

	// 解析响应体到 searchResp 结构体
	var searchResp *SearchResponse
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		global.Logger.Panic(err)
		return nil, err
	}

	// 保存到 NavigationResponse
	resp := &navigation_pb.NavigationResponse{}
	for _, place := range searchResp.PlaceList {
		resp.Location = append(resp.Location, &navigation_pb.Location{
			Latitude:  place.Latitude,
			Longitude: place.Longitude,
		})
	}

	return resp, nil
}

func makeSearchRequest(req *navigation_pb.NavigationRequest) *SearchRequest {
	return &SearchRequest{
		Resource:  []string{"z_poi", "z_address"},
		Keyword:   "駐車場",
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Radius:    1000,
		SortType:  1,
	}
}
