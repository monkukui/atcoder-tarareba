import React from 'react';
import { GET_CONTEST_HISTORY } from "../graphql/tags/getContestHistory";

import { useQuery } from 'react-apollo-hooks';

interface Contest {
  isRated: boolean;
  place: number;
  actualOldRating: number;
  actualNewRating: number;
  performance: number;
  innerPerformance: number;
  contestScreenName: string;
  contestName: string;
  contestNameEn: string;
  endTime: string;
  optimalOldRating: number;
  optimalNewRating: number;
  isParticipated: boolean;
}

interface Contests {
  contestsByUserID: Contest[];
}

type Props = {
  userID: string
}

const ContestHistory: React.FC<Props> = (props) => {

  // マウント時にリクエストが走る
  const { loading, error, data } = useQuery<Contests>(GET_CONTEST_HISTORY, {
    variables: {
      userID: props.userID
    },
  });

  if (loading) return <div>loading</div>;
  if (error) return <div>error</div>;

  return (
    <>
      <strong>userID = {props.userID}</strong>
      <ul>
        {data!.contestsByUserID.map(contest => (
          <li>
            isRated = {contest.isRated ? (
              <>
                true
              </>
            ) : (
              <>
                false
              </>
            )} 
            : place = {contest.place}
            : actualOldRating = {contest.actualOldRating}
            : actualNewRating = {contest.actualNewRating}
            : performance = {contest.performance}
            : innerPerformance = {contest.innerPerformance}
            : contestScreenName = {contest.contestScreenName}
            : contestName = {contest.contestName}
            : contestNameEn = {contest.contestNameEn}
            : optimalOldRating = {contest.optimalOldRating}
            : optimalNewRating = {contest.optimalNewRating}
            : isParticipated = {contest.isParticipated ? (
              <>
                true
              </>
            ) : (
              <>
                false
              </>
            )} 
          </li>
        ))}
      </ul>
    </>
  );
};

export default ContestHistory;
