import React, { useState } from 'react';
import Text from '../Text/Text';
import Info from './Info';
import './styles.css';

export default function Infos(props) {
  return(
  <div className="info"> 
    {props.clickable ?     
    <div onClick={() => props.titleClicked(props.title)}>
      <Text className="title-clickable">
        {props.title}
      </Text> 
    </div> : 
      <Text size="24px" weight="800">
        {props.title}
      </Text>
    }
    {props.options.map(option => <Info option={option}/>)}
  </div>
  );
}