import React, { useState } from 'react';
import * as Fonts from '../../constants/Fonts';

export default function ClassTitle(props) {
  return(
    <div >
      <h1 style={{fontFamily: Fonts.Primary, fontSize: "40px", fontWeight: "bold"}}>{props.option}</h1>
    </div>
  );
} 