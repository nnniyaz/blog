import { Lang } from '@domain/base/mlString'

import './layout.scss'

interface LayoutProps {
    children: React.ReactNode
    font: string
    lang: Lang
}

export default function Layout(props: LayoutProps) {
    return (
        <html lang={props.lang} className={props.font}>
            <body>{props.children}</body>
        </html>
    )
}
