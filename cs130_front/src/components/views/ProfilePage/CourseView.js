import React from 'react';
import Text from '../../Text/Text';
import Button from '../../Button/Button'
import * as Colors from '../../../constants/Colors';
import './styles.css';

export default function CourseView(props) {
  return(
    <div className="col">
      <div className="group-with-margin-centered">
        <Text size="30px" weight="800">Class name</Text>
        <Text size="24px">{props.item.name}</Text>
      </div>
      <div className="group-with-margin-centered">
        <Text size="30px" weight="800">Class keywords</Text>
        <div className="text-container">
          <Text size="24px">{props.item.keywords.join(', ')}</Text>
        </div>
      </div>
      <div className="group-with-margin-centered">
        <Button 
          textColor={Colors.White}
          textSize="28px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.removeCourse(props.item)}
        >
          Remove Course
        </Button>
      </div>

    </div>
  );
}