import React, { useState, useEffect } from 'react';
import GroupItem from '../GroupsPage/GroupItem';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import * as Colors from '../../../constants/Colors';

export default function ViewGroupPage(props) {
  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const currentGroup2 = { id: 123, name: "DM Squad", courseName: "Calculus", members: members, day: "friday", time: "4:30pm" }

  const [currentGroup, setCurrentGroup] = useState(currentGroup2);

  function makeRequest() {

  }

  function getGroup() {
    setCurrentGroup(currentGroup2);
  }


  const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }

  function renderMainPanel() {
    switch (currentGroup) {
      case 'noGroupSelected':
        return <Text color="black" size="24px" weight="800"> Select a group! </Text>
      default:
        return <GroupItem className="group-with-margin-centered" group={currentGroup} />
    }
  }
  //{getCourse(currentGroup.courseId).name}
  function viewGroup() {
    return (
      <div className="panel">
        <div className="column">
          <div className="text-container">
            <Text color="black" size="50px" weight="1000">
              {currentGroup.courseName}
            </Text>
          </div>
        </div>
        <div className="column">
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <Button
              textColor={Colors.White}
              textSize="28px"
              width="275px"
              height="45px"
              textWeight="800"
              color={Colors.Blue}
              onClick={() => makeRequest()}
            >
              Join Group
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return viewGroup();

}
