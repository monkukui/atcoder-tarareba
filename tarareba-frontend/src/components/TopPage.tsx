import React, { useState, useRef } from 'react'

import { Form } from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

const TopPage = () => {
  const [inputUserID, setInputUserID] = useState('')
  const [userID, setUserID] = useState('')
  const ref = useRef<null | HTMLDivElement>(null)

  return (
    <>
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
            ref!.current!.scrollIntoView({ block: 'end', behavior: 'smooth' })
          }}
        >
          Go
        </Form.Button>
      </Form>
      <div ref={ref} />
      <ContestHistory userID={userID} />
    </>
  )
}

// TODO: ref を下のコンポーネントにわたして、result に標準を合わせてスクロールするようにする

export default TopPage
