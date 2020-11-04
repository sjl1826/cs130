import React from 'react';

const UserList = (props) => { 
  return (
    <>
    { props.userList.map((data,index) => {
        if (data) {
          return (
            <div key={data.name} onClick={props.goToUserProfile(data)}>
              <h1>{data.name}</h1>
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