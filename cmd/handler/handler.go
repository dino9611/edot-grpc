package handler

import (
	"context"

	charpb "github.com/tes-edot/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	charpb.UnimplementedChartServer
}

func (*Handler) GetChart(ctx context.Context, in *emptypb.Empty) (*charpb.Chart, error) {
	result := &charpb.Chart{
		Data: []int64{29, 30, 70, 79, 40, 69, 100},
	}
	return result, nil
}
