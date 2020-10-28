import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import Text from '../Text/Text';
import triangle from './assets/triangle.png';

export default function DropdownRow({width='20vw', ...props}) {
	function changeSelection(event, item) {
		props.changeSelection(event, item);
	}

	return (
		<div
			className={css`
        display: flex;
        justify-content: space-between;
        align-items: center;
        border: 2px solid ${Colors.Black};
        width: ${width};
        &:hover {
          cursor: pointer;
        }
      `}
			onClick={event => changeSelection(event, props.item)}
		>
			<Text size="1rem">
				{props.item.name}
			</Text>
			{props.showTriangle ? <img className={css`height: 1.5rem`} src={triangle}/> : null}
		</div>
	);
}
