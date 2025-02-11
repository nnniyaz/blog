import { Lang } from '@domain/base/mlString'

interface HeaderProps {
    lang: Lang
}

export default function Header(props: HeaderProps) {
    return (
        <header>
            <p>© 2021</p>
        </header>
    )
}
