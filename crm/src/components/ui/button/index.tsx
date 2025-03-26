import classes from './button.module.scss'
import * as React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

interface ButtonProps {
    type?: 'submit' | 'button' | 'reset'
    children: React.ReactNode
    transparent?: boolean
    loading?: boolean
    onClick?: (args: unknown) => void
    rounded?: boolean
}

export const Button = (props: ButtonProps) => {
    if (props.rounded) {
        return (
            <button
                type={props.type}
                onClick={props.onClick}
                className={
                    props.transparent
                        ? classes.button__rounded__transparent
                        : classes.button__rounded
                }
            >
                {props.children}
                {props.loading && (
                    <div className={classes.button__loading__wrapper}>
                        <LoadingOutlined className={classes.button__loading} />
                    </div>
                )}
            </button>
        )
    }

    return (
        <button
            type={props.type}
            onClick={props.onClick}
            className={classes.button__rectangular}
        >
            {props.children}
            {props.loading && (
                <div className={classes.button__loading__wrapper}>
                    <LoadingOutlined className={classes.button__loading} />
                </div>
            )}
        </button>
    )
}
