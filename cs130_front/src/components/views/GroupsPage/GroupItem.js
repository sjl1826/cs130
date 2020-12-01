import React from 'react';
import Text from '../../Text/Text';
import UserList from '../../UserList/UserList';
import { withRouter } from 'react-router-dom';
import './styles.css';

function GroupItem(props) {
  function goToUserProfile(user) {
    window.open(`/profile/${user}`, "_blank");
  }
  return (
      <div className="group-container">
          <Text size="44px" weight="800"> {props.group.name} </Text>
          <Text size="24px" weight="600"> {props.group.meetingTime} </Text>
          <UserList users={props.group.members}
              goToUserProfile={goToUserProfile}
              adminPrivilege={props.adminPrivilege}
              optionalElement={false}
              optionalClick={() => { }} />
      </div>
  );
}

export default withRouter(GroupItem)
