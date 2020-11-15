import React, { useState, useEffect } from 'react';
import Tabs from '../../Tabs/Tabs';
import Dropdown from '../../Dropdown/Dropdown';
import Inviter from '../../Inviter/Inviter';
import ClassList from '../../CurrentClasses/ClassList';
import '../../../App.css';
import ClassTitle from '../../CurrentClasses/ClassTitle';
import './styles.css';

function ClassesPage(props) {
  const items = [{name: "Wow Squad", groupId: 123}, {name: "CA Squad", groupId: 124}, {name: "DM Squad", groupId: 125}]
  const user = {firstName: "Ethan", lastName: "Wow", id: 123}
  const classes = [{name: "Discrete Mathematics", courseId: 1, groups:[{name: "DM Squad"}]}, {name: "Computer Architecture", courseId: 2}]
  const [mainPanelState, setMainPanel] = useState(classes[0].name);
  //inviter is used by including it as an optional element for each user row in the user list 
  //show if invite to group button is clicked etc, then pass the user info from the row to the inviter
  //and groups information for current user can be gotten from the groups endpoint
  function handleGroupInvite(user, group){
    //send request with user and group info

  }

  function classClicked(item) {
    //set main content to be for title
    setMainPanel(item.name);
  }

  return (
    <div className="App">
        <div style={{display: 'flex', flexDirection: 'row', paddingTop: '20px', justifyContent: "center"}}>
          <ClassTitle option={mainPanelState}/>
        </div>
        <div style={{paddingTop: '0px'}} className="panel">
          <div className="column">
            <Inviter user={user} items={items} handleGroupInvite={handleGroupInvite}/> 
            <Dropdown options={items} sendSelection={() => {}}/>
          </div>
          <div className="column">
            <Tabs>
              <div type="Study Buddies"></div>
              <div type="Listings"></div>
            </Tabs>
          </div>
          <div className="column">
            <ClassList classList={classes} titleClicked={classClicked} clickable={true}/>
          </div>
        </div>
    </div>
    
  );
}

export default ClassesPage;