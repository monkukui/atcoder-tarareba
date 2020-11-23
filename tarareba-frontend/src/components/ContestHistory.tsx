import React from 'react';
import { GET_CONTEST_HISTORY } from "../graphql/tags/getContestHistory";

import { useQuery } from 'react-apollo-hooks';

interface Contest {
  isRated: boolean;
  place: number;
  oldRating: number;
  newRating: number;
  performance: number;
  innerPerformance: number;
  contestScreenName: string;
  contestName: string;
  contestNameEn: string;
  endTime: string;
  isParticipated: boolean;
}

interface Contests {
  contestsByUserID: Contest[];
}

const ContestHistory = () => {
  // マウント時にリクエストが走る
  const { loading, error, data } = useQuery<Contests>(GET_CONTEST_HISTORY, {});

  if (loading) return <div>loading</div>;
  if (error) return <div>error</div>;

  return (
    <>
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
            : oldRating = {contest.oldRating}
            : newRating = {contest.newRating}
            : performance = {contest.performance}
            : innerPerformance = {contest.innerPerformance}
            : contestScreenName = {contest.contestScreenName}
            : contestName = {contest.contestName}
            : contestNameEn = {contest.contestNameEn}
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
