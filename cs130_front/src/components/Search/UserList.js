import React from 'react';
import './UserList.css'

const UserList = (props) => {
  const ListStyling = {display: "flex", justifyContent: "center"}

  return (
    <>
    { props.userList.map((data,index) => {
        if (data) {
          return (
            <div style={ListStyling} key={data.name} onClick={() => props.goToUserProfile(data)}>
              <h1 className="listItem" onmouseover="">{data.name}</h1>
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