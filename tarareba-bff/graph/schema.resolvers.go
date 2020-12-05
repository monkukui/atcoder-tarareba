package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/generated"
	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/model"
	pbAlgorithms "github.com/monkukui/atcoder-tarareba/tarareba-bff/tarareba_algorithms_pb"
	pbHistory "github.com/monkukui/atcoder-tarareba/tarareba-bff/tarareba_competition_history_pb"
	"google.golang.org/grpc"
)

// ContestsByUserID は、AtCoder ID を入力として受け取り、レートを最大化したコンテスト情報を返します
func (r *queryResolver) ContestsByUserID(ctx context.Context, userID *string) ([]*model.Contest, error) {
	connHistory, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	fmt.Println("debug", 20)
	if err != nil {
		fmt.Println("hahihuheho")
		return nil, err
	}
	defer connHistory.Close()
	clientHistory := pbHistory.NewTararebaServiceClient(connHistory)

	// マイクロサービス `tarareba_competition_history` から、コンテスト情報を取得する
	messageHistory := &pbHistory.GetCompetitionHistoryRequest{UserId: *userID}
	resHistory, err := clientHistory.GetCompetitionHistory(context.TODO(), messageHistory)
	fmt.Println("debug", 31)
	if err != nil {
		return nil, err
	}

	fmt.Println("debug", 36)

	// マイクロサービス `tarareba_algorithms` へリクエストを送るための準備をする
	actualHistory := make([]*pbAlgorithms.ActualHistory, 0, len(resHistory.CompetitionHistory))
	for _, contest := range resHistory.CompetitionHistory {
		actualHistory = append(actualHistory, &pbAlgorithms.ActualHistory{
			IsRated:          contest.IsRated,
			Performance:      contest.Performance,
			InnerPerformance: contest.InnerPerformance,
		})
	}

	connAlgorithms, err := grpc.Dial("127.0.0.1:19004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clientAlgorithms := pbAlgorithms.NewTararebaServiceClient(connAlgorithms)

	messageAlgorithms := &pbAlgorithms.GetOptimalHistoryRequest{
		ActualHistory: actualHistory,
	}
	// マイクロサービス `tarareba_algorithms` から参加履歴を取得する
	resAlgorithms, err := clientAlgorithms.GetOptimalHistory(context.TODO(), messageAlgorithms)
	if err != nil {
		return nil, err
	}

	if len(resHistory.CompetitionHistory) != len(resAlgorithms.OptimalHistory) {
		panic("")
	}

	contests := make([]*model.Contest, 0, len(resHistory.CompetitionHistory))

	for i, contest := range resHistory.CompetitionHistory {
		optimal := resAlgorithms.OptimalHistory[i]

		contests = append(contests, &model.Contest{
			IsRated:           contest.IsRated,
			Place:             int(contest.Place),
			ActualOldRating:   int(contest.OldRating),
			ActualNewRating:   int(contest.NewRating),
			Performance:       int(contest.Performance),
			InnerPerformance:  int(contest.InnerPerformance),
			ContestScreenName: contest.ContestScreenName,
			ContestName:       contest.ContestName,
			ContestNameEn:     contest.ContestNameEn,
			EndTime:           contest.EndTime,
			OptimalOldRating:  int(optimal.OldRating),
			OptimalNewRating:  int(optimal.NewRating),
			IsParticipated:    optimal.IsParticipated,
		})
	}

	return contests, nil
}

// RatingTransitionByPerformance は、パフォーマンス列を入力として受け取り、レートの推移列を返します。
func (r *queryResolver) RatingTransitionByPerformance(ctx context.Context, isParticipated []*bool, performances []*int, innerPerformances []*int) ([]*model.RatingTransition, error) {

	if len(isParticipated) != len(performances) {
		panic("")
	}
	if len(isParticipated) != len(innerPerformances) {
		panic("")
	}

	connAlgorithms, err := grpc.Dial("127.0.0.1:19004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clientAlgorithms := pbAlgorithms.NewTararebaServiceClient(connAlgorithms)

	contestPerformance := make([]*pbAlgorithms.ContestPerformance, 0, len(isParticipated))

	for i := 0; i < len(isParticipated); i++ {
		contestPerformance = append(contestPerformance, &pbAlgorithms.ContestPerformance{
			IsParticipated:   *isParticipated[i],
			Performance:      int32(*performances[i]),
			InnerPerformance: int32(*innerPerformances[i]),
		})
	}

	messageAlgorithms := &pbAlgorithms.GetRatingTransitionRequest{
		ContestPerformance: contestPerformance,
	}

	// マイクロサービス `tarareba_algorithms` からレート推移を取得する
	resAlgorithms, err := clientAlgorithms.GetRatingTransition(context.TODO(), messageAlgorithms)
	if err != nil {
		return nil, err
	}

	if len(isParticipated) != len(resAlgorithms.RatingTransition) {
		panic("")
	}

	transitions := make([]*model.RatingTransition, 0, len(resAlgorithms.RatingTransition))

	for _, transition := range resAlgorithms.RatingTransition {
		transitions = append(transitions, &model.RatingTransition{
			OldRating: int(transition.OldRating),
			NewRating: int(transition.NewRating),
		})
	}

	return transitions, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
