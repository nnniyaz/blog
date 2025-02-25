import { Lang } from '@domain/base/ml-string.ts'
import { useLocation } from 'react-router-dom'
import { RoutesList } from '@domain/base/routes-list.tsx'
import { matchRoute } from '@lib/utils/match-route.ts'
import { translate } from '@lib/utils/translate.ts'

interface PageHeaderProps {
    lang: Lang
}

export const PageHeader = ({ lang }: PageHeaderProps) => {
    const pathname = useLocation().pathname
    const route = RoutesList.find((route) => {
        return matchRoute(route.path, pathname)
    })
    if (!route) {
        return 'Page header not identified'
    }
    return <h2>{translate(lang, route.name)}</h2>
}
