import React, { useState } from 'react';
import FormInput from './FormInput';
import Button from '../Button/Button';
import * as Colors from '../../constants/Colors';
import './styles.css';

export default function SimpleForm(props) {
  const [firstInput, setFirstInput] = useState({name: props.options[0].name, value: props.options[0].value});
  const [secondInput, setSecondInput] = useState({name: props.options[1].name, value: props.options[1].value});
  const [thirdInput, setThirdInput] = useState({name: props.options[2].name, value: props.options[2].value});
  function sendInput(name, value) {
    switch(name) {
      case firstInput.name:
        setFirstInput({name: firstInput.name, value: value})
        break;
      case secondInput.name:
        setSecondInput({name: secondInput.name, value: value})
        break;
      case thirdInput.name:
        setThirdInput({name: thirdInput.name, value: value})
        break;
      default: 
        break;
    }
  }
  return(
    <div className="form">
      {props.options.map(option => <FormInput key={option.name} option={option} sendInput={sendInput}/>)}
      <Button 
        textColor={Colors.White}
        textSize="28px"
        width="150px"
        height="50px"
        textWeight="800" 
        color={Colors.Blue}
        onClick={() => props.saveInfoClicked(firstInput, secondInput, thirdInput)}
      >
        Save
      </Button>
    </div>
  );
}