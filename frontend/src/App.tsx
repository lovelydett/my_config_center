import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import HomePage from "./pages/home";
import ExperimentalPage from "./pages/experimental";

// P for props, S for state, SS for snapshot
class App extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Routes>
                    <Route path="/" Component={HomePage} />
                    <Route path="/experimental" Component={ExperimentalPage} />
                </Routes>
            </BrowserRouter>
        );
    }
}

export default App;
