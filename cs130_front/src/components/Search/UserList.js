import React from 'react';
import './UserList.css'

const UserList = (props) => {
  const ListStyling = {display: "flex", justifyContent: "center"}

  return (
    <>
    { props.userList.map((data,index) => {
        if (data) {
          return (
            <div data-testid="user" style={ListStyling} key={data.name} onClick={props.goToUserProfile(data)}>
              <h1 className="listItem" onMouseOver={() => {}}>{data.name}</h1>
	          </div>	
    	    );	
    	  }
    	  return null;
    })
    }
    </>
  );
}

export default UserList