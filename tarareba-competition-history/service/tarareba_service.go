package service

import (
	"context"
	"errors"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/tarareba_competition_history_pb"
)

type TararebaService struct {
}

func NewTararebaService() *TararebaService {
	return &TararebaService{}
}

func (s *TararebaService) GetCompetitionHistory(ctx context.Context, message *pb.GetCompetitionHistoryRequest) (*pb.GetCompetitionHistoryResponse, error) {

	competition_history := &pb.CompetitionHistory{
		IsRated:           true,
		Place:             429,
		OldRating:         0,
		NewRating:         111,
		Performance:       1571,
		InnerPerformance:  2001111111,
		ContestScreenName: "atcoder.com",
		ContestName:       "AtCoder Beginner Contest 058",
		ContestNameEn:     "",
		EndTime:           "2020-02-02",
	}

	var res []*pb.CompetitionHistory
	res = append(res, competition_history)
	res = append(res, competition_history)
	res = append(res, competition_history)

	switch message.UserId {
	case "monkukui":
		return &pb.GetCompetitionHistoryResponse{
			CompetitionHistory: res,
		}, nil
	}

	return nil, errors.New("user is not monkukui or olphe")
}
