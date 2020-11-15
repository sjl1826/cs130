import React, {useState} from 'react';
import Text from '../../Text/Text';
import TextInput from '../../TextInput/TextInput';
import Button from '../../Button/Button';
import * as Colors from '../../../constants/Colors';
import './styles.css';

export default function CourseCreater(props) {
  const [courseName, setCourseName] = useState('');
  const [keywords, setCourseKeywords] = useState('');
  return(
    <div className="col">
      <div className="group-with-margin-centered">
        <Text size="30px" weight="800">Class name</Text>
        <TextInput 
        onChange={(event) => {
          setCourseName(event.target.value);
        }}/>
      </div>
      <div className="group-with-margin-centered">
        <Text size="30px" weight="800">Class keywords</Text>
        <div className="text-container">
          <TextInput 
          width="25vw"
          height="10vh"
          onChange={(event) => {
            setCourseKeywords(event.target.value);
          }}/>
        </div>
      </div>
      <div className="group-with-margin-centered">
        <Button 
          textColor={Colors.White}
          textSize="24px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.addCustomCourse(courseName, keywords)}
        >
          Create and Add Course
        </Button>
      </div>

    </div>
  );
}