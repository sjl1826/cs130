import React, { useState } from 'react';
import Text from '../Text/Text';
import * as Fonts from '../../constants/Fonts';
import * as Colors from '../../constants/Colors';
import './ListingItem.css';

export default function ListingItem(props) {
    return(
        <div className="box-shadow" style={{textAlign: "center"}}>
            <div style={{display: "flex", flexDirection: "column"}}>
                <div className="row">
                    <div className="col">
                        <h3 className="header clickable" onClick={props.goToUserProfile(props.data.poster)} onmouseover="">{props.data.poster}</h3>
                        <h3 style={{color: Colors.Gray, marginTop: "-10px"}}>{props.data.school} </h3>
                    </div>
                    {props.data.groupId ?
                        <div className="col">
                            <h3 className="header">Group:</h3>
                            <h3 style={{color: Colors.Blue, marginTop: "-10px"}} className="clickable" onClick={props.goToGroup(props.data.groupId)}>{props.data.groupName}</h3>
                        </div> 
                        :
                        <div></div>
                    }
                </div>
                <div className="row" style={{marginTop: "10px"}}>
                    <div className="desc" >{props.data.description}</div>
                </div>
            </div>
        </div>
    );
} 