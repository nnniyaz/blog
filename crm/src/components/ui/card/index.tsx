import classes from './card.module.scss'

interface CardProps {
    children: React.ReactNode
    title?: string
    subtitle?: string
}

export const Card = (props: CardProps) => {
    return (
        <div className={classes.card}>
            {props.title && <h2>{props.title}</h2>}
            {props.subtitle && (
                <p className={classes.card__subtitle}>{props.subtitle}</p>
            )}
            <div className={classes.card__content}>{props.children}</div>
        </div>
    )
}
