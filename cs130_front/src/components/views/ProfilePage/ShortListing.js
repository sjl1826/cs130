import React from 'react';
import Text from '../../Text/Text';
import Button from '../../Button/Button'
import './styles.css';

export default function ShortListing(props) {
  return(
    <div className="shadow-box">
     <div className="col">
      <Text weight="800">
        {props.item.courseName}
      </Text>
      <Text>
        {props.item.content}
      </Text>
     </div>
     <div className="col">
      <Button onClick={() => props.editListing(props.item)}>Close</Button>
     </div>
    </div>
  );
}