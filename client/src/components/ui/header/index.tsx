import StyledLink from '@components/ui/styled-link/styled-link'
import { Lang } from '@domain/base/ml-string'
import { translate } from '@lib/utils/translate'
import Link from 'next/link'

import classes from './header.module.scss'

interface HeaderProps {
    lang: Lang
    fullName: string
}

export default function Header({ lang, fullName }: HeaderProps) {
    return (
        <header className={classes.header}>
            <div className={classes.header__group}>
                <Link
                    href={`/${lang}`}
                    className={classes.header__group__author}
                >
                    {fullName}
                </Link>
                <nav>
                    <ul className={classes.header__group__nav_list}>
                        <li>
                            <StyledLink
                                href={`/${lang.toLowerCase()}/blog`}
                                label={translate(lang, 'blog')}
                            />
                        </li>
                        <li>
                            <StyledLink
                                href={`/${lang.toUpperCase()}/projects`}
                                label={translate(lang, 'projects')}
                            />
                        </li>
                    </ul>
                </nav>
            </div>
            <div className={classes.header__group}>
                <ul className={classes.header__group__lang__list}>
                    {Object.values(Lang).map((langItem) => {
                        if (langItem === 'key') return null
                        return (
                            <li key={langItem}>
                                <Link href={`/${langItem.toLowerCase()}`} className={classes.header__group__lang}>
                                    {langItem}
                                </Link>
                            </li>
                        )
                    })}
                </ul>
            </div>
        </header>
    )
}
