package service

import (
	"context"
	"errors"

	pb "../pb"
)

type TararebaService struct {
}

func NewTararebaService() *TararebaService {
	return &TararebaService{}
}

func (s *TararebaService) GetCompetitionHistory(ctx context.Context, message *pb.GetCompetitionHistoryRequest) (*pb.GetCompetitionHistoryResponse, error) {
	switch message.UserId {
	case "monkukui":
		return &pb.GetCompetitionHistoryResponse{
			IsRated:        true,
			Place:          429,
			OldRating:      0,
			NewRating:      11,
			Performance:    157,
			ContestName:    "AtCoder Beginner Contest 058",
			IsParticipated: true,
		}, nil
	case "olphe":
		return &pb.GetCompetitionHistoryResponse{
			IsRated:        true,
			Place:          525,
			OldRating:      0,
			NewRating:      146,
			Performance:    1197,
			ContestName:    "AtCoder Grand Contest 003",
			IsParticipated: true,
		}, nil
	}

	return nil, errors.New("user is not monkukui or olphe")
}
