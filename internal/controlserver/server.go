package controlserver

import (
	"context"

	apipb "github.com/troydai/control-surface/gen/proto/com/troydai/controlsurface/v1"
)

func New() apipb.ControlServiceServer {
	return &impl{}
}

type impl struct {
	apipb.UnimplementedControlServiceServer
}

func (s *impl) GetParameters(context.Context, *apipb.GetParametersRequest) (*apipb.GetParametersResponse, error) {
	resp := &apipb.GetParametersResponse{
		ParameterValues: []*apipb.ParameterValue{
			{
				Name: "foo",
				Value: &apipb.ParameterValue_FloatValue{
					FloatValue: 1.0,
				},
			},
		},
	}

	return resp, nil
}
