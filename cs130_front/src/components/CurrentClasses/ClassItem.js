import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import './ClassItem.css';

export default function ClassItem(props) {
  return(
  <div > 
    {props.clickable ?     
    <div onClick={() => props.classClicked(props.title)}>
      <Text class="item" style={{fontFamily: Fonts.Primary, color: Colors.Blue}}>
        {props.title}
      </Text> 
    </div> : 
      <Text size="24px" weight="800">
        {props.title}
      </Text>
    }
  </div>
  );
} 