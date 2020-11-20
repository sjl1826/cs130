import React, { useState } from 'react';
import axios from 'axios';
import { withRouter } from 'react-router-dom';
import './styles.css';
import TextInput from '../../TextInput/TextInput';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import './styles.css';
import { USER_SERVER } from '../../../Config';

function RegisterPage(props) {
  const [formErrorMessage, setFormErrorMessage] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [firstName, setFirst] = useState('');
  const [lastName, setLast] = useState('');

  async function createUser() {
    try {
      const dataToSubmit = {
        u_email: email,
        password: password,
        first_name: firstName,
        last_name: lastName
      }
      console.log(dataToSubmit)
      const response = await axios.post(`${USER_SERVER}/register`,  dataToSubmit)
      if (response) {
        console.log(response)
        props.history.push('/login');
      }
    } catch(error) {
      console.log(error);
    }
  }

  return (
    <div className="main-container">
        <div className="form-box">
          <Text>First name</Text>
          <TextInput
            width="400px"
            id="firstName"
            placeholder="Enter your first name"
            type="firstName"
            onChange={(event) => {
              setFirst(event.target.value);
            }}
          />
        </div>
        <div className="form-box">
          <Text>Last name</Text>
          <TextInput
            width="400px"
            id="lastName"
            placeholder="Enter your last name"
            type="lastName"
            onChange={(event) => {
              setLast(event.target.value);
            }}
          />
        </div>
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
          <Button onClick={() => createUser()}>
            Register
          </Button>
        </div>
    </div>
	);
}

export default withRouter(RegisterPage);