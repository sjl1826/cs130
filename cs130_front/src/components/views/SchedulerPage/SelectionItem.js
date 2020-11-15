import React, { Component } from 'react';
import './styles.css';

export default class  SelectionItem extends Component {
  render() {
    var className='item2 noselect';
    className += (this.props.isSelected ? ' selected' : '');
    return (
      <div className={className}>
        {this.props.data.time}
      </div>
    );
  }

}