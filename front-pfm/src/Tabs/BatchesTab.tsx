import * as React from "react";
import {Outlet} from "react-router-dom";


function BatchesTab() {

    return (
        <div>
            <h1>Batches</h1>
            <Outlet/>
        </div>

    );

}

export default BatchesTab