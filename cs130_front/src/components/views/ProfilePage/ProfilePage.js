import React, {useState} from 'react';
import './styles.css';
import Requests from '../../Requests/Requests';
import Infos from '../../Infos/Infos';
import SimpleForm from '../../Infos/SimpleForm';
import MyListings from './MyListings';
import SchedulerPage from '../SchedulerPage/SchedulerPage';
import * as Colors from '../../../constants/Colors';
import Text from '../../Text/Text';
import Button from '../../Button/Button';

function ProfilePage(props) {
  const invs = [{name: "CA Squad", id: 223, types: "invitation"}, {name: "Alexander Hamilton-Tuff", id: 223, types: "invitation"}]
  const contactInformation = [
    {name: "Email", value: "shirly@gmail.com"}, 
    {name: "Facebook", value: "facebook.com/shirly"},
    {name: "Discord", value: "shirly#1234"}
  ];
  const additionalInformation = [
    {name: "Name", value: "Edgar Garcia"},
    {name: "School name", value: "UCLA"}, 
    {name: "Timezone", value: "PST"},
    {name: "Biography", value: "I enjoy graphs and cookies"},
  ];
  var availability1 = [1,1,1,1,1,1];
  var availability2 = new Array(330).fill(0);
  var fullAvailability = availability1.concat(availability2);
  const l = [
    {id:123, courseName: "Discrete Mathematics", 
    content: "HMU PLS I really want to talk about bipartite graphs with people from all over the country. I am from Wisconsin. I am very cool."},
    {id:124, courseName: "Human Anatomy", content: "HMU PLS i love anatomy"}
  ];

  const [invitations, setInvitations] = useState(invs);
  const [contactInfo, setContactInfo] = useState(contactInformation);
  const [additionalInfo, setAdditionalInfo] = useState(additionalInformation);
  const [listings, setListings] = useState(l);
  const [mainPanelState, setMainPanel] = useState('Classes');
  const userId = props.match.params.id;
  const myId = 2;
  // check clickable by checking user's id with path's id


  function getProfile() {
    setInvitations(invs);
    setContactInfo(contactInfo);
    setAdditionalInfo(additionalInfo);
    setListings(l);
  }

  function titleClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }

  function saveInfoClicked(first, second, third) {
    //post to endpoint updating info
  }

  function handleInvitation(invitation){
    //update invitation with accept/decline
    // and remove from list then fetch again to update ui
  }

  function editListing(listing){
    //edit listing endpoint
  }

  function viewOnlyProfile() {
    return (
      <div className="view-panel">
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="32px" weight="800"> 
            {additionalInfo[0].value} Profile
            </Text>
            <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={false}/>
            <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={false}/>
            <Text color="black" size="24px" weight="800"> 
              Current Classes
            </Text>
          </div>
        </div>
        <div className="column">
          <SchedulerPage passedSelections={fullAvailability}/>
        </div>
      </div>
    );

  }

  function myProfile() {
    return (
      <div className="panel">
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="24px" weight="800"> 
              Hi Student, edit your class and other information to start studying with others!
            </Text>
          </div>
          <MyListings items={listings} editListing={editListing}/>
          <Requests title="Invitations" items={invitations} handleResponse={handleInvitation}/>
        </div>
        <div className="column"> 
          <SimpleForm options={mainPanelState == 'Contact Information' ? contactInfo : additionalInfo} saveInfoClicked={saveInfoClicked}/>
        </div>
        <div className="column">
          <div className="group-with-margin">
            <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={true}/>
            <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={true}/>
          </div>
          <Button 
          textColor={Colors.White}
          textSize="28px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.history.push({
            pathname: `/profile/${userId}/scheduler`,
            state: { availability: fullAvailability }
            })}
          >
            Set Availability
          </Button>
        </div>
      </div>
    );
  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
  return userId == myId ? myProfile() : viewOnlyProfile();
}

export default ProfilePage;