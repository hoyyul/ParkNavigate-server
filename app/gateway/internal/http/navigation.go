package http

import (
	"ParkNavigate/app/gateway/rpc"
	"ParkNavigate/global"
	"ParkNavigate/pkg/pb/navigation_pb"
	"ParkNavigate/pkg/resp"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParkingRequest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type FrontendParkingSpot struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsFull    bool    `json:"isFull"`
}

func ShowParkings(ctx *gin.Context) {
	var parkingReq ParkingRequest

	// 从请求体中绑定 JSON 数据
	if err := ctx.BindJSON(&parkingReq); err != nil {
		global.Logger.Panicln(err)
		resp.SendWithNotOk(http.StatusBadRequest, "Failed to bind request", ctx)
		return
	}

	// 使用绑定的数据创建 NavigationRequest
	req := navigation_pb.NavigationRequest{
		Latitude:  float32(parkingReq.Lat),
		Longitude: float32(parkingReq.Lng),
	}

	// 调用 RPC 服务
	r, err := rpc.ShowParkings(ctx, &req)
	if err != nil {
		global.Logger.Panicln(err)
		resp.SendWithNotOk(http.StatusInternalServerError, "Failed to call User RPC service", ctx)
		return
	}

	// 转换响应为前端需要的格式
	frontendResponse := make([]FrontendParkingSpot, len(r.Location))
	for i, loc := range r.Location {
		frontendResponse[i] = FrontendParkingSpot{
			Name:      "駐車場" + strconv.Itoa(i+1), // 为每个停车场生成名称
			Latitude:  float64(loc.Latitude),
			Longitude: float64(loc.Longitude),
			IsFull:    rand.Float32() < 0.5, // 随机决定是否满车
		}
	}

	// 将转换后的数据发送回客户端
	ctx.JSON(http.StatusOK, frontendResponse)
}
