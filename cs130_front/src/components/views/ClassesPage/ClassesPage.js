import React from 'react';
import Tabs from '../../Tabs/Tabs';
import Dropdown from '../../Dropdown/Dropdown';
import Inviter from '../../Inviter/Inviter';
import '../../../App.css';

function ClassesPage() {
  const items = [{name: "Wow Squad", groupId: 123}, {name: "CA Squad", groupId: 124}, {name: "DM Squad", groupId: 125}]
  const user = {firstName: "Ethan", lastName: "Wow", id: 123}
  //inviter is used by including it as an optional element for each user row in the user list 
  //show if invite to group button is clicked etc, then pass the user info from the row to the inviter
  //and groups information for current user can be gotten from the groups endpoint
  function handleGroupInvite(user, group){
    //send request with user and group info

  }
  return (
    <div className="App">
      <Tabs>
        <div type="Study Buddies"></div>
        <div type="Listings"></div>
      </Tabs>
      <Inviter user={user} items={items} handleGroupInvite={handleGroupInvite}/> 
      <Dropdown options={items} sendSelection={() => {}}/>
    </div>
  );
}

export default ClassesPage;