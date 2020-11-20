import React, { useState } from 'react';
import { withRouter } from 'react-router-dom';
import axios from 'axios';
import './styles.css';
import TextInput from '../../TextInput/TextInput';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import './styles.css';
import { USER_SERVER } from '../../../Config';
import qs from 'qs';

function LoginPage(props) {
  const [formErrorMessage, setFormErrorMessage] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  async function loginUser() {
    try {
      const config = {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      }
      const dataToSubmit = {
        email: email,
        password: password
      }
      const response = await axios.post(`${USER_SERVER}/login`,  qs.stringify(dataToSubmit), config)
      if (response) {
        localStorage.setItem('accessToken', response.data.access_token);
        localStorage.setItem('userId', response.data.ID);
        props.history.push(`/profile/${response.data.ID}`)
      }
    } catch(error) {
      console.log(error);
    }
  }

  return (
    <div className="main-container">
        <div className="form-box">
          <Text>Email</Text>
          <TextInput
            width="400px"
            id="email"
            placeholder="Enter your email"
            onChange={(event) => {
              setEmail(event.target.value);
            }}
          />
        </div>
        <div className="form-box">
          <Text>Password</Text>
          <TextInput
            width="400px"
            id="password"
            placeholder="Enter your password"
            type="password"
            onChange={(event) => {
              setPassword(event.target.value);
            }}
          />
        </div>
        {formErrorMessage && <Text error>{formErrorMessage}</Text>}
        <div className="form-box">
          <Button onClick={() => loginUser()}>
            Login
          </Button>
        </div>
    </div>
	);
}

export default withRouter(LoginPage);