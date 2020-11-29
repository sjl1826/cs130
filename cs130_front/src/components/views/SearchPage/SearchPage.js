import React, { useState, useEffect } from 'react';
import axios from 'axios';
import SearchBar from '../../Search/SearchBar';
import UserList from '../../Search/UserList';
import * as Fonts from '../../../constants/Fonts';
import { ROUTE } from '../../../Config';

const SearchPage = (props) => {
  const [input, setInput] = useState('');
  const [userListDefault, setUserListDefault] = useState();
  const [userList, setUserList] = useState();

  const userId = localStorage.getItem('userId');
  const config = {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
    }
  }

  useEffect(() => {
		async function initUsers() {
			try {
        const userResponse = await getUsers();
        console.log(userResponse);
        handleUserResponse(userResponse.data.users);
			} catch (err) {
				// Handle err here. Either ignore the error, or surface the error up to the user somehow.
			}
		}
    initUsers();
  }, []);

  function getUsers(){
    return axios.get(`${ROUTE}/getAllUsers?u_id=${userId}`, config);
  }

  function handleUserResponse(data){
    const users = [];
    Object.keys(data).forEach(function(key) {
      const name = data[key]["FirstName"] + " " + data[key]["LastName"];
      const id = data[key]["ID"].toString();
      users.push({name: name, id: id});
    });
    setUserListDefault(users);
  }

  const updateInput = async (input) => {
    const filtered = userListDefault.filter(user => {
      return (user.name.toLowerCase().includes(input.toLowerCase()) || user.id.toLowerCase().includes(input.toLowerCase()))
    })
    setInput(input);
    setUserList(filtered);
  }

  const goToUserProfile = user => () => { props.history.push(`/profile/${user.id}`); }
  
  return (
    <div style={{fontFamily: Fonts.Primary, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
      <h1 style = {{fontSize: '40px'}}>Search Users</h1>
      <div> 
        <SearchBar input={input} onChange={updateInput} width="30rem" fontSize="30px"/>
      </div>
      { input && <UserList userList={userList} goToUserProfile={goToUserProfile}/> }
    </div>
  );
}

export default SearchPage