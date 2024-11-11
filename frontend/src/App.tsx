import React from "react";
import { BrowserRouter } from "react-router-dom";

import HomePage from "./pages/home";

// P for props, S for state, SS for snapshot
class App extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <HomePage />
            </BrowserRouter>
        );
    }
}

export default App;
