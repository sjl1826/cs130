import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import './ClassItem.css'

export default function GroupItem(props) {
  return(
  <div > 
    {props.clickable ?     
    <div onClick={() => props.titleClicked(props.group.name)}>
      <Text className="item clickable-item" style={{fontFamily: Fonts.Primary, color: Colors.Blue}}>
      {props.group.name}
      </Text> 
    </div> : 
      <Text size="24px" weight="800">
        {props.group.name}
      </Text>
    }
  </div>
  );
} 