import { Lang } from '@domain/base/ml-string'

import classes from './blog.module.scss'

interface BlogProps {
    lang: Lang
}

export default function Blog(props: BlogProps) {
    return <div className={classes.blog}>{props.lang}</div>
}
