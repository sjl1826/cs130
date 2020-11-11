import React, { useState } from 'react';
import FormInput from './FormInput';
import Button from '../Button/Button';
import * as Colors from '../../constants/Colors';
import './styles.css';

export default function SimpleForm(props) {
  function sendInput(name, value) {

  }
  return(
    <div className="form">
      {props.options.map(option => <FormInput option={option} sendInput={sendInput}/>)}
      <Button textColor={Colors.White} color={Colors.Blue}>Save</Button>
    </div>
  );
}