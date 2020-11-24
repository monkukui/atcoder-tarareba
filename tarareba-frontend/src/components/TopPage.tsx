import React, { useState } from 'react'

import { Form } from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

const TopPage = () => {
  const [inputUserID, setInputUserID] = useState('')
  const [userID, setUserID] = useState('')

  return (
    <>
      <h1>AtCoder tarareba</h1>
      <h1>
        「このコンテストがなかったら、あのコンテストに参加していなければ」
      </h1>
      <h1>
        AtCoder tarareba は、過去を改竄してレートを最大化するサービスです。
      </h1>

      <Form style={{ margin: '10em' }}>
        <Form.Input
          fluid={true}
          placeholder={'AtCoder ID'}
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
