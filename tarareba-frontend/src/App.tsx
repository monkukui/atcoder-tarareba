import React from 'react';
import './App.css';

import { ApolloProvider } from 'react-apollo';
import {
  ApolloProvider as ApolloHooksProvider,
} from 'react-apollo-hooks';

import { appClient } from './graphql/client';

import TopPage from './components/TopPage';

const App = () => {
  return (
    <ApolloProvider client={appClient}>
      <ApolloHooksProvider client={appClient}>
        <TopPage />
      </ApolloHooksProvider>
    </ApolloProvider>
  );
};

export default App;
