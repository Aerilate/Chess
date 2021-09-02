import './App.css';
import React from 'react';
import Chessground from 'react-chessground'
import 'react-chessground/dist/styles/chessground.css'

const TIMEOUT = 250;

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      ws: null,
      serverData: {
        fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
      }
    };
    this.onMove = this.onMove.bind(this);
    this.calcMovable = this.calcMovable.bind(this);
  }

  timeout = TIMEOUT;

  componentDidMount() {
    this.connect();
  }

  connect = () => {
    var ws = new WebSocket("ws://127.0.0.1:8080/ws");
    let that = this;
    var connectInterval;

    ws.onmessage = (ev) => {
      console.log(ev.data)
      this.setState({
        serverData: JSON.parse(ev.data)
      }, () => {
        const validMoves = this.state.serverData?.validMoves
        console.log(validMoves)
      })
    }

    ws.onopen = () => {
      console.log("Connected to server!");

      this.setState({ ws: ws });
      that.timeout = TIMEOUT; // reset timer to 250 on open of websocket connection 
      clearTimeout(connectInterval); // clear Interval on on open of websocket connection
    };

    ws.onclose = e => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
          10000 / 1000,
          (that.timeout + that.timeout) / 1000
        )} second.`,
        e.reason
      );

      that.timeout = that.timeout + that.timeout; //increment retry interval
      connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
    };

    // websocket onerror event listener
    ws.onerror = err => {
      console.error(
        "Socket encountered error: ",
        err.message,
        "Closing socket"
      );
      ws.close();
    };
  };

  check = () => {
    const { ws } = this.state;
    if (!ws || ws.readyState === WebSocket.CLOSED) this.connect();
  };

  onMove(from, to) {
    console.log(from)
    console.log(to)
    const { ws } = this.state;
    try {
      ws.send(from + to)
    } catch (error) {
      console.log(error)
    }
  }

  calcMovable = () => {
    let dests = new Map()
    const validMoves = this.state.serverData?.validMoves
    if (validMoves) {
      dests = new Map(Object.entries(validMoves))
    }

    return {
      free: false,
      dests,
      color: "both"
    }
  }

  render() {
    return <Chessground onMove={this.onMove} movable={this.calcMovable()} fen={this.state.serverData?.fen}/>;
  }
}
