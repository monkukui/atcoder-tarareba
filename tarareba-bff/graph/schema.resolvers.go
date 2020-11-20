package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"tarareba-bff/graph/generated"
	"tarareba-bff/graph/model"
)

func (r *queryResolver) ContestsByUserID(ctx context.Context, userID *string) ([]*model.Contest, error) {
	contests := make([]*model.Contest, 0, 2)
	contests = append(contests, &model.Contest{
		IsRated:           true,
		Place:             128,
		OldRating:         0,
		NewRating:         10,
		Performance:       1000,
		InnerPerformance:  1,
		ContestScreenName: "atcoder.jp",
		ContestName:       "AtCoder Grand Contest 001",
		ContestNameEn:     "",
		EndTime:           "2020-10-10-1",
		IsParticipated:    true,
	})

	contests = append(contests, &model.Contest{
		IsRated:           true,
		Place:             528,
		OldRating:         110,
		NewRating:         1011,
		Performance:       4000,
		InnerPerformance:  1111,
		ContestScreenName: "atcoder.jp",
		ContestName:       "AtCoder Grand Contest 111",
		ContestNameEn:     "",
		EndTime:           "2000-10-10-1",
		IsParticipated:    false,
	})

	return contests, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
