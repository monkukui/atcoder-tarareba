import React, { useState } from 'react';

import { 
  Button, Form 
} from 'semantic-ui-react'

import ContestHistory from './ContestHistory'

const TopPage = () => {

  const [inputUserID, setInputUserID] = useState("");
  const [userID, setUserID] = useState("");

  return (
    <>
      <h1>AtCoder tarareba</h1>

      <Form>
        <Form.Input
          fluid={true}
          label="AtCoder ID"
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

      <ContestHistory 
        userID={userID}
      />
    </>
  );
};

export default TopPage;
