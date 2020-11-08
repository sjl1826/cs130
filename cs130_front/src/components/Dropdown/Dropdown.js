import React, {  useState, useEffect } from 'react';
import { css } from 'emotion';
import DropdownRow from './DropdownRow';

export default function Dropdown({width = '20vw', options=[], ...props}) {
  const dropdown = css`
    display: flex;
    flex-direction: column;
    align-items: center; 
    width: ${width};
  `;

  const dropdownAbsolute = css`
    display: flex;
    flex-direction: column;
    align-items: center; 
    width: ${width};
    position: absolute;
    background-color: white;
    margin-top: 1.6rem;
  `;

	const [items, setItems] = useState(options);
	const [showMenuState, setMenu] = useState(false);
  const [selected, setSelected] = useState(options[0]);
  const { sendSelection } = props;
  useEffect(() => {
    sendSelection(selected);
  }, [selected, sendSelection]);

	function showMenu(event, item) {
		event.preventDefault();
		setMenu(!showMenuState);
    setSelected(item);
	}

	if (items === null || selected === null) {
		return null;
  }

	return (
		<div className={dropdown}>
			{<DropdownRow item={selected != undefined ? selected : 'None available'} changeSelection={showMenu} showTriangle={true}/>}
			{
				showMenuState && items.length > 0 ?
					<div className={dropdownAbsolute}>
						{
							items.map(item =>
								<DropdownRow
									key={item.name}
									item={item}
									changeSelection={showMenu}
									showTriangle={false}
								/>)
						}
					</div> :
					null
			}
		</div>
	);
}
