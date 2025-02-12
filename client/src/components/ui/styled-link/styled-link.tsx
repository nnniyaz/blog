import Link from 'next/link'

import classes from './styled-link.module.scss'

interface StyledLinkProps {
    href: string
    label: string
}

export default function StyledLink(props: StyledLinkProps) {
    return (
        <Link href={props.href} className={classes.link}>
            {props.label}
        </Link>
    )
}
