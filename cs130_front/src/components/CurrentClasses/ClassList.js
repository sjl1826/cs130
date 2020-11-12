import React from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import ClassItem from './ClassItem';
import './ClassList.css';

const ClassList = (props) => {

  return (
    <>
    <div style = {{fontFamily: Fonts.Primary}} className="class-box">
    <Text size="24px" weight="800">Current Classes</Text>
    { props.classList.map((data,index) => {
        if (data) {
          return (
            <div >
              <ClassItem title={data.name} classClicked={props.classClicked} clickable = {true}/>
	          </div>	
    	    );	
    	  }
    	  return null;
    })
    }
    </div>
    </>
  );
}

export default ClassList