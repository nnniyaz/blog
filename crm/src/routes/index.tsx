import {BrowserRouter, Route, Routes} from 'react-router-dom'


import {Article} from './pages/article'
import {Author} from './pages/author'
import {Book} from './pages/book'
import {Home} from './pages/home/Home'
import {Project} from './pages/project'
import {Settings} from './pages/settings'
import {User} from './pages/user'
import {Layout} from '../components/layout'
import {Routes as RoutesEnum} from '../domain/base/routes/routes'

export const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route element={<Layout/>}>
                    <Route path={RoutesEnum.HOME} element={<Home/>}/>
                    <Route path={RoutesEnum.USER} element={<User/>}/>
                    <Route path={RoutesEnum.AUTHOR} element={<Author/>}/>
                    <Route path={RoutesEnum.PROJECT} element={<Project/>}/>
                    <Route path={RoutesEnum.ARTICLE} element={<Article/>}/>
                    <Route path={RoutesEnum.BOOK} element={<Book/>}/>
                    <Route path={RoutesEnum.SETTINGS} element={<Settings/>}/>
                </Route>
            </Routes>
        </BrowserRouter>
    )
}