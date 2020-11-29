import React, { useState } from 'react';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import UserList from '../../UserList/UserList';
import { withRouter } from 'react-router-dom';
import './styles.css';

function GroupItem(props) {
    const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }
    const [courseName, setCourseName] = useState('');
    const [keywords, setCourseKeywords] = useState('');
    return (
        <div className="group-container">
            <Text size="44px" weight="800"> {props.group.name} </Text>
            <Text size="24px" weight="600"> {props.group.meetingTime}</Text>
            <UserList users={props.group.members}
                goToUserProfile={goToUserProfile}
                adminPrivilege={props.adminPrivilege}
                optionalElement={false}
                optionalClick={() => { }} />
        </div>
    );
}

export default withRouter(GroupItem)