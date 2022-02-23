import React from 'react';
import logo from './logo.svg';
import './App.css';
// import {TaskServiceClient} from "./protobuf/generated/task_grpc_web_pb"
import {ResultServiceClient} from "./protobuf/generated/result_grpc_web_pb"

function App() {

/*  const getTasks = function ():void {
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
  }*/
  const getResults = function ():void {
    const messages = require('./protobuf/generated/result_pb');
    let listTaskRequest = new messages.StreamResultsRequest({"batch_uuid":""})
    let metadata = {};
    let taskService = new ResultServiceClient("http://localhost:8080", null, null)
    let stream = taskService.streamResults(listTaskRequest,metadata)

    // @ts-ignore
    stream.on('data', function(response) {
      console.log("data");
      debugger
      console.log(response.getResponseMessage());
    });
    // @ts-ignore
    stream.on('status', function(status) {
      console.log("status");
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    // @ts-ignore
    stream.on('end', function(end) {
      console.log("end");
      stream.cancel()
    });

// to close the stream

  }


  // getTasks()
  getResults()
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
