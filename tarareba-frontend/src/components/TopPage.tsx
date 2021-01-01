import React, { useState } from 'react'

import {
  Form,
  Header,
  Container,
  Button,
  Visibility,
  Segment,
  Menu,
} from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

// TODO: コンポーネントを分ける
const HomepageHeading = () => (
  <Container text>
    <Header
      as="h1"
      content="AtCoder tarareba"
      inverted
      style={{
        fontSize: '4em',
        fontWeight: 'normal',
        marginBottom: 0,
        marginTop: '2em',
      }}
    />
    <Header.Subheader
      as="h3"
      style={{
        fontSize: '1.2em',
        fontWeight: 'normal',
        marginTop: '1.5em',
      }}
    >
      「このコンテストがなかっ<span style={{ color: 'red' }}>たら</span>」
      <br />
      「あのコンテストに参加していなけ
      <span style={{ color: 'red' }}>れば</span>」
      <br />
      AtCoder tarareba は、過去を改竄してレートを最大化するサービスです
      <br />
    </Header.Subheader>
    <Header
      as="h2"
      content=""
      inverted
      style={{
        fontSize: '1.7em',
        fontWeight: 'normal',
        marginTop: '1.5em',
      }}
    />
    <Button primary size="huge">
      早速試してみる
    </Button>
    <Button secondary size="huge">
      使い方
    </Button>
  </Container>
)

const TopPage = () => {
  const [fixed, setFixed] = useState(false)
  const showFixedMenu = () => {
    setFixed(true)
  }
  const hideFixedMenu = () => {
    setFixed(false)
  }
  const [inputUserID, setInputUserID] = useState('')
  const [userID, setUserID] = useState('')

  return (
    <>
      <Visibility
        once={false}
        onBottomPassed={showFixedMenu}
        onBottomPassedReverse={hideFixedMenu}
      >
        <Segment
          inverted
          textAlign="center"
          style={{ minHeight: 600, padding: '1em 0em' }}
          vertical
        >
          <Menu
            fixed={fixed ? 'top' : undefined}
            inverted={!fixed}
            pointing={!fixed}
            secondary={!fixed}
            size="large"
          >
            <Container>
              <Menu.Item as="a" active>
                Top
              </Menu.Item>
              <Menu.Item as="a">How to use</Menu.Item>
              <Menu.Item as="a">Contact</Menu.Item>
              <Menu.Item position="right">
                <Button
                  as="a"
                  inverted={!fixed}
                  primary={fixed}
                  style={{ marginLeft: '0.5em' }}
                >
                  GitHub
                </Button>
              </Menu.Item>
            </Container>
          </Menu>
          <HomepageHeading />
        </Segment>
      </Visibility>
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
