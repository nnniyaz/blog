import { Layout } from '@components/layout'
import { Routes as RoutesEnum } from '@domain/base/routes/routes.tsx'
import { Article } from '@routes/pages/article'
import { Author } from '@routes/pages/author'
import { Book } from '@routes/pages/book'
import { Home } from '@routes/pages/home/Home.tsx'
import { Project } from '@routes/pages/project'
import { Settings } from '@routes/pages/settings'
import { User } from '@routes/pages/user'
import { BrowserRouter, Route, Routes } from 'react-router-dom'

export const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route element={<Layout />}>
                    <Route path={RoutesEnum.HOME} element={<Home />} />
                    <Route path={RoutesEnum.USER} element={<User />} />
                    <Route path={RoutesEnum.AUTHOR} element={<Author />} />
                    <Route path={RoutesEnum.PROJECT} element={<Project />} />
                    <Route path={RoutesEnum.ARTICLE} element={<Article />} />
                    <Route path={RoutesEnum.BOOK} element={<Book />} />
                    <Route path={RoutesEnum.SETTINGS} element={<Settings />} />
                </Route>
            </Routes>
        </BrowserRouter>
    )
}
