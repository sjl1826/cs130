import React, { useState, useEffect } from 'react';
import GroupItem from '../GroupsPage/GroupItem';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import * as Colors from '../../../constants/Colors';
import axios from 'axios';
import { USER_SERVER, GROUP_SERVER, INVITATION_SERVER } from '../../../Config';

export default function ViewGroupPage(props) {
  const config = {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
    }
  }

  const members = [
    { name: "Shirly fang", school: "UCLA", id: 123, discord: "shirly#123", email: "shirly@gmail.com" },
    { name: "Shirly fang", id: 123, discord: "shirly#123", email: "shirly@gmail.com" }
  ]

  const currentGroup2 = { id: 123, name: "DM Squad", courseName: "Calculus", members: members, day: "friday", time: "4:30pm" }

  const [currentGroup, setCurrentGroup] = useState(null);

  useEffect(() => {
    async function initGroup() {
      try {
        const classesResponse = await getClassesInfo();
        const groupResponse = await getGroup();
        handleGroupResponse(groupResponse.data, classesResponse.data.courses);
      } catch (err) {
        // Handle err here. Either ignore the error, or surface the error up to the user somehow.
      }
    }
    initGroup();
  }, []);

  function makeRequest() {
    const groupId = props.match.params.id;
    const myId = localStorage.getItem('userId'); //receiveid
    const myName = localStorage.getItem('userName'); //receivename
    const body = {
      group_name: currentGroup.name,
      group_id: parseInt(groupId),
      receive_id: parseInt(myId),
      receive_name: myName,
      type: true
    }
    axios.post(`${INVITATION_SERVER}/create?u_id=${myId}`, body, config);
  }

  async function getClassesInfo() {
    return axios.get(`${USER_SERVER}/classes-info`);
  }

  function getGroup() {
    const groupId = props.match.params.id; // from url
    return axios.get(`${GROUP_SERVER}?g_id=${groupId}`, config)
  }

  function handleGroupResponse(groupResponse, classes) {
    const name = groupResponse["name"];
    const id = groupResponse["g_id"];
    const meetingTime = groupResponse["meeting_time"];

    // Populate members
    const members = []
    Object.keys(groupResponse["members"]).forEach(function (key2) {
      const memberName = groupResponse["members"][key2]["first_name"] + " " + groupResponse["members"][key2]["last_name"];
      const school = groupResponse["members"][key2]["school_name"];
      const memberId = groupResponse["members"][key2]["u_id"];
      const facebook = groupResponse["members"][key2]["facebook"];
      const discord = groupResponse["members"][key2]["discord"];
      const email = groupResponse["members"][key2]["u_email"];
      members.push({ name: memberName, school: school, id: memberId, facebook: facebook, discord: discord, email: email });
    });

    Object.keys(classes).forEach(function (key) {
      Object.keys(classes[key]).forEach(function (key2) {
        Object.keys(classes[key][key2]).forEach(function (key3) {
          const targetId = groupResponse["course_id"];
          if (classes[key][key2][key3]["id"] == targetId) {
            const courseName = classes[key][key2][key3]["name"];
            setCurrentGroup({ id: id, name: name, courseName: courseName, meetingTime: meetingTime, members: members });
          }
        });
      });
    });
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
        {currentGroup != null ?
          <div className="panel-third">
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
          </div>
          :
          null
        }
        <div className="column">
          <div className="text-container">
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

  if (currentGroup != null) {
    return viewGroup();
  } else {
    return (
      <Text color="black" size="44px" weight="800">
        Error viewing group :/
      </Text>
    );
  }

}
