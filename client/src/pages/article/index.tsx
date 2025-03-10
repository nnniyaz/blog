import { Lang } from '@domain/base/ml-string'

import classes from './article.module.scss'

interface ArticleProps {
    lang: Lang
}

export default function Article(props: ArticleProps) {
    return <div className={classes.article}>{props.lang}</div>
}
