import React, { useState, useEffect } from 'react';
import ClassList from '../../CurrentClasses/ClassList';
import GroupItem from '../GroupsPage/GroupItem';
import Text from '../../Text/Text';
import '../../../App.css';
import CreateGroup from '../GroupsPage/CreateGroup';

export default function ViewGroupPage(props) {
  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const currentGroup2 = { id: 123, name: "DM Squad", courseName: "Calculus", members: members, day: "friday", time: "4:30pm" }

  const classes2 = [
    { name: "Discrete Mathematics", courseId: 1, groups: [currentGroup2], },
    { name: "Computer Architecture", courseId: 2, groups: [currentGroup2] }
  ]

  const groupInformation = [
    { name: "Group name", value: "" },
  ];

  const [currentGroup, setCurrentGroup] = useState(classes2[0].groups[0]);
  const [classes, setClasses] = useState(classes2);

  // Pass this to ClassList
  function groupClicked(group) {
    //set main content to be for title
    setCurrentGroup(group);
  }

  function getClassesList() {
    setClasses(classes2);
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
            <Text color="black" size="24px" weight="800">
              {currentGroup.courseName}
            </Text>
          </div>
          <CreateGroup options={groupInformation} />
        </div>
        <div className="column">
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={classes} titleClicked={groupClicked} clickable={true} />
          </div>
        </div>
      </div>
    );
  }

  return viewGroup();

}
