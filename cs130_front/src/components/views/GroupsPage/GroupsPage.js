import React, { useState, useEffect } from 'react';
import ClassList from '../../CurrentClasses/ClassList';
import GroupItem from './GroupItem';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import Requests from '../../Requests/Requests';
import '../../../App.css';
import CreateGroup from './CreateGroup';
import axios from 'axios';
import { USER_SERVER_AUTH, GROUP_SERVER } from '../../../Config';

function GroupsPage(props) {
  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const reqs = [{ name: "John Smith", id: 223, type: "request", userId: 1 }, { name: "John Oliver", id: 223, type: "request", userId: 2 }]

  const currentGroup2 = { id: 123, name: "DM Squad", courseName: "Calculus", members: members, day: "friday", time: "4:30pm", requests: reqs }

  const classes2 = [
    { name: "Discrete Mathematics", courseId: 1, groups: [currentGroup2], },
    { name: "Computer Architecture", courseId: 2, groups: [currentGroup2] }
  ]

  const groupInformation = [
    { name: "Group name", value: "" },
  ];

  const [currentGroup, setCurrentGroup] = useState(null);
  const [classes, setClasses] = useState(null);
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
        const groupResponse = await getGroups();
        const classResponse = await getClasses();
        handleClassesAndGroupsResponse(groupResponse.data.group_responses, classResponse.data);
      } catch (err) {
        // Handle err here. Either ignore the error, or surface the error up to the user somehow.
      }
    }
    initGroups();
  }, []);

  function getGroups() {
    return axios.get(`${USER_SERVER_AUTH}/getUserGroups?u_id=${userId}`, config);
    //fetch the endpoints for groups and courses in order to populate current classes section. 
  }

  function getClasses() {
    return axios.get(`${USER_SERVER_AUTH}/getBuddiesListings?u_id=${userId}`, config);
  }

  function handleClassesAndGroupsResponse(groupResponse, classResponse) {
    //console.log(classResponse);
    // Popoulate classes
    const classes = []
    Object.keys(classResponse).forEach(function (key) {
      const name = classResponse[key]["CourseName"];
      classes.push({ name: name, courseId: key, groups: [] });
    });

    // Populate groups
    const groups = []
    Object.keys(groupResponse).forEach(function (key) {
      const name = groupResponse[key]["name"];
      const id = groupResponse[key]["g_id"];
      const meetingTime = groupResponse[key]["meeting_time"];

      // Populate members
      const members = []
      Object.keys(groupResponse[key]["members"]).forEach(function (key2) {
        const memberName = groupResponse[key]["members"][key2]["first_name"] + " " + groupResponse[key]["members"][key2]["last_name"];
        const school = groupResponse[key]["members"][key2]["school_name"];
        const memberId = groupResponse[key]["members"][key2]["u_id"];
        const facebook = groupResponse[key]["members"][key2]["facebook"];
        const discord = groupResponse[key]["members"][key2]["discord"];
        const email = groupResponse[key]["members"][key2]["u_email"];
        members.push({ name: memberName, school: school, id: memberId, facebook: facebook, discord: discord, email: email });
      });

      // Populate requests
      const reqs = []
      Object.keys(groupResponse[key]["invitations"]).forEach(function (key2) {
        const reqName = groupResponse[key]["invitations"][key2]["receive_name"];
        const reqId = groupResponse[key]["invitations"][key2]["receive_id"];
        const reqGroupId = groupResponse[key]["invitations"][key2]["group_id"];
        const reqInviteId = groupResponse[key]["invitations"][key2]["id"];
        reqs.push({ name: reqName, groupId: reqGroupId, inviteId: reqInviteId, id: reqId, type: "request" });
      });

      Object.keys(classes).forEach(function (key2) {
        const targetId = groupResponse[key]["course_id"];
        if (classes[key2]["courseId"] == targetId) {
          const courseName = classes[key2]["name"];
          groups.push({ id: id, name: name, courseName: courseName, meetingTime: meetingTime, members: members, requests: reqs });
          classes[key2]["groups"].push({ id: id, name: name, courseName: courseName, meetingTime: meetingTime, members: members, requests: reqs });
        }
      });
    });
    //these sets are for mocked values now but should be the real values from response
    setClasses(classes); //handle null case if no classes, show a message about adding course in order to add groups
    setCurrentGroup(classes[0].groups[0]); // handle null case if no groups at all. this is just setting to first group of first class.
    //not necessarily handle here but need to handle overall
  }

  function handleRequest(status, request) {
    console.log(status, request);
    const body = {
      u_id: parseInt(request.id),
      invitation_id: request.inviteId,
      status: status ? 'ACCEPT' : 'DECLINE'
    }

    axios.put(`${USER_SERVER_AUTH}/updateInvitation`, body, config).then(response => {
      return axios.all[getGroups(), getClasses()];
    }).then(axios.spread((groupResponse, classResponse) => {
      handleClassesAndGroupsResponse(groupResponse.data.group_responses, classResponse.data);
    }));
  }

  function createGroup(group, course) {
    console.log(group, course);

    const body = {
      admin_id: parseInt(userId),
      name: group.value,
      course_id: parseInt(course.courseId)
    }
    axios.post(`${GROUP_SERVER}/create`, body, config).then(response => {
      return axios.all([getGroups(), getClasses()]);
    }).then(axios.spread((groupResponse, classResponse) => {
      console.log(groupResponse);
      console.log(classResponse);
      handleClassesAndGroupsResponse(groupResponse.data.group_responses, classResponse.data);
    }));
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

  function myGroupAdmin() {
    return (
      <div className="panel">
        {currentGroup != null ?
          <div className="panel-third">
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
          </div>
          : null
        }
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={classes} titleClicked={groupClicked} clickable={true} />
          </div>
          {classes.length > 0 ? <CreateGroup options={groupInformation} createGroup={createGroup} courses={classes} /> : null}
        </div>
      </div >
    );
  }

  if (currentGroup != null && classes != null) {
    return myGroupAdmin();
  } else {
    return (
      <Text color="black" size="44px" weight="800">
        Join some groups!
      </Text>
    );
  }

}

export default GroupsPage;