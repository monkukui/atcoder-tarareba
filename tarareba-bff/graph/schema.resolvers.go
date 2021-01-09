package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/generated"
	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/model"
	pbAlgorithms "github.com/monkukui/atcoder-tarareba/tarareba-bff/tarareba_algorithms_pb"
	pbHistory "github.com/monkukui/atcoder-tarareba/tarareba-bff/tarareba_competition_history_pb"
	"google.golang.org/grpc"
)

func (r *queryResolver) ContestsByUserID(ctx context.Context, userID *string) ([]*model.Contest, error) {
	connHistory, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	// connHistory, err := grpc.Dial("tarareba-competition-history:19003", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer connHistory.Close()
	clientHistory := pbHistory.NewTararebaServiceClient(connHistory)

	// マイクロサービス `tarareba_competition_history` から、コンテスト情報を取得する
	messageHistory := &pbHistory.GetCompetitionHistoryRequest{UserId: *userID}
	resHistory, err := clientHistory.GetCompetitionHistory(context.TODO(), messageHistory)
	if err != nil {
		return nil, err
	}

	// マイクロサービス `tarareba_algorithms` へリクエストを送るための準備をする
	actualHistory := make([]*pbAlgorithms.ActualHistory, 0, len(resHistory.CompetitionHistory))
	for _, contest := range resHistory.CompetitionHistory {
		actualHistory = append(actualHistory, &pbAlgorithms.ActualHistory{
			RateChange:       contest.GetRateChange(),
			Performance:      contest.GetPerformance(),
			InnerPerformance: contest.GetInnerPerformance(),
		})
	}

	connAlgorithms, err := grpc.Dial("127.0.0.1:19004", grpc.WithInsecure())
	// connAlgorithms, err := grpc.Dial("tarareba-algorithms:19004", grpc.WithInsecure())
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
			RateChange:        contest.GetRateChange(),
			Place:             int(contest.GetPlace()),
			ActualOldRating:   int(contest.GetOldRating()),
			ActualNewRating:   int(contest.GetNewRating()),
			Performance:       int(contest.GetPerformance()),
			InnerPerformance:  int(contest.GetInnerPerformance()),
			ContestScreenName: contest.GetContestScreenName(),
			ContestName:       contest.GetContestName(),
			ContestNameEn:     contest.GetContestNameEn(),
			EndTime:           contest.GetEndTime(),
			OptimalOldRating:  int(optimal.GetOldRating()),
			OptimalNewRating:  int(optimal.GetNewRating()),
			IsParticipated:    optimal.GetIsParticipated(),
		})
	}

	return contests, nil
}

func (r *queryResolver) RatingTransitionByPerformance(ctx context.Context, rateChange []*string, isParticipated []*bool, performances []*int, innerPerformances []*int) ([]*model.RatingTransition, error) {
	if len(isParticipated) != len(performances) {
		panic("")
	}
	if len(isParticipated) != len(innerPerformances) {
		panic("")
	}

	connAlgorithms, err := grpc.Dial("127.0.0.1:19004", grpc.WithInsecure())
	// connAlgorithms, err := grpc.Dial("tarareba-algorithms:19004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clientAlgorithms := pbAlgorithms.NewTararebaServiceClient(connAlgorithms)

	contestPerformance := make([]*pbAlgorithms.ContestPerformance, 0, len(isParticipated))

	for i := 0; i < len(isParticipated); i++ {
		contestPerformance = append(contestPerformance, &pbAlgorithms.ContestPerformance{
			RateChange:       *rateChange[i],
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
			OldRating: int(transition.GetOldRating()),
			NewRating: int(transition.GetNewRating()),
		})
	}

	return transitions, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
