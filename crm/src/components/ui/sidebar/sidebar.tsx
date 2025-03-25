import classes from './sidebar.module.scss'
import { StyledLink } from '@components/ui/styled-link/styled-link.tsx'
import { PrivateRoutesList } from '@domain/base/routes-list.tsx'
import { charCountInString } from '@lib/utils/charCountInString.ts'
import { translate } from '@lib/utils/translate.ts'
import { Lang } from '@domain/base/ml-string'

interface SidebarProps {
    lang: Lang
}

export const Sidebar = ({ lang }: SidebarProps) => {
    return (
        <section className={classes.sidebar}>
            <h1 className={classes.sidebar__header}>
                {translate(lang, 'blog')}
            </h1>
            <nav className={classes.sidebar__body}>
                <ul className={classes.sidebar__body__list}>
                    {PrivateRoutesList.map((route) => {
                        if (charCountInString(route.path, '/') > 1) {
                            return null
                        }
                        return (
                            <li key={route.name}>
                                <StyledLink
                                    href={route.path}
                                    label={translate(lang, route.name)}
                                />
                            </li>
                        )
                    })}
                </ul>
            </nav>
        </section>
    )
}
