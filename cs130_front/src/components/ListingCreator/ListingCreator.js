import React, { useState } from 'react';
import axios from 'axios';
import { withRouter } from 'react-router-dom';
import TextArea from '../TextInput/TextArea';
import Text from '../Text/Text';
import Button from '../Button/Button';
import { USER_SERVER } from '../../Config';
import './ListingCreator.css'
import Dropdown from '../Dropdown/Dropdown';

function ListingCreator(props) {
  const [message, setMessage] = useState('');
  const[group, setGroup] = useState(props.items[0]);
  const noGroupOption = {name: "None",groupId: 0};
  const dataToSubmit = {
    poster: props.userId,
    course_id: props.course_id,
    text_description: message,
    course_name: props.course_name
  }

  function sendSelection(group){
    setGroup(group);
  }

  function createListing(){
    dataToSubmit.group_id = group.groupId;
    dataToSubmit.group_name = group.name;
    props.createListing(dataToSubmit);
    document.getElementById("message").value = "";
  }

  return (
    <div>
        <div className="small-box">
            <div className="input-box">
                <Text>Add a listing to this class</Text>
                <TextArea
                    width="400px"
                    height="200px"
                    id="message"
                    placeholder=""
                    type="message"
                    onChange={(event) => {
                    setMessage(event.target.value);
                    }}
                />
            </div>
            <div className="input-box">
                <Text>Tag group in listing</Text>
                <Dropdown width="20vw" options={[noGroupOption].concat(props.items)} sendSelection={sendSelection}/>
            </div>
        </div>
        <div className="small-box">
          <Button onClick={() => createListing()}>Post</Button>
        </div>
    </div>
	);
}

export default withRouter(ListingCreator);