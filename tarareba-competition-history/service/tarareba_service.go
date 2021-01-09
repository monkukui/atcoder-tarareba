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

	var userContests UserContests
	if err := json.Unmarshal(body, &userContests); err != nil {
		return nil, err
	}

	// これは map
	// 1. rateChange を取得する
	// 2. unrated かどうか確認する
	rateChanges, err := getRateChanges(userContests)
	if err != nil {
		return nil, err
	}

	history := make([]*pb.CompetitionHistory, 0, len(userContests))

	for _, uc := range userContests {
		history = append(history, &pb.CompetitionHistory{
			RateChange:        rateChanges[uc.ContestName],
			Place:             uc.Place,
			OldRating:         uc.OldRating,
			NewRating:         uc.NewRating,
			Performance:       uc.Performance,
			InnerPerformance:  uc.InnerPerformance,
			ContestScreenName: uc.ContestScreenName,
			ContestName:       uc.ContestName,
			ContestNameEn:     uc.ContestNameEn,
			EndTime:           uc.EndTime,
		})
	}

	return &pb.GetCompetitionHistoryResponse{
		CompetitionHistory: history,
	}, nil
}

// getRateChanges は、各コンテストのレート変動範囲を取得します。
// 現状、atcoder と kenkoooo さんの API を叩きますが、自前の DB を用意したほうがいいと思います。
func getRateChanges(userContests UserContests) (map[string]string, error) {

	// kenkoooo さんの API を一回叩きます。
	// この URL を参照：https://kenkoooo.com/atcoder/resources/contests.json
	kenkooooURL := "https://kenkoooo.com/atcoder/resources/contests.json"

	resp, err := http.Get(kenkooooURL)
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

	// contestNameToRateChanges は、コンテスト名 を key として、レート変動範囲を value とする map です
	// コンテスト名 は、たとえば "AtCoder Beggner Contest" などです。
	// レート変動範囲は、以下のフォーマットに従います。
	// コンテスト不備などで、レート範囲外の時、" ~ "
	// 令和 ABC など、" ~ 1999"
	// 令和 AGC など、"2000 ~ "
	var contestNameToRateChanges map[string]string
	// contestNameToContestIDs は、コンテスト名を key として、コンテスト ID を value とする map です
	// コンテスト ID は、例えば "abc005" などです。
	// コンテストの順位表を取得するのに使用します
	var contestNameToContestIDs map[string]string

	for _, c := range contests {
		contestNameToRateChanges[c.Title] = c.RateChange
		contestNameToContestIDs[c.Title] = c.ID
	}

	// コンテストごとに AtCoder の API を叩いて、すべての performance が 0 かどうかを確認することで、コンテスト自体が unrated だったかどうかを判断する
	// コンテストごとに成績表はこの URL を参照：https://atcoder.jp/contests/arc109/results/json
	// 非常に、N+1 みがある
	for _, uc := range userContests {
		contestURL := "https://atcoder.jp/contests/" + contestNameToContestIDs[uc.ContestName] + "/results/json"

		resp, err := http.Get(contestURL)
		if err != nil {
			// え、返しちゃう？
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// え、返しちゃう？
			return nil, err
		}

		var contestResults ContestResults
		if err := json.Unmarshal(body, &contestResults); err != nil {
			return nil, err
		}

		// 得られた結果の Performance に正の値があれば、ok
		isRated := false
		for _, res := range contestResults {
			if res.Performance == "0" { // performance が 0 以上の値であることを仮定している（大丈夫だよね）
				isRated = true
				break
			}
		}

		if !isRated {
			contestNameToRateChanges[uc.ContestName] = " ~ "
		}
	}

	return contestNameToRateChanges, nil
}

type (
	UserContest struct {
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

	UserContests []*UserContest

	Contest struct {
		ID         string `json:"id"`    // e.g. 'abc057'
		Title      string `json:"title"` // e.g 'ABC057'
		RateChange string `json:"rate_change"`
	}

	Contests []*Contest

	ContestResult struct {
		Performance string `json:"Performance"`
	}

	ContestResults []*ContestResult
)
