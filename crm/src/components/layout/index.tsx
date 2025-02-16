import { Outlet } from 'react-router-dom'
import classes from './layout.module.scss'
import { Sidebar } from '@components/ui/sidebar/sidebar.tsx'

export const Layout = () => {


    return (
        <div className={classes.layout}>
            <Sidebar />
            <div className={classes.content}>
                <Outlet />
            </div>
        </div>
    )
}
