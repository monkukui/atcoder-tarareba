package service

import (
	"testing"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/tarareba_competition_history_pb"
	"github.com/stretchr/testify/assert"
)

// 最初に受けたコンテストが取得できることをテストします。
func TestCompetitionHistory_CanGetCompetitionHistory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		userID       string
		firstContest *pb.CompetitionHistory
	}{
		{
			name:   "ユーザーが monkukui のとき",
			userID: "monkukui",
			firstContest: &pb.CompetitionHistory{
				IsRated:           true,
				Place:             429,
				OldRating:         0,
				NewRating:         11,
				Performance:       157,
				InnerPerformance:  157,
				ContestScreenName: "abc058.contest.atcoder.jp",
				ContestName:       "AtCoder Beginner Contest 058",
				ContestNameEn:     "",
				EndTime:           "2017-04-08T22:40:00+09:00",
			},
		},
		{
			name:   "ユーザーが olphe のとき",
			userID: "olphe",
			firstContest: &pb.CompetitionHistory{
				IsRated:           true,
				Place:             525,
				OldRating:         0,
				NewRating:         146,
				Performance:       1197,
				InnerPerformance:  1197,
				ContestScreenName: "agc003.contest.atcoder.jp",
				ContestName:       "AtCoder Grand Contest 003",
				ContestNameEn:     "",
				EndTime:           "2016-08-21T22:50:00+09:00",
			},
		},
		{
			name:   "ユーザーが chokudai のとき",
			userID: "chokudai",
			firstContest: &pb.CompetitionHistory{
				IsRated:           true,
				Place:             59,
				OldRating:         0,
				NewRating:         1255,
				Performance:       2455,
				InnerPerformance:  2455,
				ContestScreenName: "arc061.contest.atcoder.jp",
				ContestName:       "AtCoder Regular Contest 061",
				ContestNameEn:     "",
				EndTime:           "2016-09-11T22:40:00+09:00",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTararebaService()
			request := pb.GetCompetitionHistoryRequest{UserId: tt.userID}
			response, err := s.GetCompetitionHistory(nil, &request)

			assert.NoError(t, err)
			assert.NotEqual(t, 0, len(response.CompetitionHistory))
			assert.Equal(t, tt.firstContest, response.CompetitionHistory[0])
		})
	}
}

// 存在しないユーザー or コンテストに出場していないとき、空配列が変えることをテストします。
func TestCompetitionHistory_ReturnEmptySlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		userID string
	}{
		{
			name:   "存在しないユーザーのとき",
			userID: "not_exist_user_987315073764316328",
		},
		{
			name:   "存在はするが、コンテストへの出場歴がないとき",
			userID: "hitachi_admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTararebaService()
			request := pb.GetCompetitionHistoryRequest{UserId: tt.userID}
			response, err := s.GetCompetitionHistory(nil, &request)

			assert.NoError(t, err)
			assert.Equal(t, 0, len(response.CompetitionHistory))
		})
	}
}
