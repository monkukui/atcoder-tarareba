import React, { useState } from 'react'
import './App.css'

import { ApolloProvider } from 'react-apollo'
import { ApolloProvider as ApolloHooksProvider } from 'react-apollo-hooks'

import { appClient } from './graphql/client'

import TopPage from './components/TopPage'
import HowToUse from './components/HowToUse'
import Contact from './components/Contact'
import Ranking from './components/Ranking'

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
      inverted
      style={{
        fontSize: '3.5em',
        fontWeight: 'normal',
        marginBottom: 0,
        marginTop: '2em',
      }}
    >
      At
      <span style={{ color: 'rgb(229, 115, 110)' }}>C</span>
      oder tarareba
    </Header>
    <Header.Subheader
      as="h2"
      style={{
        fontSize: '1.2em',
        fontWeight: 'normal',
        marginTop: '1.5em',
      }}
    >
      「このコンテストが存在しなかっ
      <span style={{ color: 'rgb(229, 115, 110)', fontWeight: 'bold' }}>
        たら
      </span>
      」
      <br />
      「あのコンテストに参加していなけ
      <span style={{ color: 'rgb(229, 115, 110)', fontWeight: 'bold' }}>
        れば
      </span>
      」
      <br />
      AtCoder tarareba は、過去を改竄してレートを最大化するサービスです
      <br />
    </Header.Subheader>
  </Container>
)

const App = () => {
  const [fixed, setFixed] = useState(false)
  const [page, setPage] = useState('top')
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
                style={{
                  minHeight: page === 'top' ? 550 : 40,
                  padding: '0.2em 0em',
                }}
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
                    <Menu.Item
                      as="a"
                      active={page === 'top'}
                      onClick={() => {
                        setPage('top')
                      }}
                    >
                      Top
                    </Menu.Item>
                    <Menu.Item
                      as="a"
                      active={page === 'howtouse'}
                      onClick={() => {
                        setPage('howtouse')
                      }}
                    >
                      How to use
                    </Menu.Item>
                    <Menu.Item
                      as="a"
                      active={page === 'ranking'}
                      onClick={() => {
                        setPage('ranking')
                      }}
                    >
                      Ranking
                    </Menu.Item>
                    <Menu.Item
                      as="a"
                      active={page === 'contact'}
                      onClick={() => {
                        setPage('contact')
                      }}
                    >
                      Contact
                    </Menu.Item>
                    <Menu.Item position="right">
                      <Button
                        as="a"
                        inverted={!fixed}
                        primary={fixed}
                        style={{ marginLeft: '0.5em' }}
                        href="https://github.com/monkukui/atcoder-tarareba"
                      >
                        GitHub
                      </Button>
                    </Menu.Item>
                  </Container>
                </Menu>
                {page === 'top' ? <HomepageHeading /> : null}
              </Segment>
            </Visibility>
            <Container style={{ width: '70%' }}>
              {page === 'top' ? <TopPage /> : null}
              {page === 'howtouse' ? <HowToUse /> : null}
              {page === 'contact' ? <Contact /> : null}
              {page === 'ranking' ? <Ranking /> : null}
            </Container>
            <Segment
              inverted={true}
              vertical={true}
              style={{ padding: '1em 0em', marginTop: '10em' }}
            >
              <Container textAlign="center">
                <p>AtCoder tarareba 2020</p>
              </Container>
            </Segment>
          </ApolloHooksProvider>
        </ApolloProvider>
      </header>
    </div>
  )
}

export default App
