import { Lang } from '@domain/base/mlString'

import classes from './project.module.scss'

interface ProjectProps {
    lang: Lang
}

export default function Project(props: ProjectProps) {
    return <div className={classes.project}>{props.lang}</div>
}
