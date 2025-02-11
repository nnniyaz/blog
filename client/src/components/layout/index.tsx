import Footer from '@components/footer'
import Header from '@components/header'
import { Lang } from '@domain/base/mlString'

import './layout.scss'

interface LayoutProps {
    children: React.ReactNode
    lang: Lang
}

export default function Layout(props: LayoutProps) {
    return (
        <html lang={props.lang}>
            <body>
                <Header lang={props.lang} />
                <main>{props.children}</main>
                <Footer lang={props.lang} />
            </body>
        </html>
    )
}
