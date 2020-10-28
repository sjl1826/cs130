import React, { useState } from 'react';
import { css } from 'emotion';
import Tab from './Tab';

export default function Tabs(props) {
	const [activeTab, setActiveTab] = useState(React.Children.toArray(props.children)[0].props.type);
	return (
		<div>
			<ol
				className={css`
          display: flex;
        `}
			>
				{React.Children.map(props.children, child => {
					return (
						<Tab
							activeTab={activeTab}
							key={child.props.type}
							label={child.props.type}
							onClick={tab => setActiveTab(tab)}
						/>
					);
				})}
			</ol>
			<div>
				{React.Children.map(props.children, child => {
					return child.props.type !== activeTab ? undefined : child;
				})}
			</div>
		</div>
	);
}
