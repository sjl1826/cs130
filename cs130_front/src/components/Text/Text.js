import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import * as Fonts from '../../constants/Fonts';

function Text({
	children,
	color = Colors.Black,
	size = '18px',
	weight = 400,
	error = false,
	...props
}) {
	return (
		<div
			className={css`
				color: ${color};
				font-size: ${size};
				font-weight: ${weight};
				font-family: ${Fonts.Primary};
			`}
			{...props}
		>
			{children}
		</div>
	);
}

export default Text;