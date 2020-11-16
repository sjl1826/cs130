import React, { useState, useEffect } from 'react';
import UserList from '../../UserList/UserList';
import '../../../App.css';

function GroupsPage(props) {


  const classes = [
    {name: "Discrete Mathematics", courseId: 1, groups: [{name: "DM Squad", groupId: 1}, {name: "DM Squad II", groupId: 1}] }, 
    {name: "Computer Architecture", courseId: 2, groups: [{name: "CA Squad", groupId: 3}]}
  ]
  const [mainPanelState, setMainPanel] = useState();

  const members= [
    {name: "Shirly fang", school: "UCLA", id:123, discord:"shirly#123", email:"shirly@gmail.com"},
    {name: "Shirly fang", id:123, discord:"shirly#123", email:"shirly@gmail.com"}
  ]
  // Pass this to ClassList
  function groupClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }


  const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }


  return (
    <div className="App">
      <UserList users={members} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={() => {}}/>
    </div>
  );
}

export default GroupsPage;