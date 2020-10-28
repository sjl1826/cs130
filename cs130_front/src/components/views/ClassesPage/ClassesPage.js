import React from 'react';
import Tabs from '../../Tabs/Tabs';
import Dropdown from '../../Dropdown/Dropdown'
import '../../../App.css';

function ClassesPage() {
  const items = [{name: "Wow", groupId: 123}, {name: "CA", groupId: 123}, {name: "DM", groupId: 123},]
  return (
    <div className="App">
      <Tabs>
        <div type="Study Buddies"></div>
        <div type="Listings"></div>
      </Tabs>
      <Dropdown width="100vw" options={items} />
    </div>
  );
}

export default ClassesPage;