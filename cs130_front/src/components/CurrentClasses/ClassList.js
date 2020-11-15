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
    <Text size="28px" weight="800">Current Classes</Text>
    { props.classList.map((data,index) => {
        if (data) {
          return (
            <div >
              <ClassItem data={data} titleClicked={props.titleClicked} clickable={props.clickable}/>
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
