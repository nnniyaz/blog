import { Layout } from '@components/layout'
import { Route, Routes } from 'react-router-dom'
import { RoutesList } from '@domain/base/routes-list.tsx'

export const AppRoutes = () => {
    return (
        <Routes>
            <Route element={<Layout />}>
                {RoutesList.map((route) => (
                    <Route
                        key={route.name}
                        path={route.path}
                        element={route.element}
                    />
                ))}
            </Route>
        </Routes>
    )
}
