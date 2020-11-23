import React, { useState, useEffect } from 'react';
import Tabs from '../../Tabs/Tabs';
import Dropdown from '../../Dropdown/Dropdown';
import Inviter from '../../Inviter/Inviter';
import ClassList from '../../CurrentClasses/ClassList';
import ListingList from '../../Listings/ListingList';
import '../../../App.css';
import ClassTitle from '../../CurrentClasses/ClassTitle';
import './styles.css';
import StudyBuddyList from '../../StudyBuddies/StudyBuddyList';
import SearchBar from '../../Search/SearchBar';
import Text from '../../Text/Text';

function ClassesPage(props) {
  const items = [{name: "Wow Squad", groupId: 123}, {name: "CA Squad", groupId: 124}, {name: "DM Squad", groupId: 125}]
  const user = {firstName: "Ethan", lastName: "Wow", id: 123}
  const classes = [{name: "Discrete Mathematics", courseId: 1, listings: [{poster: "Edgar Garcia", school: "University of Nevada, Reno", description: "Looking for a group of Nevada kids to talk about bipartite graphs! Hit me up!"}, {poster: "Colin Kaepernick", school: "University of Nevada, Reno", description: "Looking for 2 more Nevada kids!", groupId: 5, groupName: "DM Squad"}]}, {name: "Computer Architecture", courseId: 2, listings: [{poster: "Glenn Reinman", school: "University of California, Los Angeles", description: "Looking for a group of students to workout with!"}]}]
  //, groups:[{name: "DM Squad"}]
  const [mainTitleState, setMainTitle] = useState(classes[0].name);
  const [mainListingsDefault, setMainListingsDefault] = useState(classes[0].listings);
  const [mainListings, setMainListings] = useState(classes[0].listings);
  const [currentTab, setCurrentTab] = useState("studybuddies")
  const [input, setInput] = useState('');
  //inviter is used by including it as an optional element for each user row in the user list 
  //show if invite to group button is clicked etc, then pass the user info from the row to the inviter
  //and groups information for current user can be gotten from the groups endpoint
  function handleGroupInvite(user, group){
    //send request with user and group info

  }

  function classClicked(item) {
    //set main content to be for title
    setMainTitle(item.name);
    setMainListingsDefault(item.listings);
    setMainListings(item.listings);
  }

  const updateInput = async (input) => {
    const filtered = mainListingsDefault.filter(listing => {
      return (listing.poster.toLowerCase().includes(input.toLowerCase()) || listing.description.toLowerCase().includes(input.toLowerCase()))
    })
    setInput(input);
    setMainListings(filtered);
  }

  function setTabVar(name){
    setCurrentTab(name);
    if (name == "studybuddies"){
      setMainListings(mainListingsDefault);
    }
  }

  const goToUserProfile = user => () => { props.history.push(`/profile/${user}`); }
  const goToGroup = groupId => () => { props.history.push(`/groups/group/${groupId}`); }

  return (
    <div className="App">
        <div style={{display: 'flex', flexDirection: 'row', paddingTop: '20px', justifyContent: "center"}}>
          <ClassTitle option={mainTitleState}/>
        </div>
        <div style={{paddingTop: '0px'}} className="panel">
          <div className="column">
            <Inviter user={user} items={items} handleGroupInvite={handleGroupInvite}/> 
            <Dropdown options={items} sendSelection={() => {}}/>
          </div>
          <div className="column">
            <Tabs>
              <div type="Study Buddies">
                <StudyBuddyList setTabVar={setTabVar} />
              </div>
              <div type="Listings">
                <ListingList listingList={mainListings} goToUserProfile={goToUserProfile} goToGroup={goToGroup} setTabVar={setTabVar}/>
              </div>
            </Tabs>
          </div>
          <div className="column">
            <ClassList classList={classes} titleClicked={classClicked} clickable={true}/>
            {currentTab == "listing" ?
              <div style={{paddingTop: "50px"}}>
                <Text size="28px" weight="800" style={{display: "flex", flexDirection: "column", justifyContent: "flex-start", alignItems: "flex-start", minWidth: "20vw"}}>Search Listings</Text>
                <SearchBar input={input} onChange={updateInput} width="18rem" fontSize="25px"/> 
              </div>
              : 
              <div></div>
            }
          </div>
        </div>
    </div>
    
  );
}

export default ClassesPage;