import React, { useState, useEffect } from 'react';
import SearchBar from '../../Search/SearchBar';
import UserList from '../../Search/UserList';
import * as Fonts from '../../../constants/Fonts';

const SearchPage = (props) => {
  const [input, setInput] = useState('');
  const [userListDefault, setUserListDefault] = useState();
  const [userList, setUserList] = useState();

  const fetchData = async () => {
    return await fetch('https://restcountries.eu/rest/v2/all')
      .then(response => response.json())
      .then(data => {
         setUserList(data) 
         setUserListDefault(data)
        }
      );
  }

  const updateInput = async (input) => {
    const filtered = userListDefault.filter(user => {
      return user.name.toLowerCase().includes(input.toLowerCase())
    })
    setInput(input);
    setUserList(filtered);
  }

  useEffect( () => {fetchData()},[]);
  const goToUserProfile = user => () => { props.history.push(`/profile/${user.name}`); }
  
  return (
    <div style={{fontFamily: Fonts.Primary, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
      <h1 style = {{fontSize: '40px'}}>Search User</h1>
      <div> 
        <SearchBar input={input} onChange={updateInput} width="30rem" fontSize="30px"/>
      </div>
      { input && <UserList userList={userList} goToUserProfile={goToUserProfile}/> }
    </div>
  );
}

export default SearchPage