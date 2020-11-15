import React from 'react';
import UserItem from "./UserItem";
import './styles.css'

function UserList(props) {
    return (
        <div className='container'>
            {props.users.map((user, index) => (
                <UserItem user={user} />
            ))}
        </div>
    );
}

export default UserList;
