import { Outlet } from 'react-router-dom'
import classes from './layout.module.scss'
import { Sidebar } from '@components/ui/sidebar/sidebar.tsx'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { Breadcrumbs } from '@components/ui/breadcrumbs/breadcrumbs.tsx'

export const Layout = () => {
    const { lang } = useTypedSelector((state) => state.system)
    return (
        <main className={classes.layout}>
            <Sidebar lang={lang} />
            <section className={classes.content}>
                <Breadcrumbs lang={lang}/>
                <Outlet />
            </section>
        </main>
    )
}
