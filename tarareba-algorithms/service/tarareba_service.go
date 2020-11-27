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

// GetOptimalHistory は、パフォーマンス列を入力として受け取り、レートを最大化する関数です
func (s *TararebaService) GetOptimalHistory(ctx context.Context, message *pb.GetOptimalHistoryRequest) (*pb.GetOptimalHistoryResponse, error) {

	var res []*pb.OptimalHistory
	rating := int32(0)
	for i := 0; i < len(message.ActualHistory); i++ {
		res = append(res, &pb.OptimalHistory{
			OldRating:      rating,
			NewRating:      rating + 100,
			IsParticipated: (int32(i)*rating)%5 > 0,
		})
		rating += 100
	}

	return &pb.GetOptimalHistoryResponse{
		OptimalHistory: res,
	}, nil
}

// GetRatingTransition は、パフォーマンス列を入力として受け取り、isPerticipated が true のコンテストだけでレート推移を計算する関数です
func (s *TararebaService) GetRatingTransition(ctx context.Context, message *pb.GetRatingTransitionRequest) (*pb.GetRatingTransitionResponse, error) {

	var res []*pb.RatingTransition
	rating := int32(0)
	for i := 0; i < len(message.ContestPerformance); i++ {
		res = append(res, &pb.RatingTransition{
			OldRating: rating,
			NewRating: rating + 100,
		})
		rating += 100
	}

	return &pb.GetRatingTransitionResponse{
		RatingTransition: res,
	}, nil
}
