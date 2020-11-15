import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import GroupItem from './GroupItem'
import './ClassItem.css';

export default function ClassItem(props) {

  return( props.data.groups ?
    <div>    
      <Text className="item" style={{fontFamily: Fonts.Primary, color: Colors.Black, fontSize: "24px"}}>{props.data.name}</Text>
      <div>
      {props.data.groups.map(group => <GroupItem group={group} titleClicked={props.titleClicked} clickable={true}/>)} 
      </div>
    </div>  
    : 
    <div className="item" onClick={() => props.titleClicked(props.data)}>
      {props.clickable ?  <Text className="clickable-item" style={{fontFamily: Fonts.Primary, color: Colors.Blue, fontSize: "24px"}}>{props.data.name}</Text> :
        <Text size="24px">
          {props.data.name}
        </Text>
        }
    </div>
    );
} 
