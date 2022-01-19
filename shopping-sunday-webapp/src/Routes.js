import React from 'react';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {HomePage} from "./home";

export const RoutesList = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route exact path="/:date" element={<HomePage/>}></Route>
            </Routes>
        </BrowserRouter>
    );
};
