package service

import (
	"context"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-algorithms/tarareba_algorithms_pb"
)

type TararebaService struct {
}

func NewTararebaService() *TararebaService {
	return &TararebaService{}
}

func (s *TararebaService) GetOptimalHistory(ctx context.Context, message *pb.GetOptimalHistoryRequest) (*pb.GetOptimalHistoryResponse, error) {

	var res []*pb.OptimalHistory
	rating := 0
	for i := 0; i < len(message.ActualHistory); i++ {
		res = append(res, &pb.OptimalHistory{
			OldRating:      uint32(rating),
			NewRating:      uint32(rating + 100),
			IsParticipated: false,
		})
	}

	return &pb.GetOptimalHistoryResponse{
		OptimalHistory: res,
	}, nil
}
