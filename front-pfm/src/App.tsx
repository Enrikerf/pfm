import React from 'react';
import './App.scss';
import AppBar from "./Components/AppBar";
import {Outlet} from "react-router-dom";


function App() {
    return (
        <div className="App">
            <AppBar className="App-header"/>
            <div className="App-content">
                <Outlet/>
            </div>
        </div>
    );
}

export default App;
