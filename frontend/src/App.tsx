import { VFC } from 'react';
import Header from 'component/header/header';
import Home from 'page/home/home';
import DiaryEntry from 'page/diary/diaryEntry';
import DiaryView from 'page/diary/diaryView';
import Search from 'page/search/search';
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";

const App: VFC = () =>{
    return(
        <Router>
            <Header />
            <Routes>
                <Route path='/' element={<Home />} />
                <Route path='/entry/:newOrEdit' element={<DiaryEntry />} />
                <Route path='/view' element={<DiaryView />} />
                <Route path='/search' element={<Search />} />
                <Route path='*' element={<Navigate to='/' replace/>} />
            </Routes>
        </Router>
    )
}


export default App;