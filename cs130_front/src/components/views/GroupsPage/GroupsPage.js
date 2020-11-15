import React, { useState, useEffect } from 'react';
import logo from '../../../logo.svg';
import '../../../App.css';

function GroupsPage() {


  const classes = [
    {name: "Discrete Mathematics", courseId: 1, groups: [{name: "DM Squad", groupId: 1}, {name: "DM Squad II", groupId: 1}] }, 
    {name: "Computer Architecture", courseId: 2, groups: [{name: "CA Squad", groupId: 3}]}
  ]
  const [mainPanelState, setMainPanel] = useState();

  // Pass this to ClassList
  function groupClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }


  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default GroupsPage;