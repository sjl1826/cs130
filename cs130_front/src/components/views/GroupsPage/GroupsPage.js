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


  const classes2 = [
    { name: "Discrete Mathematics", courseId: 1, groups: [{ name: "DM Squad", groupId: 1 }, { name: "DM Squad II", groupId: 1 }] },
    { name: "Computer Architecture", courseId: 2, groups: [{ name: "CA Squad", groupId: 3 }] }
  ]

  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const groupInformation = [
    { name: "Group name", value: "" },
  ];

  const reqs = [{ name: "Al Squad", id: 223, types: "invitation" }, { name: "Calc Gang", id: 223, types: "invitation" }]

  const [requests, setRequests] = useState(reqs);
  const [mainPanelState, setMainPanel] = useState();
  const [currentGroup, setCurrentGroup] = useState({ name: "DM Squad", groupId: 1 });
  const [classes, setClasses] = useState(classes2);

  // Pass this to ClassList
  function groupClicked(title) {
    //set main content to be for title
    setMainPanel(title);
    setCurrentGroup();
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
    switch (mainPanelState) {
      case 'noGroupSelected':
        return <Text color="black" size="24px" weight="800"> Select a group! </Text>
      default:
        return <GroupItem className="group-with-margin-centered" day="Firday" time="4 PM" groupName={currentGroup.name} groupId={currentGroup.id} />
    }
  }
  //{getCourse(currentGroup.courseId).name}
  function myGroupAdmin() {
    return (
      <div className="panel">
        <div className="column">
          <div className="text-container">
            <Text color="black" size="24px" weight="800">
              Course
            </Text>
          </div>
          <CreateGroup options={groupInformation} />
          <Requests title="Requests" items={requests} handleResponse={handleRequest} />
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

  /*return (
    <div className="App">
      <UserList users={members} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={() => {}}/>
    </div>
  );*/
  return myGroupAdmin();

}

export default GroupsPage;