import React from 'react';
import UserItem from "./UserItem";
import './styles.css'

function UserList(props) {
    return (
        <div className='user-container'>
            {props.users.map((user, index) => (
                <UserItem user={user}
                    goToUserProfile={props.goToUserProfile}
                    adminPrivilege={props.adminPrivilege}
                    optionalElement={props.optionalElement}
                    optionalClick={props.optionalClick} />
            ))}
        </div>
    );
}

export default UserList;
