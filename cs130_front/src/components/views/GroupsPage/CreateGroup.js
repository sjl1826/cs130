import React, { useState } from 'react';
import FormInput from '../../Infos/FormInput';
import Button from '../../Button/Button';
import Text from '../../Text/Text';
import * as Colors from '../../../constants/Colors';
import './styles.css';
import Dropdown from '../../Dropdown/Dropdown';

export default function CreateGroup(props) {
    const [firstInput, setFirstInput] = useState({ name: props.options[0].name, value: props.options[0].value });
    const [classSelection, setSelection] = useState(props.courses != null ? props.courses[0] : null);
    function sendInput(name, value) {
        switch (name) {
            case firstInput.name:
                setFirstInput({ name: firstInput.name, value: value })
                break;
            default:
                break;
        }
    }

    function sendSelection(item) {
      setSelection(item);
    }
    return (
        <div className="col">
            <Text weight="800" size="24px">Create a new group</Text>
            <div className="col-spaced">
              {props.options.map(option => <FormInput key={option.name} option={option} sendInput={sendInput} />)}
            </div>
            <div className="group-with-margin-bottom">
              <Text> For course:</Text>
              <Dropdown width="20vw" options={props.courses} sendSelection={sendSelection}/>
            </div>
            <Button
                textColor={Colors.White}
                textSize="28px"
                width="275px"
                height="65px"
                textWeight="800"
                color={Colors.Blue}
                onClick={() => props.createGroup(firstInput, classSelection)}
            >
                Create Group
            </Button>
        </div>
    );
}
