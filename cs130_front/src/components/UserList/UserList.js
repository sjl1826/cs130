import React from 'react';
import UserItem from "./UserItem";
import './styles.css'

function UserList(props) {
  console.log(props.users)
    return (
        <div className='userlist-container'>
            {props.users.map((user, index) => (
                <UserItem user={user}
                    key={user.id}
                    goToUserProfile={props.goToUserProfile}
                    adminPrivilege={props.adminPrivilege}
                    optionalElement={props.optionalElement}
                    optionalClick={props.optionalClick} />
            ))}
        </div>
    );
}

export default UserList;
