import React, {useState} from 'react';
import Text from '../../Text/Text';
import Button from '../../Button/Button'
import './styles.css';

export default function ShortListing(props) {
  const [content, setContent] = useState(props.item.content);
  return(
    <div className="shadow-box">
     <div className="col">
      <Text weight="800">
        {props.item.courseName}
      </Text>
      <textarea className="content-box" onChange={event => setContent(event.target.value)} cols="30" rows="5" defaultValue={props.item.content}/>
     </div>
     <div className="col-spaced">
      <Button onClick={() => props.editListing(content, props.item)}>Update</Button>
      <Button onClick={() => props.editListing('Close', props.item)}>Close</Button>
     </div>
    </div>
  );
}
