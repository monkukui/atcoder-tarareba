import React, { useState } from 'react'

import { Form, Header, List } from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

const TopPage = () => {
  const [inputUserID, setInputUserID] = useState('')
  const [userID, setUserID] = useState('')

  return (
    <>
      <Header as="h1">
        <Header.Content>AtCoder tarareba</Header.Content>
        <Header.Subheader>
          「このコンテストがなかっ<span style={{ color: 'red' }}>たら</span>」
          <br />
          「あのコンテストに参加していなけ
          <span style={{ color: 'red' }}>れば</span>」
          <br />
          AtCoder tarareba は、過去を改竄してレートを最大化するサービスです
          <br />
        </Header.Subheader>
      </Header>
      <Header as="h2">How to use</Header>
      <List ordered={true}>
        <List.Item>AtCoder のユーザー名を入力してください</List.Item>
        <List.Item>架空のレーティングをご覧ください</List.Item>
        <List.Item>
          Twitter
          で自慢しましょう（例：あのコンテストがなかったら、俺、今頃黄色コーダーなんだが？）
        </List.Item>
      </List>
      <Form style={{ marginTop: '10em' }}>
        <Form.Input
          fluid={true}
          label={'AtCoder ID'}
          placeholder={'monkukui'}
          value={inputUserID}
          onChange={(e) => setInputUserID(e.target.value)}
        />
        <Form.Button
          color="green"
          onClick={() => {
            setUserID(inputUserID)
          }}
        >
          Go
        </Form.Button>
      </Form>

      <ContestHistory userID={userID} />
    </>
  )
}

export default TopPage
