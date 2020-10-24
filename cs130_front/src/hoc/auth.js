import React, { useEffect } from 'react';
import { auth } from '../_actions/user_actions';
import { useSelector, useDispatch } from 'react-redux';

export default function (SpecificComponent, option, adminRoute = null) {
	function AuthenticationCheck(props) {
		const user = useSelector(state => state.user);
		const dispatch = useDispatch();
		useEffect(() => {
			// To know my current status, send Auth request
			dispatch(auth()).then(response => {
				// Not Loggined in Status
				if (!response.payload.isAuth) {
					if (option) {
						props.history.push('/login');
					}
					// Loggined in Status
				} else {
					if (adminRoute && !response.payload.isAdmin) {
						// supposed to be Admin page, but not admin person wants to go inside
						props.history.push('/');
					} else if (option === false) {
					// Logged in Status, but Try to go into log in page
						props.history.push('/');
					}
				}
			})
				.catch(() => {
					if (adminRoute) {
						props.history.push('/');
					}
				});
		}, []);

		return (
			<SpecificComponent {...props} user={user} />
		);
	}
	return AuthenticationCheck;
}