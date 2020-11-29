import React from 'react';
import { Menu } from 'antd';
import { withRouter } from 'react-router-dom';

function RightMenu(props) {
	const logoutHandler = () => {
    localStorage.setItem('userId', 'nothing');
		props.history.push('/');
	};


	if (localStorage.getItem('userId') !== 'nothing') {
    const classesRef = `/classes/${localStorage.getItem('userId')}`
    const groupsRef = `/groups/${localStorage.getItem('userId')}`
    const profileRef = `/profile/${localStorage.getItem('userId')}`
		return (
      <Menu mode={props.mode}>
      <Menu.Item key="search">
				<a href={"/search"}>Search</a> 
			</Menu.Item>
			<Menu.Item key="classes">
				<a href={classesRef}>Classes</a> 
			</Menu.Item>
			<Menu.Item key="groups">
				<a href={groupsRef}>Groups</a>
			</Menu.Item>
      <Menu.Item key="profile">
				<a href={profileRef}>Profile</a>
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
