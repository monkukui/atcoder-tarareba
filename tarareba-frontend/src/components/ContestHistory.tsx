import React, { useState, useEffect, useRef } from 'react'
import {
  GET_CONTEST_HISTORY,
  GET_RATING_TRANSITION,
} from '../graphql/tags/getContestHistory'

import { useQuery } from 'react-apollo-hooks'
import { useLazyQuery } from '@apollo/react-hooks'

import { Table, Card, Button, Checkbox, Icon } from 'semantic-ui-react'

import getRatingColorStyle from '../utils/getRatingColorStyle'

interface Contest {
  rateChange: string
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
  const [contest, setContest] = useState<Contest[]>([])
  const { loading, error, data } = useQuery<Contests>(GET_CONTEST_HISTORY, {
    variables: {
      userID: props.userID,
    },
  })
  const [getRatingTransision, { data: ratingTransision }] = useLazyQuery(
    GET_RATING_TRANSITION,
  )

  const ref = useRef<null | HTMLDivElement>(null)

  // GraphQL からデータが取得できたら、contest にデータをバインディングする
  // 以降は、contest をフロントエンドで管理して、表示を切り替える（AtCoder にアクセスするのは最初に一回だけ）
  useEffect(() => {
    if (!loading && !error && data) {
      setContest(data.contestsByUserID)
      if (ref != null && ref.current != null) {
        ref!.current!.scrollIntoView({ block: 'end', behavior: 'smooth' })
      }
    }
  }, [loading, error, data])

  // GraphQL のクエリが完了して、ratingTransision が変更したきっかけで、contest を update する
  useEffect(() => {
    if (
      ratingTransision &&
      ratingTransision.ratingTransitionByPerformance.length === contest.length
    ) {
      let c: Contest[] = []
      contest.forEach((record, i) => {
        c.push(record)
        c[i].optimalOldRating =
          ratingTransision.ratingTransitionByPerformance[i].oldRating
        c[i].optimalNewRating =
          ratingTransision.ratingTransitionByPerformance[i].newRating
      })
      setContest(c)
    }
  }, [contest, ratingTransision])

  // index 番目の contest の参加履歴を更新する
  // GraphQL サーバーにレート推移計算のクエリを投げる
  // contest 全体を更新する
  const updateContest = (index: number) => {
    let c: Contest[] = []
    let rateChanges: string[] = []
    let isParticipated: boolean[] = []
    let performances: number[] = []
    let innerPerformances: number[] = []

    contest.forEach((record, i) => {
      c.push(record)
      if (i === index) {
        c[c.length - 1].isParticipated = !c[c.length - 1].isParticipated
      }
      rateChanges.push(c[c.length - 1].rateChange)
      isParticipated.push(c[c.length - 1].isParticipated)
      performances.push(c[c.length - 1].performance)
      innerPerformances.push(c[c.length - 1].innerPerformance)
    })
    setContest(c)

    // GraphQL にクエリを投げる
    getRatingTransision({
      variables: {
        rateChanges: rateChanges,
        isParticipated: isParticipated,
        performances: performances,
        innerPerformances: innerPerformances,
      },
    })
  }

  if (loading || !contest) return <div>loading</div>
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

  if (data!.contestsByUserID.length === 0) {
    if (props.userID !== '') {
      return (
        <div>
          <p>ユーザーが見つかりません</p>
        </div>
      )
    }

    return <></>
  }

  const currentRating = data!.contestsByUserID[
    data!.contestsByUserID.length - 1
  ].actualNewRating
  const optimalRating = data!.contestsByUserID[
    data!.contestsByUserID.length - 1
  ].optimalNewRating

