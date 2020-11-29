import gql from 'graphql-tag'

// AtCoder ID を受け取り、コンテスト情報を返す
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
// パフォーマンス列を受け取り、レート推移を返す
export const GET_RATING_TRANSITION = gql`
  query(
    $isParticipated: [Boolean]!
    $performances: [Int]!
    $innerPerformances: [Int]!
  ) {
    ratingTransitionByPerformance(
      isParticipated: $isParticipated
      performances: $performances
      innerPerformances: $innerPerformances
    ) {
      oldRating
      newRating
    }
  }
`
