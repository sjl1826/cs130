import React from 'react';
import './styles.css';
import Button from '../Button/Button';

const userItem = (props) => (
    <div className='userBox'>
        <div>
            <p className='topLineName' onClick={props.goToUserProfile(props.user)}> {props.user.name} </p>
            <p className='bottomLine'> {props.user.school} </p>
        </div>
        <div>
            {props.user.discord
                ? <p className='topLine'> Discord: </p>
                : <p className='topLine'> Email: </p>
            }

            {props.user.discord
                ? <p className='bottomLine'> {props.user.discord} </p>
                : <p className='bottomLine'> {props.user.email} </p>
            }
        </div>
        <div className="col-centered">
            {props.optionalElement ? <Button onClick={() => props.optionalClick(props.user)}> Invite to group </Button> : null}
            {props.adminPrivilege ? <Button onClick={() => props.optionalClick(props.user)}> Remove </Button> : null}
        </div>
    </div>
)

export default userItem