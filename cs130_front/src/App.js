import React, { Suspense } from 'react';
import { Route, Switch } from 'react-router-dom';
import Auth from './hoc/auth';
// pages for this product
import LandingPage from './components/views/LandingPage/LandingPage.js';
import LoginPage from './components/views/LoginPage/LoginPage.js';
import RegisterPage from './components/views/RegisterPage/RegisterPage.js';
import ClassesPage from './components/views/ClassesPage/ClassesPage';
import GroupsPage from './components/views/GroupsPage/GroupsPage.js';
import ProfilePage from './components/views/ProfilePage/ProfilePage.js';
import NavBar from './components/views/NavBar/NavBar';
import SearchPage from './components/SearchPage/SearchPage.js';

// null   Anyone Can go inside
// true   only logged in user can go inside
// false  logged in user can't go inside

function App() {
	return (
		<Suspense fallback={<div>Loading...</div>}>
			<NavBar />
			<div style={{ paddingTop: '69px', minHeight: 'calc(100vh - 80px)' }}>
				<Switch>
					<Route exact path="/" component={Auth(LandingPage, null)} />
					<Route exact path="/login" component={Auth(LoginPage, false)} />
					<Route exact path="/register" component={Auth(RegisterPage, false)} />
          			<Route exact path="/classes/" component={Auth(ClassesPage, true)} />
					<Route exact path="/groups/" component={Auth(GroupsPage, true)} />
          			<Route exact path="/profile/" component={Auth(ProfilePage, true)} />
					<Route exact path="/search/" component={Auth(SearchPage, true)} />
				</Switch>
			</div>
		</Suspense>
	);
}

export default App;