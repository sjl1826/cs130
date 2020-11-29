import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Tabs from '../../Tabs/Tabs';
import Inviter from '../../Inviter/Inviter';
import ClassList from '../../CurrentClasses/ClassList';
import ListingList from '../../Listings/ListingList';
import ClassTitle from '../../CurrentClasses/ClassTitle';
import SearchBar from '../../Search/SearchBar';
import Text from '../../Text/Text';
import ListingCreator from '../../ListingCreator/ListingCreator'
import UserList from '../../UserList/UserList';
import { USER_SERVER_AUTH, COURSE_SERVER, INVITATION_SERVER } from '../../../Config';

function ClassesPage(props) {
  const [classes, setClasses] = useState([]);
  const [mainTitleState, setMainTitle] = useState("");
  const [mainClassId, setMainClassId] = useState(null);
  const [mainListingsDefault, setMainListingsDefault] = useState([]);
  const [mainListings, setMainListings] = useState([]);
  const [addedListing, setAddedListing] = useState(null);
  const [mainStudyBuddies, setMainStudyBuddies] = useState([]);
  const [allGroups, setAllGroups] = useState({});
  const [mainGroups, setMainGroups] = useState([]);
  const [currentTab, setCurrentTab] = useState("Study Buddies");
  const [input, setInput] = useState('');
  const [invitedUser, setInvitedUser] = useState(null);
  const [successMessage, setSuccessMessage] = useState('');

  const userId = localStorage.getItem('userId');
  const config = {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
    }
  }
  
  useEffect(() => {
		async function initCourses() {
			try {
        const userResponse = await getUser();
        const allGroups = handleUserResponse(userResponse.data.groups);
        const courseResponse = await getCourses();
        handleBuddiesListingsResponse(courseResponse.data, allGroups);
			} catch (err) {
				// Handle err here. Either ignore the error, or surface the error up to the user somehow.
			}
		}
    initCourses();
  }, []);

  useEffect(() => {
    async function initCourses() {
      try {
        classes.forEach(course => {
          if (addedListing.course_id == course.courseId) {
            classClicked(course);
          }
        })
      } catch (err) {
        // Handle err here. Either ignore the error, or surface the error up to the user somehow.
      }
    }
    initCourses();
  }, [addedListing, classes]);

  function getCourses() {
    return axios.get(`${USER_SERVER_AUTH}/getBuddiesListings?u_id=${userId}`, config);
  }

  function getUser() {
    return axios.get(`${USER_SERVER_AUTH}?u_id=${userId}`, config);
  }

  function handleUserResponse(data) {
    const allGroups = {};
    Object.keys(data).forEach(function(key) {
      const name = data[key]["name"];
      const groupId = data[key]["id"];
      const currentGroup = {name: name, groupId: groupId};
      if (!(data[key]["course_name"] in allGroups)){
        allGroups[data[key]["course_name"]] = [];
      }
      allGroups[data[key]["course_name"]].push(currentGroup);
    });
    setAllGroups(allGroups);
    return allGroups;
  }

  function handleBuddiesListingsResponse(data, allGroups){

    const i = [{name: "Wow Squad", groupId: 123}, {name: "CA Squad", groupId: 124}, {name: "DM Squad", groupId: 125}]

    const classes = []
    Object.keys(data).forEach(function(key) {
      const studyBuddies = [];
      const userDict = {};
      Object.keys(data[key]["StudyBuddies"]).forEach(function(key2) {
        const name = data[key]["StudyBuddies"][key2]["first_name"] + " " + data[key]["StudyBuddies"][key2]["last_name"];
        const id = data[key]["StudyBuddies"][key2]["u_id"];
        const school = data[key]["StudyBuddies"][key2]["school_name"];
        const email = data[key]["StudyBuddies"][key2]["u_email"];
        const discord = null;
        if (!(id == parseInt(userId))){
          studyBuddies.push({name: name, school: school, id: id, discord: discord, email: email});
        }
        userDict[id] = {name: name, school: school};
      });

      const courseId = key;
      const name = data[key]["CourseName"];
      const listings = [];
      console.log(data[key]["Listings"]);
      Object.keys(data[key]["Listings"]).forEach(function(key2) {
        const poster = data[key]["Listings"][key2]["poster"];
        const name = userDict[poster].name;
        const school = userDict[data[key]["Listings"][key2]["poster"]].school;
        const description = data[key]["Listings"][key2]["text_description"];
        const group_id = data[key]["Listings"][key2]["group_id"];
        const group_name = data[key]["Listings"][key2]["group_name"];
        listings.push({poster: poster, name: name, school: school, description: description, group_id: group_id, group_name: group_name});
      });
      classes.push({name:name, courseId: courseId, listings: listings, studyBuddies: studyBuddies});
    });
    setClasses(classes);
    setMainTitle(classes[0].name);
    setMainClassId(classes[0].courseId);
    setMainListingsDefault(classes[0].listings);
    setMainListings(classes[0].listings);
    setMainStudyBuddies(classes[0].studyBuddies);
    setMainGroups(allGroups[classes[0].name]);
  }

  function createListing(listing) {
    const body = {
      poster: parseInt(listing.poster),
      course_id: parseInt(listing.course_id),
      text_description: listing.text_description,
      course_name: listing.course_name
    }

    if (listing.group_id != 0){
      body["group_id"] = listing.group_id;
      body["group_name"] = listing.group_name;
    }

    axios.post(`${COURSE_SERVER}/addListing?u_id=${userId}`, body, config).then(response => {
      return axios.all([getCourses(), getUser()]);
    }).then( axios.spread((data1, data2) => {
      const allGroups = handleUserResponse(data2.data.groups);
      handleBuddiesListingsResponse(data1.data, allGroups);
      setAddedListing(listing);
    }));
  }

  function createInvitation(user, group){
    const body = {
      group_name: group.name,
      group_id: group.groupId,
      receive_id: user.id,
      receive_name: user.name,
      type: false
    }
    axios.post(`${INVITATION_SERVER}/create?u_id=${userId}`, body, config).then(response =>{
      setSuccessMessage('Invitation sent!');
      setTimeout(() => {
				setSuccessMessage('');
			}, 3000);
    })
  }

  function classClicked(item) {
    //set main content to be for title
    setMainTitle(item.name);
    setMainClassId(item.courseId);
    setMainListingsDefault(item.listings);
    setMainListings(item.listings);
    setMainGroups(allGroups[item.name]);
    setMainStudyBuddies(item.studyBuddies);
    setInvitedUser(null);
  }

  const updateInput = async (input) => {
    const filtered = mainListingsDefault.filter(listing => {
      return (listing.name.toLowerCase().includes(input.toLowerCase()) || listing.description.toLowerCase().includes(input.toLowerCase()))
    })
    setInput(input);
    setMainListings(filtered);
  }

  function setTabVar(name){
    setCurrentTab(name);
    if (name == "Listings"){
      setMainListings(mainListingsDefault);
    }
    if (name == "Study Buddies"){
      setInvitedUser(null);
    }
  }

  const goToUserProfile = user => () => { props.history.push(`/profile/${user}`); }
  const goToGroup = groupId => () => { props.history.push(`/groups/group/${groupId}`); }

  function optionalClick(user){
    setInvitedUser(user);
  }

  if ( classes.length == 0 ) {
    return (
      <Text color="black" size="44px" weight="800">
      Add some courses to view Study Buddies and Listings!
    </Text>
    )  
  }
  return (
    <div className="panel">
        <div style={{paddingTop: '0px'}} className="panel">
          <div className="column">
            <div className="group-with-margin">
              {currentTab == "Listings" ?
              
              <ListingCreator userId={parseInt(userId)} course_id={mainClassId} course_name={mainTitleState} items={mainGroups} createListing={createListing}/>
              :
              <Inviter user={invitedUser} items={mainGroups} handleGroupInvite={createInvitation}/>   
              }
              {successMessage && <Text color="green">{successMessage}</Text> }
            </div>
          </div>
          <div className="column">
            <ClassTitle option={mainTitleState}/>
            <Tabs setTabVar={setTabVar} >
                <UserList type="Study Buddies" users={mainStudyBuddies} goToUserProfile={goToUserProfile} optionalElement={true} optionalClick={optionalClick} />
                <ListingList type="Listings" listingList={mainListings} goToUserProfile={goToUserProfile} goToGroup={goToGroup}/>
            </Tabs>
          </div>
          <div className="column">
            <ClassList classList={classes} titleClicked={classClicked} clickable={true}/>
            {currentTab == "Listings" ?
              <div style={{paddingTop: "50px", display: "flex", flexDirection: "column", alignItems: "flex-start"}}>
                <Text size="28px" weight="800" style={{display: "flex", flexDirection: "column", justifyContent: "flex-start", alignItems: "flex-start", minWidth: "20vw"}}>Search Listings</Text>
                <SearchBar input={input} onChange={updateInput} width="18rem" fontSize="25px"/> 
              </div>
              : 
              null
            }
          </div>
        </div>
    </div>
    
  );
}

export default ClassesPage;