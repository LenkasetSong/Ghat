// src/App.js 
import React, {Component} from "react";
import "./App.css";
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log('hello world!');
    sendMsg('hello world!');
  }

  render() {
    return (
      <div className="App">
        <button onClick={this.send}>Click</button>
      </div>
    );
  }
}

export default App;