  return (
    <>
      <div style={{ marginTop: '3em', fontSize: 'calc(6px + 1vmin)' }}>
        <Card
          style={{
            margin: '0 auto',
            marginBottom: '5em',
            width: '70%',
            outlineColor: 'rgb(67, 121, 178)',
          }}
        >
          <Card.Content header="結果" />
          <Table basic="very" style={{ width: '90%', margin: '0 auto' }}>
            <Table.Body>
              <Table.Row>
                <Table.Cell>ユーザー名</Table.Cell>
                <Table.Cell>{props.userID}</Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell>実際のレート</Table.Cell>
                <Table.Cell style={getRatingColorStyle(currentRating)}>
                  {currentRating}
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell>架空のレート</Table.Cell>
                <Table.Cell style={getRatingColorStyle(optimalRating)}>
                  {optimalRating}
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell>差分</Table.Cell>
                <Table.Cell style={{ fontWeight: 'bold' }}>
                  {optimalRating > currentRating ? (
                    <>+{optimalRating - currentRating}</>
                  ) : (
                    <>-{currentRating - optimalRating}</>
                  )}
                </Table.Cell>
              </Table.Row>
            </Table.Body>
          </Table>
          <Card.Content style={{ marginTop: '2em' }}>
            <Button basic>
              <Icon name="twitter" style={{ color: 'rgb(74, 161, 235)' }} />
              で投稿する
            </Button>
          </Card.Content>
        </Card>
        <Table celled={true} style={{ fontSize: 'calc(6px + 1vmin)' }}>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>日付け</Table.HeaderCell>
              <Table.HeaderCell>コンテスト</Table.HeaderCell>
              <Table.HeaderCell>順位</Table.HeaderCell>
              <Table.HeaderCell>パフォーマンス</Table.HeaderCell>
              <Table.HeaderCell>実際の Rating</Table.HeaderCell>
              <Table.HeaderCell>架空の Rating</Table.HeaderCell>
              <Table.HeaderCell>参加・不参加</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <div ref={ref}></div>

          <Table.Body>
            {contest!.map((record, index) => {
              return (
                <Table.Row
                  key={record.endTime}
                  positive={record.isParticipated && record.rateChange !== '-'}
                  negative={
                    !(record.isParticipated && record.rateChange !== '-')
                  }
                >
                  <Table.Cell>{record.endTime}</Table.Cell>
                  <Table.Cell>
                    <a
                      href={'https://' + record.contestScreenName}
                      rel="noreferrer noreferrer"
                      target="_blank"
                    >
                      {record.contestName}
                    </a>
                  </Table.Cell>
                  <Table.Cell>
                    <a
                      href={
                        'https://' +
                        record.contestScreenName +
                        '/standings?watching=' +
                        props.userID
                      }
                      rel="noreferrer noreferrer"
                      target="_blank"
                    >
                      {record.place}
                    </a>
                  </Table.Cell>
                  {record.rateChange !== '-' ? (
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
                        &nbsp;
                        {record.actualNewRating > record.actualOldRating ? (
                          <>
                            (+{record.actualNewRating - record.actualOldRating})
                          </>
                        ) : (
                          <>
                            (-{record.actualOldRating - record.actualNewRating})
                          </>
                        )}
                      </Table.Cell>
                      {record.isParticipated ? (
                        <>
                          <Table.Cell
                            style={getRatingColorStyle(record.optimalNewRating)}
                          >
                            {record.optimalNewRating}
                            &nbsp;
                            {record.optimalNewRating >
                            record.optimalOldRating ? (
                              <>
                                (+
                                {record.optimalNewRating -
                                  record.optimalOldRating}
                                )
                              </>
                            ) : (
                              <>
                                (-
                                {record.optimalOldRating -
                                  record.optimalNewRating}
                                )
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
                  <Table.Cell>
                    <Checkbox
                      checked={record.isParticipated}
                      onClick={() => {
                        // 1. useLazyQuery で、bff サーバーにクエリを投げる
                        // 2. 今の contest と帰ってきた情報を組み合わせて、新しい contest を作成する
                        // 3. setContest([...]) で contest を更新
                        updateContest(index)
                      }}
                    />
                  </Table.Cell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>

        <Button
          color="instagram"
          onClick={() => {
            if (ref != null && ref.current != null) {
              ref!.current!.scrollIntoView({ block: 'end', behavior: 'smooth' })
            }
          }}
        >
          トップヘ
        </Button>
      </div>
    </>
  )
}

export default ContestHistory
