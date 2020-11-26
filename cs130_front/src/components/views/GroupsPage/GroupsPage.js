import React, { useState, useEffect } from 'react';
import ClassList from '../../CurrentClasses/ClassList';
import GroupItem from './GroupItem';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import Requests from '../../Requests/Requests';
import * as Colors from '../../../constants/Colors';
import UserList from '../../UserList/UserList';
import '../../../App.css';
import CreateGroup from './CreateGroup';

function GroupsPage(props) {


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

  const reqs = [{ name: "Al Squad", id: 223, types: "invitation" }, { name: "Calc Gang", id: 223, types: "invitation" }]

  const [requests, setRequests] = useState(reqs);
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

  function getCourse(courseId) {
    //get course from backend
  }

  function getMeetingTime() {
    //get time from backend
  }

  function handleRequest(request) {
    //update request with accept/decline
    // and remove from list then fetch again to update ui
  }

  const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }

  function renderMainPanel() {
    switch (currentGroup) {
      case 'noGroupSelected':
        return <Text color="black" size="24px" weight="800"> Select a group! </Text>
      default:
        return <GroupItem className="group-with-margin-centered"
          group={currentGroup} />
    }
  }
  //{getCourse(currentGroup.courseId).name}
  function myGroupAdmin() {
    return (
      <div className="panel">
        <div className="column-left">
          <div className="text-container">
            <Text color="black" size="44px" weight="800">
              {currentGroup.courseName}
            </Text>
          </div>
          <Requests title="Requests" items={requests} handleResponse={handleRequest} />
        </div>
        <div className="column">
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={classes} titleClicked={groupClicked} clickable={true} />
          </div>
          <CreateGroup options={groupInformation} />
        </div>
      </div>
    );
  }

  /*return (
    <div className="App">
      <UserList users={members} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={() => {}}/>
    </div>
  );*/
  return myGroupAdmin();

}

export default GroupsPage;