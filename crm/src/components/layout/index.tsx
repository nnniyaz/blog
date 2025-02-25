import { Outlet } from 'react-router-dom'
import classes from './layout.module.scss'
import { Sidebar } from '@components/ui/sidebar/sidebar.tsx'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { Breadcrumbs } from '@components/ui/breadcrumbs/breadcrumbs.tsx'
import { PageHeader } from '@components/ui/page-header/page-header.tsx'

export const Layout = () => {
    const { lang } = useTypedSelector((state) => state.system)
    return (
        <main className={classes.layout}>
            <Sidebar lang={lang} />
            <section className={classes.content}>
                <Breadcrumbs lang={lang} />
                <PageHeader lang={lang} />
                <Outlet />
            </section>
        </main>
    )
}
