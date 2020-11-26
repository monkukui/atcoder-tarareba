import React from 'react'
import { GET_CONTEST_HISTORY } from '../graphql/tags/getContestHistory'

import { useQuery } from 'react-apollo-hooks'

import { Table } from 'semantic-ui-react'

import getRatingColorStyle from '../utils/getRatingColorStyle'

interface Contest {
  isRated: boolean
  place: number
  actualOldRating: number
  actualNewRating: number
  performance: number
  innerPerformance: number
  contestScreenName: string
  contestName: string
  contestNameEn: string
  endTime: string
  optimalOldRating: number
  optimalNewRating: number
  isParticipated: boolean
}

interface Contests {
  contestsByUserID: Contest[]
}

type Props = {
  userID: string
}

const ContestHistory: React.FC<Props> = (props) => {
  // マウント時にリクエストが走る
  const { loading, error, data } = useQuery<Contests>(GET_CONTEST_HISTORY, {
    variables: {
      userID: props.userID,
    },
  })

  if (loading) return <div>loading</div>
  if (error) {
    return (
      <div>
        <p>なんらかのエラーが発生しました。</p>
        <p>
          <a href="https://twitter.com/monkukui2">こちら</a>にご連絡ください
        </p>
      </div>
    )
  }
  return (
    <>
      <div style={{ marginTop: '5em' }}>
        <Table celled={true} style={{ fontSize: 'calc(6px + 1vmin)' }}>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>日付け</Table.HeaderCell>
              <Table.HeaderCell>コンテスト</Table.HeaderCell>
              <Table.HeaderCell>順位</Table.HeaderCell>
              <Table.HeaderCell>パフォーマンス</Table.HeaderCell>
              <Table.HeaderCell>実際の Rating</Table.HeaderCell>
              <Table.HeaderCell>架空の Rating</Table.HeaderCell>
            </Table.Row>
          </Table.Header>

          <Table.Body>
            {data!.contestsByUserID.map((record) => {
              return (
                <Table.Row key={record.endTime}>
                  <Table.Cell>{record.endTime}</Table.Cell>
                  <Table.Cell>
                    <a
                      target="_blank"
                      href={'https://' + record.contestScreenName}
                    >
                      {record.contestName}
                    </a>
                  </Table.Cell>
                  <Table.Cell>{record.place}</Table.Cell>
                  {record.isRated ? (
                    <>
                      <Table.Cell
                        style={getRatingColorStyle(record.performance)}
                      >
                        {record.performance}
                      </Table.Cell>
                      <Table.Cell
                        style={getRatingColorStyle(record.actualNewRating)}
                      >
                        {record.actualNewRating}
                        {record.actualNewRating > record.actualOldRating ? (
                          <>
                            （+{record.actualNewRating - record.actualOldRating}
                            ）
                          </>
                        ) : (
                          <>
                            （-{record.actualOldRating - record.actualNewRating}
                            ）
                          </>
                        )}
                      </Table.Cell>
                      {record.isParticipated ? (
                        <>
                          <Table.Cell
                            style={getRatingColorStyle(record.optimalNewRating)}
                          >
                            {record.optimalNewRating}
                            {record.optimalNewRating >
                            record.optimalOldRating ? (
                              <>
                                （+
                                {record.optimalNewRating -
                                  record.optimalOldRating}
                                ）
                              </>
                            ) : (
                              <>
                                （-
                                {record.optimalOldRating -
                                  record.optimalNewRating}
                                ）
                              </>
                            )}
                          </Table.Cell>
                        </>
                      ) : (
                        <Table.Cell>-</Table.Cell>
                      )}
                    </>
                  ) : (
                    <>
                      <Table.Cell>-</Table.Cell>
                      <Table.Cell>-</Table.Cell>
                      <Table.Cell>-</Table.Cell>
                    </>
                  )}
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>
      </div>
    </>
  )
}

export default ContestHistory
