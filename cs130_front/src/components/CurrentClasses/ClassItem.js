import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import GroupItem from './GroupItem'
import './ClassItem.css';

export default function ClassItem(props) {

  return( props.data.groups ?
    <div >    
      <Text class="item" style={{fontFamily: Fonts.Primary, color: Colors.Black}}>{props.data.name}</Text>
      <div>
      {props.data.groups.map(group => <GroupItem group={group} titleClicked={props.titleClicked} clickable={true}/>)} 
      </div>
    </div>  
    : 
    <div>
      {props.clickable ?     
      <div onClick={() => props.titleClicked(props.data.name)}>
        <Text class="item clickable-item" style={{fontFamily: Fonts.Primary, color: Colors.Blue}}>{props.data.name}</Text>
      </div>
      : 
      <Text size="24px" weight="800">
        {props.data.name}
      </Text>
      }
    </div>
    );
} 