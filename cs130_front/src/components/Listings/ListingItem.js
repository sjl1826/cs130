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
                    <h3 className="header">{props.data.poster}</h3>
                </div>
                <div className="row" style={{marginTop: "-10px"}}>
                    <h3 style={{color: Colors.Gray}}>{props.data.school} </h3>
                </div>
                <div className="row" style={{marginTop: "10px"}}>
                    <div className="desc" >{props.data.description}</div>
                </div>
            </div>
        </div>
    );
} 