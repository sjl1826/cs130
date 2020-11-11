import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Colors from '../../constants/Colors';
import './styles.css';

export default function Infos(props) {
  return(
  <div className="indented-line"> 
    <Text size="18px">{props.option.name}</Text>
    <Text size="18px" color="gray">{props.option.value}</Text>
  </div>
  );
}