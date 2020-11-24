import gql from 'graphql-tag';

// クエリ定義
export const GET_CONTEST_HISTORY = gql`
  query($userID: String!) {
    contestsByUserID(userID: $userID) {
      isRated
      place
      actualOldRating
      actualNewRating
      performance
      innerPerformance
      contestScreenName
      contestName
      contestNameEn
      endTime
      optimalOldRating
      optimalNewRating
      isParticipated
    }
  }
`
