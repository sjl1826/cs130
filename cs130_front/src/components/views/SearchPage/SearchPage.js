import React, { useState, useEffect } from 'react';
import SearchBar from '../../Search/SearchBar';
import UserList from '../../Search/UserList';

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
    <>
      <h1>Search User</h1>
      <SearchBar 
       input={input} 
       onChange={updateInput}
      />
      { input && <UserList userList={userList} goToUserProfile={goToUserProfile}/> }
    </>
  );
}

export default SearchPage