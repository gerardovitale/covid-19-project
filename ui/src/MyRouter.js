import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import BarChart from './charts/BarChart';
import TestData from './charts/testData';


const MyRouter = () => {

    return (
        <BrowserRouter>
            <Routes>
                <Route path='/' element={<BarChart data={TestData} />}></Route>
            </Routes>
        </BrowserRouter>
    )
}

export default MyRouter;
