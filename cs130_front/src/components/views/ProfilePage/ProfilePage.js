import React, {useState} from 'react';
import './styles.css';
import Requests from '../../Requests/Requests';
import Infos from '../../Infos/Infos';
import SimpleForm from '../../Infos/SimpleForm';

function ProfilePage(props) {
  const invs = [{name: "CA Squad", id: 223, types: "invitation"}, {name: "Alexander Hamilton-Tuff", id: 223, types: "invitation"}]
  const contactInformation = [
    {name: "Email", value: "shirly@gmail.com"}, 
    {name: "Facebook", value: "facebook.com/shirly"},
    {name: "Discord", value: "shirly#1234"}
  ];
  const additionalInformation = [
    {name: "School name", value: "shirly@gmail.com"}, 
    {name: "Timezone", value: "facebook.com/shirly"},
    {name: "Biography", value: "I enjoy graphs and cookies"}
  ];

  const [invitations, setInvitations] = useState(invs);
  const [contactInfo, setContactInfo] = useState(contactInformation);
  const [additionalInfo, setAdditionalInfo] = useState(additionalInformation);
  const userId = props.match.params.id;
  // check clickable by checking user's id with path's id


  function getProfile() {
    setInvitations(invs);
    setContactInfo(contactInformation);
    setAdditionalInfo(additionalInformation);
  }

  function titleClicked(title) {
    //set main content to be for title
  }

  function handleInvitations(invitation){
    //update invitation with accept/decline
    // and remove from list then fetch again to update ui
  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
  return (
    <div className="panel">
      <div className="column"> <Requests title="Invitations" items={invitations} handleResponse={handleInvitations}/></div>
      <div className="column"><SimpleForm options={contactInformation}/></div>
      <div className="column">
        <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={true}/>
        <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={true}/>
      </div>
    </div>
  );
}

export default ProfilePage;