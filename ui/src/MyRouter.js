import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Todo from './charts/Todo';

const MyRouter = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/' element={<Todo />}></Route>
            </Routes>
        </BrowserRouter>
    )
}

export default MyRouter;
