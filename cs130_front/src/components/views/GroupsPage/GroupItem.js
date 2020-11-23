import React, { useState } from 'react';
import Text from '../../Text/Text';
import TextInput from '../../TextInput/TextInput';
import Button from '../../Button/Button';
import UserList from '../../UserList/UserList';
import * as Colors from '../../../constants/Colors';
import './styles.css';

const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
]

export default function GroupItem(props) {
    const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }
    const [courseName, setCourseName] = useState('');
    const [keywords, setCourseKeywords] = useState('');
    return (
        <div className="form">
            <div className="form-input">
                <Text size="30px" weight="800"> Group meeting time is {props.day}, {props.time}. </Text>
            </div>
            <div className="form-input">
                <Text size="30px" weight="800"> {props.groupName} </Text>
                {props.optionalElement ? <Button onClick={() => props.optionalClick(props.groupId)}> Delete Group </Button> : null}
            </div>
            <div className="form-input">
                <UserList users={members} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={() => { }} />
            </div>
        </div>
    );
}