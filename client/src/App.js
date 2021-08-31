import './App.css';
import React from 'react';
import Chessground from 'react-chessground'
import 'react-chessground/dist/styles/chessground.css'

export default class App extends React.Component {
  onMove(from, to ) {
    console.log(from)
    console.log(to)
  }

  render () {
    return <Chessground onMove={this.onMove}/>
  }
}
