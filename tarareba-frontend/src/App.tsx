import React from 'react';
import './App.css';

import { ApolloProvider } from 'react-apollo';
import {
  ApolloProvider as ApolloHooksProvider,
} from 'react-apollo-hooks';

import { appClient } from './graphql/client';
import { GET_CONTEST_HISTORY } from './graphql/tags/getContestHistory';

import ContestHistory from './components/ContestHistory';

const App = () => {
  return (
    <ApolloProvider client={appClient}>
      <ApolloHooksProvider client={appClient}>
        <ContestHistory />
      </ApolloHooksProvider>
    </ApolloProvider>
  );
};

export default App;
