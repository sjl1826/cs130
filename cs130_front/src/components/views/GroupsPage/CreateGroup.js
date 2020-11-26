import React, { useState } from 'react';
import FormInput from '../../Infos/FormInput';
import Button from '../../Button/Button';
import * as Colors from '../../../constants/Colors';
import './styles.css';

export default function CreateGroup(props) {
    const [firstInput, setFirstInput] = useState({ name: props.options[0].name, value: props.options[0].value });
    function sendInput(name, value) {
        switch (name) {
            case firstInput.name:
                setFirstInput({ name: firstInput.name, value: value })
                break;
            default:
                break;
        }
    }
    return (
        <div className="col">
            {props.options.map(option => <FormInput key={option.name} option={option} sendInput={sendInput} />)}
            <Button
                textColor={Colors.White}
                textSize="28px"
                width="275px"
                height="65px"
                textWeight="800"
                color={Colors.Blue}
                onClick={() => props.saveInfoClicked(firstInput)}
            >
                Create Group
            </Button>
        </div>
    );
}
