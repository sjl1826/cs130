import React from 'react';
import './styles.css';

const userItem = (props) => (
    <div className='userBox'>
        <div className='userInfo'>
            <p className='topLineName' onClick={props.goToUserProfile(props.user)}> {props.user.name} </p>
            <p className='bottomLine'> {props.user.school} </p>
        </div>
        <div className='userInfo'>
            {props.user.discord
                ? <p className='topLine'> Discord: </p>
                : <p className='topLine'> Email: </p>
            }

            {props.user.discord
                ? <p className='bottomLine'> {props.user.discord} </p>
                : <p className='bottomLine'> {props.user.email} </p>
            }
        </div>
    </div>
)

export default userItem