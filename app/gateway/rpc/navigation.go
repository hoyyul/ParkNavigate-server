package rpc

import (
	"ParkNavigate/pkg/pb/navigation_pb"
	"context"
)

func ShowParkings(ctx context.Context, req *navigation_pb.NavigationRequest) (*navigation_pb.NavigationResponse, error) {
	r, err := NavigationClient.ShowParkings(ctx, req)
	if err != nil {
		return nil, err
	}

	return r, nil
}
