package grpc

import (
	"context"
	"encoding/json"
	dh "indexer/data_handler"
)

// ServerImpl implements the Server gRPC service
type ServerImpl struct {
	UnimplementedServerServer
}

// GetPositionsJson returns positions data in JSON format
func (s *ServerImpl) GetPositionsJson(_ context.Context, _ *Empty) (*PositionsResponse, error) {
	// Convert positions map to JSON
	data, err := json.Marshal(dh.Positions)
	if err != nil {
		return nil, err
	}

	return &PositionsResponse{
		Data: data,
	}, nil
}
