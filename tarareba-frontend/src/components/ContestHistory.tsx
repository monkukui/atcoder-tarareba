import React from 'react';
import { GET_CONTEST_HISTORY } from "../graphql/tags/getContestHistory";

import { useQuery } from 'react-apollo-hooks';

interface Contest {
  performance: number;
  isParticipated: boolean;
  contestName: string;
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
          <li>{contest.contestName} : {contest.performance} : {contest.isParticipated}</li>
        ))}
      </ul>
    </>
  );
};

export default ContestHistory;
