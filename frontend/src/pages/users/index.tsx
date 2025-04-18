import type { Metadata } from 'next'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import React from "react";
export const metadata: Metadata = {
    title: 'GoWeb',
    description: 'GoWeb前端',
}

const App = () => {
    return <div>
       hello
    </div>
}

export default App
