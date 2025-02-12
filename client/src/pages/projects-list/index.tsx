import { Lang } from '@domain/base/ml-string'

import classes from './projects-list.module.scss'

interface ProjectsListProps {
    lang: Lang
}

export default function ProjectsList(props: ProjectsListProps) {
    return <div className={classes.projects_list}>{props.lang}</div>
}
