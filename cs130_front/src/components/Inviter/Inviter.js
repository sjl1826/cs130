import React, {useState} from 'react';
import Dropdown from '../Dropdown/Dropdown';
import Text from '../Text/Text';
import Button from '../Button/Button';
import './styles.css';

export default function Inviter(props) {
  const[group, setGroup] = useState(props.items[0]);

  function sendSelection(group){
    setGroup(group);
  }

  return(
    <div className="small-rounded-box">
      <Text size="24px" weight="800">Invite to group</Text>
      <Dropdown width="20vw" options={props.items} sendSelection={sendSelection}/>
      <Button textWeight="800" onClick={() => props.handleGroupInvite(props.user, group)}>Invite</Button>
    </div>
  );
}