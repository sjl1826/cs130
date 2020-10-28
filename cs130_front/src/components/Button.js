import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import Text from '../Text/Text';

function Button({
	children,
	color = Colors.Blue,
	height = '42px',
	width = '150px',
	disabled = false,
	...props
}) {
	return (
		<button
			className={css`
				height: ${height};
				width: ${width};
				background-color: ${disabled ? Colors.Gray : color};
				border-radius: 16px;
				display: flex;
				align-items: center;
				justify-content: center;
				cursor: ${disabled ? 'not-allowed' : 'pointer'};
				border: none;
				&:hover {
					background-color: ${disabled ? Colors.Gray : Colors.DarkBlue};
					transition: 0.3s;
				}
			`}
			{...props}
		>
			<Text color={Colors.White}>{children}</Text>
		</button>
	);
}

export default Button;

// Usage:
/*
<Button
  onClick={() => {
    console.log("hello");
  }}
  disabled
>
  Hello
</Button>;
*/