import React from 'react';
import { css } from 'emotion';
import { Route } from 'react-router-dom';
import Auth from '../../../hoc/auth';
import RegisterPage from '../RegisterPage/RegisterPage.js';
import studyPic from "../../../studyPic.jpg";
import Button from "../../Button/Button";
import '../../../App.css';
import Text from "../../Text/Text"

function LandingPage() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={studyPic} className="App-logo" alt="logo" />
        <p className={css`
          width: 1680px;
          height: 100px;
          font-family: Poppins;`
        }>
          Broaden your scope, broaden your network, broaden your mind.
        </p>
        <Button height="80px" width="200px">
          <Text size="30px" color="white"> Join </Text>
          <Route exact path="/register" component={Auth(RegisterPage, false)} />
        </Button>
      </header>
    </div>
  );
}

export default LandingPage;