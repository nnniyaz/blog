import { Lang } from '@domain/base/ml-string'

interface FooterProps {
    lang: Lang
}

export default function Footer(props: FooterProps) {
    return (
        <footer>
            <p>© 2021</p>
        </footer>
    )
}
