import React, { useState } from 'react';
import { css } from 'emotion';
import Tab from './Tab';

export default function Tabs(props) {
  function setTab(tab) {
    setActiveTab(tab);
    props.setTabVar(tab);
  }
	const [activeTab, setActiveTab] = useState(React.Children.toArray(props.children)[0].props.type);
	return (
		<div>
			<ol
				className={css`
          display: flex;
          justify-content: center;
        `}
			>
				{React.Children.map(props.children, child => {
					return (
						<Tab
							activeTab={activeTab}
							key={child.props.type}
							label={child.props.type}
							onClick={tab => setTab(tab)}
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
