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
            <Text size="24px" weight="600"> Group meeting time is {props.group.day}, {props.group.time}. </Text>
            <Text size="30px" weight="800"> {props.group.name} </Text>
            <UserList users={props.group.members} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={() => { }} />
        </div>
    );
}

export default withRouter(GroupItem)