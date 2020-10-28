import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import * as Fonts from '../../constants/Fonts';

function TextInput({
	border = `2px solid ${Colors.Blue}`,
	height = '42px',
	width = '300px',
	...props
}) {
	return (
		<input
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

export default TextInput;

// Usage:
// <TextInput
//   value={value}
//   onChange={(event) => {
//     setValue(event.target.value);
//   }}
// />