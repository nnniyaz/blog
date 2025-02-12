import { Lang } from '@domain/base/ml-string'

import classes from './home.module.scss'

interface HomeProps {
    lang: Lang
}

export default async function Home(props: HomeProps) {
    return <div className={classes.home}>{props.lang}</div>
}
