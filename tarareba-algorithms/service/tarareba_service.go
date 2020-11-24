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
	rating := int32(0)
	for i := 0; i < len(message.ActualHistory); i++ {
		res = append(res, &pb.OptimalHistory{
			OldRating:      rating,
			NewRating:      rating + 100,
			IsParticipated: (int32(i)*rating)%5 > 0,
		})
		rating += 31
	}

	return &pb.GetOptimalHistoryResponse{
		OptimalHistory: res,
	}, nil
}
