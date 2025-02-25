import classes from './breadcrumbs.module.scss'
import { StyledLink } from '@components/ui/styled-link/styled-link.tsx'
import { useLocation } from 'react-router-dom'
import { RoutesList } from '@domain/base/routes-list.tsx'
import { translate } from '@lib/utils/translate.ts'
import { Lang } from '@domain/base/ml-string.ts'
import { matchRoute } from '@lib/utils/match-route.ts'

interface BreadcrumbsProps {
    lang: Lang
}

export const Breadcrumbs = ({ lang }: BreadcrumbsProps) => {
    const pathname = useLocation().pathname.split('/').slice(1)
    return (
        <nav className={classes.crumbs}>
            <ol>
                {pathname.map((item, index) => {
                    const route = RoutesList.find((route) => {
                        return matchRoute(
                            route.path,
                            `/${pathname.slice(0, index + 1).join('/')}`,
                        )
                    })

                    if (!route) {
                        return null
                    }

                    if (index === pathname.length - 1) {
                        return (
                            <li className={classes.crumb} key={item}>
                                {translate(lang, route.name)}
                            </li>
                        )
                    }

                    return (
                        <li className={classes.crumb} key={item}>
                            <StyledLink
                                href={route.path}
                                label={translate(lang, route.name)}
                            />
                        </li>
                    )
                })}
            </ol>
        </nav>
    )
}
