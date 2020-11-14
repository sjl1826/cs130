import React, { useState } from 'react';
import Selection from './Selection';
import SelectionItem from './SelectionItem';
import _ from 'lodash';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import './styles.css';

const times = [
  {time: "12:00am", slot: 0 }, 
  {time: "12:30am", slot: 1 }, 
  {time: "1:00am", slot: 2 }, 
  {time: "1:30am", slot: 3 },
  {time: "2:00am", slot: 4 }, 
  {time: "2:30am", slot: 5 }, 
  {time: "3:00am", slot: 6 }, 
  {time: "3:30am", slot: 7 },
  {time: "4:00am", slot: 8 }, 
  {time: "4:30am", slot: 9 }, 
  {time: "5:00am", slot: 10 }, 
  {time: "5:30am", slot: 11 },
  {time: "6:00am", slot: 12 }, 
  {time: "6:30am", slot: 13 }, 
  {time: "7:00am", slot: 14 }, 
  {time: "7:30am", slot: 15 },
  {time: "8:00am", slot: 16 }, 
  {time: "8:30am", slot: 17 }, 
  {time: "9:00am", slot: 18 },
  {time: "9:30am", slot: 19 }, 
  {time: "10:00am", slot: 20 }, 
  {time: "10:30am", slot: 21 }, 
  {time: "11:00am", slot: 22 },
  {time: "11:30am", slot: 23 }, 
  {time: "12:00pm", slot: 24 },
  {time: "12:30pm", slot: 25 },
  {time: "1:00pm", slot: 26 },
  {time: "1:30pm", slot: 27 },
  {time: "2:00pm", slot: 28 },
  {time: "2:30pm", slot: 29 },
  {time: "3:00pm", slot: 30 },
  {time: "3:30pm", slot: 31 },
  {time: "4:00pm", slot: 32 },
  {time: "4:30pm", slot: 33 },
  {time: "5:00pm", slot: 34 },
  {time: "5:30pm", slot: 35 },
  {time: "6:00pm", slot: 36 },
  {time: "6:30pm", slot: 37 },
  {time: "7:00pm", slot: 38 },
  {time: "7:30pm", slot: 39 },
  {time: "8:00pm", slot: 40 },
  {time: "8:30pm", slot: 41 },
  {time: "9:00pm", slot: 42 },
  {time: "9:30pm", slot: 43 },
  {time: "10:00pm", slot: 44 },
  {time: "10:30pm", slot: 45 },
  {time: "11:00pm", slot: 46 },
  {time: "11:30pm", slot: 47 },
]

function SchedulerPage(props) {
  var selections = new Array(336).fill(0);
  var initSelections = new Array(336).fill(0);
  var userId = '';
  if(props.location != undefined) {
    const { state } = props.location;
    if (state != undefined && state.availability != undefined) {
      initSelections = state.availability;
      selections = state.availability;
    } 
    userId = props.match.params.id;
  }
  if(props.passedSelections != undefined) {
    initSelections = props.passedSelections;
    selections = props.passedSelections;
  }

  function updateSelection(items) {
    if (items.length < 1) {
      return;
    }
    var low = 0;
    var high = 48;
    if(items[0] >= 48 && items[0] < 96) {
      low = 48;
      high = 96;
    } else if (items[0] >= 96 && items[0] < 144) {
      low = 96;
      high = 144;
    } else if (items[0] >= 144 && items[0] < 192) {
      low = 144;
      high = 192;
    } else if (items[0] >= 192 && items[0] < 240) {
      low = 192;
      high = 240;
    }
    else if (items[0] >= 240 && items[0] < 288) {
      low = 240;
      high = 288;
    }
    else if (items[0] >= 288 && items[0] < 336) {
      low = 288;
      high = 336;
    }
    for(var i = low; i < high; i++) {
      if(items.includes(i.toString())) {
        selections[i] = 1;
      } else {
        selections[i] = 0;
      }
    }
  }

  function saveSelections() {
    // send selections array to backend
    //go back to profile and reload it? 
    props.history.push(`/profile/${userId}`)
  }

  const initSelectedSlots = [];

  var data = [];
  for(var i = 0; i < 7; i++) {
    for(var j = 0; j < 48; j++) {
      const newEntry = {time: times[j].time, slot: i * 48 + times[j].slot};
      if (initSelections[i * 48 + times[j].slot]) {
        initSelectedSlots.push((i * 48 + times[j].slot).toString())
      } else {
        initSelectedSlots.push("");
      }
      data.push(
        <SelectionItem key={newEntry.slot} data={newEntry}/>
      );
    }
  }
  const monInitial = initSelectedSlots.slice(0,48);
  const tuesInitial = initSelectedSlots.slice(48,96);
  const wedInitial = initSelectedSlots.slice(96,144);
  const thursInitial = initSelectedSlots.slice(144,192);
  const friInitial = initSelectedSlots.slice(192,240);
  const satInitial = initSelectedSlots.slice(240,288);
  const sunInitial = initSelectedSlots.slice(288,336);

  return (
    <div className="scheduler-container">
      <div className="scheduler-parent-container">
        <Text>Monday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={monInitial}>
          {data.slice(0,48)}
        </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Tuesday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={tuesInitial}>
          {data.slice(48,96)}
        </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Wednesday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={wedInitial}>
          {data.slice(96,144)}
        </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Thursday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={thursInitial}>
          {data.slice(144,192)}
        </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Friday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={friInitial}>
          {data.slice(192,240)}
        </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Saturday</Text>
      <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={satInitial}>
        {data.slice(240,288)}
      </Selection>
      </div>
      <div className="scheduler-parent-container">
        <Text>Sunday</Text>
        <Selection enabled={props.passedSelections == undefined} onSelectionChange={updateSelection} initialSelected={sunInitial}>
          {data.slice(288,336)}
        </Selection>
      </div>
      {props.passedSelections == undefined &&       
      <div className="column-center">
        <Button onClick={() => saveSelections()}>
          <Text color="white">Save</Text>
        </Button>
        <Text color>Indicate your availability! Drag to select or click. </Text>
      </div>
      }

    </div>
  );
}

export default SchedulerPage;