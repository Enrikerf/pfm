import React from 'react';
import './App.css';
import AppBar from "./Components/AppBar";
import {Outlet} from "react-router-dom";


function App() {
    return (
        <div className="App">
            <AppBar/>
            <Outlet/>
        </div>
    );
}

export default App;
