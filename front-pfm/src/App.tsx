import React from 'react';
import './App.css';
import AppBar from "./Components/AppBar";
import TasksTab from "./Tabs/TasksTab";
import { Routes, Route } from "react-router-dom";
import BatchesTab from "./Tabs/BatchesTab";
import ResultsTab from "./Tabs/ResultsTab";

function App() {
    return (
        <div className="App">
            <AppBar/>
            <Routes>
                <Route path="/" element={<TasksTab/>} />
                <Route path="/tasks" element={<TasksTab/>} />
                <Route path="/batches" element={<BatchesTab/>} />
                <Route path="/results" element={<ResultsTab/>} />
            </Routes>
        </div>
    );
}

export default App;
