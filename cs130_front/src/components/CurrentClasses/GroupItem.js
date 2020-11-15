import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import './ClassItem.css'

export default function GroupItem(props) {
  return(
  <div className="item clickable-item inner" onClick={() => props.titleClicked(props.group)}> 
    <Text  style={{fontFamily: Fonts.Primary, color: Colors.Blue}}>
      {props.group.name}
    </Text> 
  </div>
  );
} 
