import React, { useState } from 'react';
import Text from '../Text/Text';
import TextInput from '../TextInput/TextInput';
import './styles.css';

export default function FormInput(props) {
  return(
    <div className="form-input">
      <Text> {props.option.name} </Text>
      <TextInput onChange={event => props.sendInput(props.option.name, event.target.value)}/>
    </div>
  );
}