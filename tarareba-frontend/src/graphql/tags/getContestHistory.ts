import gql from 'graphql-tag';

// クエリ定義
export const GET_CONTEST_HISTORY = gql`
  query {
    contestsByUserID(userID: "monkukui") {
      performance
      isParticipated
      contestName
    }
  }
`
