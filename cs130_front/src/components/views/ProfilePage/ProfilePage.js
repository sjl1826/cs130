import React, {useState} from 'react';
import '../../../App.css';
import Requests from '../../Requests/Requests';

function ProfilePage(props) {
  const invs = [{name: "CA Squad", id: 223, types: "invitation"}, {name: "Alexander Hamilton-Tuff", id: 223, types: "invitation"}]
  const [invitations, setInvitations] = useState(invs);
  const userId = props.match.params.id;

  function getProfile() {
    setInvitations(invs);
  }

  function handleInvitations(invitation){
    //update invitation with accept/decline
    // and remove from list then fetch again to update ui
  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
  return (
    <div className="App">
      <Requests title="Invitations" items={invitations} handleResponse={handleInvitations}/>
    </div>
  );
}

export default ProfilePage;