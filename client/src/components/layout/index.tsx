import { Lang } from '@domain/base/mlString'

import './layout.scss'

interface LayoutProps {
    children: React.ReactNode
    lang: Lang
}

export default function Layout(props: LayoutProps) {
    return (
        <html lang={props.lang}>
            <body>{props.children}</body>
        </html>
    )
}
