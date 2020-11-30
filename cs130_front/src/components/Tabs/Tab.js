import React from 'react';
import { css } from 'emotion';
import * as Colors from '../../constants/Colors';
import Text from '../Text/Text';

const tabListItem = css`
  list-style: none;
  border-width: 1px;
  padding: 0.5rem;
  width: 250px;
  text-align: center;
  &:hover {
    cursor: pointer;
  } 
`;
const tabListActive = css`
  ${tabListItem}
  border-bottom: solid ${Colors.Blue};
`;

export default function Tab(props) {
	return (
		<li
			className={props.activeTab === props.label ? tabListActive : tabListItem}
			onClick={() => props.onClick(props.label)}
		>
			<Text size="2rem">{props.label}</Text>
		</li>
	);
}
