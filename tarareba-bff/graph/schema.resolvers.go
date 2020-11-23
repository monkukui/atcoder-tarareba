package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/generated"
	"github.com/monkukui/atcoder-tarareba/tarareba-bff/graph/model"
	"github.com/monkukui/atcoder-tarareba/tarareba-bff/pb"
	"google.golang.org/grpc"
)

func (r *queryResolver) ContestsByUserID(ctx context.Context, userID *string) ([]*model.Contest, error) {

	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewTararebaServiceClient(conn)

	message := &pb.GetCompetitionHistoryRequest{UserId: *userID}
	res, err := client.GetCompetitionHistory(context.TODO(), message)
	if err != nil {
		return nil, err
	}

	contests := make([]*model.Contest, 0, len(res.CompetitionHistory))

	for _, contest := range res.CompetitionHistory {
		contests = append(contests, &model.Contest{
			IsRated:           contest.IsRated,
			Place:             int(contest.Place),
			OldRating:         int(contest.OldRating),
			NewRating:         int(contest.NewRating),
			Performance:       int(contest.Performance),
			InnerPerformance:  int(contest.InnerPerformance),
			ContestScreenName: contest.ContestScreenName,
			ContestName:       contest.ContestName,
			ContestNameEn:     contest.ContestNameEn,
			EndTime:           contest.EndTime,
			IsParticipated:    contest.IsParticipated,
		})
	}

	return contests, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
