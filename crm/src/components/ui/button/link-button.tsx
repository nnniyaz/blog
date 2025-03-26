import classes from './button.module.scss'
import { Link } from 'react-router-dom'
import * as React from 'react'

interface LinkButtonProps {
    children: React.ReactNode
    href: string
}

export const LinkButton = (props: LinkButtonProps) => {
    return (
        <Link
            to={props.href}
            className={classes.button__rectangular}
            style={{ textDecoration: 'none', fontWeight: '400' }}
        >
            {props.children}
        </Link>
    )
}
