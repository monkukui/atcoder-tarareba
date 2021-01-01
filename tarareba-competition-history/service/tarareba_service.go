package service

import (
	"context"
	"encoding/json"
	"fmt"
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
			Place:             contest.Place,
			OldRating:         contest.OldRating,
			NewRating:         contest.NewRating,
			Performance:       contest.Performance,
			InnerPerformance:  contest.InnerPerformance,
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
		Place             int32  `json:"Place"`
		OldRating         int32  `json:"OldRating"`
		NewRating         int32  `json:"NewRating"`
		Performance       int32  `json:"Performance"`
		InnerPerformance  int32  `json:"InnerPerformance"`
		ContestScreenName string `json:"ContestScreenName"`
		ContestName       string `json:"ContestName"`
		ContestNameEn     string `json:"ContestNameEn"`
		EndTime           string `json:"EndTime"`
	}

	Contests []*Contest
)
