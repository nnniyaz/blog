import classes from './input.module.scss'
import * as React from 'react'
import { Input as AntdInput } from 'antd'

interface InputProps {
    placeholder?: string
    name?: string
    type?: string
    icon?: {
        element: React.ReactNode
        onClick?: () => void
    }
    value?: string
    onChange?: (value: string) => void
}

export const Input = (props: InputProps) => {
    return (
        <div className={classes.input__wrapper}>
            <AntdInput
                className={classes.input}
                type={props.type}
                name={props.name}
                placeholder={props.placeholder}
                value={props.value}
                onChange={(e) =>
                    props.onChange && props.onChange(e.target.value)
                }
            />
            {!!props.icon && (
                <div
                    className={classes.input__icon}
                    onClick={props.icon.onClick}
                >
                    {props.icon.element}
                </div>
            )}
        </div>
    )
}
