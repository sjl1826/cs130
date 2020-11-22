import React from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import ListingItem from './ListingItem';

const ListingList = (props) => {

  return (
    <>
    <div style = {{fontFamily: Fonts.Primary}}>
    { props.listingList.map((data,index) => {
        if (data) {
          return (
            <div >
              <ListingItem data={data} />
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

export default ListingList