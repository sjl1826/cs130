import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import * as Fonts from '../../constants/Fonts';

function TextArea({
	border = `2px solid ${Colors.Blue}`,
	height = '42px',
	width = '300px',
	...props
}) {
	return (
		<textarea
			className={css`
				border: ${border};
				height: ${height};
				width: ${width};
				border-radius: 16px;
				padding: 12px;
				font-family: ${Fonts.Primary};
				outline: none;
			`}
			type="text"
			{...props}
		/>
	);
}

export default TextArea;