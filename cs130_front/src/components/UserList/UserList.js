import React from 'react';
import UserItem from "./UserItem";
import './styles.css'

function UserList(props) {
    return (
        <div className='userlist-container'>
            {props.users.map((user, index) => (
                <UserItem user={user} goToUserProfile={props.goToUserProfile} optionalElement={props.optionalElement} optionalClick={props.optionalClick} />
            ))}
        </div>
    );
}

export default UserList;
