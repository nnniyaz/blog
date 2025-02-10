import { Lang } from '@domain/base/mlString'

import classes from './blog.module.scss'

interface BlogProps {
    lang: Lang
}

export default function Blog(props: BlogProps) {
    return <div className={classes.blog}>{props.lang}</div>
}
