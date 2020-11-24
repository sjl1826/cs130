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

  async function createListing() {
    try {
      const dataToSubmit = {
        poster: props.user,
        course_name: props.course_name,
        course_id: 0, // TODO, get courseID from course
        text_description: message,
        group_id: 0, // TODO, get groupID from group
      }
      console.log(dataToSubmit)
      const response = await axios.post(`${USER_SERVER}/classes/addListing`,  dataToSubmit)
      if (response) {
        console.log(response)
      }
    } catch(error) {
      console.log(error);
    }
  }

  function sendSelection(group){
    setGroup(group);
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