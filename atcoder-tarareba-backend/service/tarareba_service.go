package service

import (
	"context"
	"errors"

	pb "github.com/monkukui/atcoder-tarareba/atcoder-tarareba-backend/pb"
)

type TararebaService struct {
}

func (s *TararebaService) GetCompetitionHistory(ctx context.Context, message *pb.GetCompetitionHistoryRequest) (*pb.GetCompetitionHistoryResponse, error) {
	switch message.UserId {
	case "monkukui":
		return &pb.GetCompetitionHistoryResponse{
			IsRated:   true,
			Place:     429,
			OldRating: 0,
			NewRating: 11,
			InnerPerformance, 157,
			Performance, 157,
			ContestScreenName: "abc058.contest.atcoder.jp",
			ContestName:       "AtCoder Beginner Contest 058",
			ContestNameEn:     "",
			IsParticipated:    true,
		}, nil
	case "olphe":
		return &pb.GetCompetitionHistoryResponse{
			IsRated:   true,
			Place:     525,
			OldRating: 0,
			NewRating: 146,
			InnerPerformance, 1197,
			Performance, 1197,
			ContestScreenName: "agc003.contest.atcoder.jp",
			ContestName:       "AtCoder Grand Contest 003",
			ContestNameEn:     "",
			IsParticipated:    true,
		}, nil
	}

	return nil, errors.New("user is not monkukui or olphe")
}
