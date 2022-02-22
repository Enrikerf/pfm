import React from 'react';
import logo from './logo.svg';
import './App.css';
import {TaskServiceClient} from "./protobuf/generated/task_grpc_web_pb"

function App() {

  const call = function ():void {

    const messages = require('./protobuf/generated/task_pb');
    let listTaskRequest = new messages.ListTasksRequest()
    let metadata = {};
    let taskService = new TaskServiceClient("http://localhost:8080", null, null)
    taskService.listTasks(listTaskRequest,metadata,function (err,response){
      if (err){
        console.log(err);
      }else{
        console.log(response);
      }
    })
  }
  call()
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
