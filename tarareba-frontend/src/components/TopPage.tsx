import React, { useState } from 'react'

import { Form } from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

const TopPage = () => {
  const [inputUserID, setInputUserID] = useState('')
  const [userID, setUserID] = useState('')

  return (
    <>
      <Form style={{ marginTop: '5em' }}>
        <Form.Input
          fluid={true}
          label={'AtCoder ID'}
          placeholder={'monkukui'}
          value={inputUserID}
          onChange={(e) => setInputUserID(e.target.value)}
        />
        <Form.Button
          color="instagram"
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
