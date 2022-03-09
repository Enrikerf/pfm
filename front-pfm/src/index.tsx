import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import TasksTab from "./Tabs/TasksTab/TasksTab";
import StepsTab from "./Tabs/StepsTab";
import BatchesList from "./Tabs/BatchesTab/BatchesList";
import ResultsTab from "./Tabs/ResultsTab";
import RunnerTab from "./Tabs/RunnerTab/RunnerTab";
import BatchesTab from "./Tabs/BatchesTab/BatchesTab";


ReactDOM.render(
    <React.StrictMode>
        <meta name="viewport" content="initial-scale=1, width=device-width"/>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<App/>}>
                    <Route index element={<TasksTab/>}/>
                    <Route path=":uuid/steps" element={<StepsTab/>}/>
                    <Route path=":uuid/batches" element={<BatchesTab/>}>
                        <Route index element={<BatchesList/>}/>
                        <Route path=":uuid/results" element={<ResultsTab/>}/>
                    </Route>
                    <Route path=":uuid/runner" element={<RunnerTab/>}/>
                </Route>
                <Route
                    path="*"
                    element={
                        <main style={{padding: "1rem"}}>
                            <p>There's nothing here!</p>
                        </main>
                    }
                />
            </Routes>
        </BrowserRouter>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
