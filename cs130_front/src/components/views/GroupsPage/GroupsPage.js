import React, { useState, useEffect } from 'react';
import ClassList from '../../CurrentClasses/ClassList';
import GroupItem from './GroupItem';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import Requests from '../../Requests/Requests';
import '../../../App.css';
import CreateGroup from './CreateGroup';
import axios from 'axios';
import { USER_SERVER_AUTH } from '../../../Config';

function GroupsPage(props) {
  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const reqs = [{ name: "John Smith", id: 223, type: "request", userId: 1 }, { name: "John Oliver", id: 223, type: "request", userId: 2 }]

  const currentGroup2 = { id: 123, name: "DM Squad", courseName: "Calculus", members: members, day: "friday", time: "4:30pm" , requests: reqs}

  const classes2 = [
    { name: "Discrete Mathematics", courseId: 1, groups: [currentGroup2], },
    { name: "Computer Architecture", courseId: 2, groups: [currentGroup2] }
  ]

  const groupInformation = [
    { name: "Group name", value: "" },
  ];

  const [currentGroup, setCurrentGroup] = useState(classes2[0].groups[0]);
  const [classes, setClasses] = useState(classes2);
  const userId = localStorage.getItem('userId');
  const config = {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
    }
  }

  // Pass this to ClassList
  function groupClicked(group) {
    //set main content to be for title
    setCurrentGroup(group);
  }

  useEffect(() => {
    async function initGroups() {
      try {
        const response = await getClassesAndGroups();
        console.log(response);
        handleClassesAndGroupsResponse(response);
      } catch (err) {
        // Handle err here. Either ignore the error, or surface the error up to the user somehow.
      }
    }
    initGroups();
  }, []);

  function getClassesAndGroups() {
    return axios.get(`${USER_SERVER_AUTH}/getUserGroups?u_id=${userId}`, config);
    //fetch the endpoints for groups and courses in order to populate current classes section. 
  }

  function handleClassesAndGroupsResponse(response){
    //these sets are for mocked values now but should be the real values from response
    setClasses(classes2); //handle null case if no classes, show a message about adding course in order to add groups
    setCurrentGroup(classes2[0].groups[0]); // handle null case if no groups at all. this is just setting to first group of first class.
    //not necessarily handle here but need to handle overall
  }

  function handleRequest(status, request) {
    console.log(status, request);
    //update request with accept/decline
    //fetch again to update ui
  }

  function createGroup(group, course) {
    console.log(group, course);
    //create a group for this course and fetch classes and groups again to rerender 
    //current classes to show new group.
  }

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
          {currentGroup.requests.length > 0 ? <Requests title="Requests" items={currentGroup.requests} handleResponse={handleRequest} /> : null}
        </div>
        <div className="column">
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={classes} titleClicked={groupClicked} clickable={true} />
          </div>
            {classes.length > 0 ? <CreateGroup options={groupInformation} createGroup={createGroup} courses={classes}/> : null }
        </div>
      </div>
    );
  }

  return myGroupAdmin();

}

export default GroupsPage;