import React from 'react';
import * as Colors from '../../constants/Colors';
import * as Fonts from '../../constants/Fonts';

const SearchBar = ({input:keyword, onChange:setKeyword}) => {
  const BarStyling = {fontFamily: Fonts.Primary, 
                      border: `2px solid ${Colors.Blue}`, 
                      borderRadius: '35px',
                      width: '30rem',
                      textAlign: 'left',
                      outline: 'none',
                      paddingLeft: '1.5rem',
                      fontSize: '30px',
                      background: '#F2F1F9',
                      marginBottom: '30px'}
  return (
    <input 
     style={BarStyling}
     key="random1"
     value={keyword}
     placeholder={"Search"}
     onChange={(e) => setKeyword(e.target.value)}
    />
  );
}

export default SearchBar