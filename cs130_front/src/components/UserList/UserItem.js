import React from 'react';
import './styles.css';
import Button from '../Button/Button';

const userItem = (props) => (
    <div className='userBox'>
        <div className='halfUserBox'>
            <p className='topLineName' onClick={props.goToUserProfile(props.user.id)}> {props.user.name} </p>
            <p className='bottomLine'> {props.user.school} </p>
        </div>
        <div className='halfUserBox'>
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
            {props.optionalElement ? <Button height="40px" width="150px" onClick={() => props.optionalClick(props.user)}> Invite to group </Button> : null}
            {props.adminPrivilege ? <Button height="40px" width="150px" onClick={() => props.optionalClick(props.user)}> Remove </Button> : null}
        </div>
    </div>
)

export default userItem