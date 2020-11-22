import React from 'react';
import { Link } from 'react-router-dom';
import * as Colors from '../../constants/Colors';
import { css } from 'emotion';
import Button from '../Button/Button';
import './styles.css';

export default function RequestRow(props) {
  return(
    <div className="horizontal-row">
      <Link 
        to={ props.item.type === 'invitation' ? `/groups/${props.item.groupId}` : `/profile/${props.item.id}`}
        className={css`
          font-size: 20px;
          width: 180px;
        `}
      >
        {props.item.name}
      </Link>
      <Button height="35px" width="85px" onClick={() => props.handleResponse(true, props.item)}>Accept</Button>
      <Button height="35px" width="85px" onClick={() => props.handleResponse(false, props.item)}>Decline</Button>
    </div>
  );
}