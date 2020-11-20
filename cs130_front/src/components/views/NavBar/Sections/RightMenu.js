import React from 'react';
import { Menu } from 'antd';
import { withRouter } from 'react-router-dom';

function RightMenu(props) {
	const logoutHandler = () => {
    localStorage.setItem('userId', 'nothing');
		props.history.push('/');
	};

  console.log(localStorage.getItem('userId'));
	if (localStorage.getItem('userId') !== 'nothing') {
		return (
      <Menu mode={props.mode}>
			<Menu.Item key="classes">
				<a href="/classes">Classes</a> 
			</Menu.Item>
			<Menu.Item key="groups">
				<a href="/groups">Groups</a>
			</Menu.Item>
      <Menu.Item key="profile">
				<a href="/profile">Profile</a>
			</Menu.Item>
			<Menu.Item key="logout">
				<a onClick={logoutHandler}>Logout</a>
			</Menu.Item>
		</Menu>
		);
  }
  // we need to pass current user's id here in the future.
	return (
		<Menu mode={props.mode}>
			<Menu mode={props.mode}>
				<Menu.Item key="mail">
					<a href="/login">Login</a>
				</Menu.Item>
				<Menu.Item key="app">
					<a href="/register">Register</a>
				</Menu.Item>
			</Menu>
    </Menu>
	);
}

export default withRouter(RightMenu);
