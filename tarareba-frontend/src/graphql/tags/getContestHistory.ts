import gql from 'graphql-tag';

// クエリ定義
export const GET_CONTEST_HISTORY = gql`
  query {
    contestsByUserID(userID: "monkukui") {
      isRated
      place
      oldRating
      newRating
      performance
      innerPerformance
      contestScreenName
      contestName
      contestNameEn
      endTime
      isParticipated
    }
  }
`
