import React, { useState } from 'react'
import './App.css'

import { ApolloProvider } from 'react-apollo'
import { ApolloProvider as ApolloHooksProvider } from 'react-apollo-hooks'

import { appClient } from './graphql/client'

import TopPage from './components/TopPage'
import {
  Button,
  Container,
  Header,
  Menu,
  Segment,
  Visibility,
} from 'semantic-ui-react'

const HomepageHeading = () => (
  <Container text>
    <Header
      as="h1"
      content="AtCoder tarareba"
      inverted
      style={{
        fontSize: '3.5em',
        fontWeight: 'normal',
        marginBottom: 0,
        marginTop: '2em',
      }}
    />
    <Header.Subheader
      as="h2"
      style={{
        fontSize: '1.2em',
        fontWeight: 'normal',
        marginTop: '1.5em',
      }}
    >
      「このコンテストがなかっ
      <span style={{ color: 'rgb(235, 50, 35)', fontWeight: 'bold' }}>
        たら
      </span>
      」
      <br />
      「あのコンテストに参加していなけ
      <span style={{ color: 'rgb(235, 50, 35)', fontWeight: 'bold' }}>
        れば
      </span>
      」
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
  </Container>
)

const App = () => {
  const [fixed, setFixed] = useState(false)
  const showFixedMenu = () => {
    setFixed(true)
  }
  const hideFixedMenu = () => {
    setFixed(false)
  }
  return (
    <div className="App">
      <header className="App-header">
        <ApolloProvider client={appClient}>
          <ApolloHooksProvider client={appClient}>
            <Visibility
              once={false}
              onBottomPassed={showFixedMenu}
              onBottomPassedReverse={hideFixedMenu}
            >
              <Segment
                textAlign="center"
                inverted
                style={{ minHeight: 570, padding: '1em 0em' }}
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
            <TopPage />
          </ApolloHooksProvider>
        </ApolloProvider>
      </header>
    </div>
  )
}

export default App
