import { Link } from 'react-router-dom'

import classes from './styled-link.module.scss'

interface StyledLinkProps {
    href: string
    label: string
}

export function StyledLink(props: StyledLinkProps) {
    return (
        <Link to={props.href} className={classes.link}>
            {props.label}
        </Link>
    )
}
