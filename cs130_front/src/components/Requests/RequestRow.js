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
        to={ props.item.type === 'request' ? `/groups/${props.item.id}` : `/profile/${props.item.id}`}
        className={css`
          font-size: 20px;
          width: 180px;
        `}
      >
        {props.item.name}
      </Link>
      <Button height="35px" width="85px" onClick={() => props.handleResponse(true)}>Accept</Button>
      <Button height="35px" width="85px" onClick={() => props.handleResponse(false)}>Decline</Button>
    </div>
  );
}