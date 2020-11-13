import React from 'react';
import Text from '../../Text/Text';
import ShortListing from './ShortListing';
import './styles.css';

export default function MyListings(props) {
  return(
    <div className="col">
      <Text color="black" size="24px" weight="800">My Listings</Text>
      {props.items.map(item => <ShortListing key={item.id} item={item} editListing={props.editListing}/>)}
    </div>
  );
}