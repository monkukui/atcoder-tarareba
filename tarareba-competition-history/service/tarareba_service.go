package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/tarareba_competition_history_pb"
)

type TararebaService struct {
}

func NewTararebaService() *TararebaService {
	return &TararebaService{}
}

func (s *TararebaService) GetCompetitionHistory(ctx context.Context, message *pb.GetCompetitionHistoryRequest) (*pb.GetCompetitionHistoryResponse, error) {

	// コンテスト成績表に関する JSON を取得する
	// コンテスト成績表はここを参照：https://atcoder.jp/users/monkukui/history
	// コンテスト成績表に関する JSON は、この URL で取れる：https://atcoder.jp/users/monkukui/history/json

	url := "https://atcoder.jp/users/" + message.UserId + "/history/json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var contests Contests
	if err := json.Unmarshal(body, &contests); err != nil {
		return nil, err
	}

	history := make([]*pb.CompetitionHistory, 0, len(contests))

	for _, contest := range contests {
		history = append(history, &pb.CompetitionHistory{
			IsRated:           contest.IsRated,
			Place:             uint32(contest.Place),
			OldRating:         uint32(contest.OldRating),
			NewRating:         uint32(contest.NewRating),
			Performance:       uint32(contest.Performance),
			InnerPerformance:  uint32(contest.InnerPerformance),
			ContestScreenName: contest.ContestScreenName,
			ContestName:       contest.ContestName,
			ContestNameEn:     contest.ContestNameEn,
			EndTime:           contest.EndTime,
		})
	}

	return &pb.GetCompetitionHistoryResponse{
		CompetitionHistory: history,
	}, nil
}

type (
	Contest struct {
		IsRated           bool   `json:"IsRated"`
		Place             int    `json:"Place"`
		OldRating         int    `json:"OldRating"`
		NewRating         int    `json:"NewRating"`
		Performance       int    `json:"Performance"`
		InnerPerformance  int    `json:"InnerPerformance"`
		ContestScreenName string `json:"ContestScreenName"`
		ContestName       string `json:"ContestName"`
		ContestNameEn     string `json:"ContestNameEn"`
		EndTime           string `json:"EndTime"`
	}

	Contests []*Contest
)
