import React, { useEffect } from 'react';
import { auth } from '../_actions/user_actions';
import { useSelector, useDispatch } from 'react-redux';

export default function (SpecificComponent, option, adminRoute = null) {
	function AuthenticationCheck(props) {
		const user = useSelector(state => state.user);
		const dispatch = useDispatch();
		return (
			<SpecificComponent {...props} user={user} />
		);
	}
	return AuthenticationCheck;
}