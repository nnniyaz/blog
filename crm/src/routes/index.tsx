import { Layout } from '@components/layout'
import { Navigate, Route, Routes } from 'react-router-dom'
import {
    PrivateRoutesList,
    PublicRoutesList,
} from '@domain/base/routes-list.tsx'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { RoutesPaths } from '@domain/base/routes-paths.ts'
import { useActions } from '@lib/hooks/useActions.tsx'
import { useEffect } from 'react'

export const AppRoutes = () => {
    const { isAuth, isLoadingSignIn } = useTypedSelector((state) => state.auth)
    const { currentUser, isLoadingGetCurrentUser } = useTypedSelector((state) => state.currentUser)
    const { getCurrentUser } = useActions()

    useEffect(() => {
        if (!currentUser) {
            getCurrentUser()
        }
    }, [])

    if (isLoadingGetCurrentUser && !isLoadingSignIn) {
        return <div>Loading...</div>
    }

    return (
        <Routes>
            {isAuth ? (
                <Route element={<Layout />}>
                    {PrivateRoutesList.map((route) => (
                        <Route
                            key={route.name}
                            path={route.path}
                            element={route.element}
                        />
                    ))}
                    <Route
                        path={'*'}
                        element={
                            <Navigate to={RoutesPaths.DASHBOARD} replace />
                        }
                    />
                </Route>
            ) : (
                <>
                    {PublicRoutesList.map((route) => (
                        <Route
                            key={route.name}
                            path={route.path}
                            element={route.element}
                        />
                    ))}
                    <Route
                        path={'*'}
                        element={<Navigate to={RoutesPaths.SIGN_IN} replace />}
                    />
                </>
            )}
        </Routes>
    )
}
