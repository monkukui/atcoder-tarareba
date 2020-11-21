import React from 'react';
import './App.css';

import { ApolloProvider, gql, ApolloClient, InMemoryCache, useQuery } from '@apollo/client';


// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <div>hello monkukui</div>
//       </header>
//     </div>
//   );
// }

// const query = gql`
//   query {
//     contestsByUserID(userID: "monkukui") {
//       performance
//       isParticipated
//       contestName
//     }
//   }
// `

// React 側で cors の設定をする必要があることがわかった
const uri = 'http://localhost:8080';

const client = new ApolloClient({
  uri: uri,
  cache: new InMemoryCache(),
});

function App() {
  return (
    <ApolloProvider client={client}>
      <Hoge />
    </ApolloProvider>
  );
}

// ----------

// クエリ定義
const GET_CONTEST_HISTORY = gql`
  query {
    contestsByUserID(userID: "monkukui") {
      performance
      isParticipated
      contestName
    }
  }
`

function Hoge() {

  // マウント時にリクエストが走る
  const { loading, error, data } = useQuery(GET_CONTEST_HISTORY, {});

  if (loading) return <div>loading</div>;
  if (error) return <div>error</div>;

  return (
    <div>Hello Hoge</div>
  )
}

export default App;
