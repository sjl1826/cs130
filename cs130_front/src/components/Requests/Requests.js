import React from 'react';
import Text from '../Text/Text';
import RequestRow from './RequestRow';
import './styles.css';

export default function Requests(props) {
  return(
    <div className="rounded-box">
      <Text size="24px" weight="800">{props.title}</Text>
      {props.items.map(item => <RequestRow item={item} handleResponse={props.handleResponse}/>)}
    </div>
  );
}