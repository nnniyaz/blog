import { Lang } from '@domain/base/ml-string'

import classes from './home.module.scss'
import { translate } from '@lib/utils/translate'

interface HomeProps {
    lang: Lang
}

export default async function Home({ lang }: HomeProps) {
    return (
        <div className={classes.home}>
            <section className={classes.home__bio}>
                <img
                    className={classes.home__group__cover}
                    src="https://dynamic-media-cdn.tripadvisor.com/media/photo-o/13/f8/5c/05/picture-lake.jpg?w=900&h=500&s=1"
                    alt={translate(lang, 'author')}
                />
                <div className={classes.home__group__bio}></div>
            </section>
            <section>
                <h2>{translate(lang, 'featured_projects')}</h2>
                <ul>
                    <li>1</li>
                    <li>2</li>
                    <li>3</li>
                    <li>4</li>
                </ul>
            </section>
            <section>
                <h2>{translate(lang, 'articles')}</h2>
                <ul>
                    <li>1</li>
                    <li>2</li>
                    <li>3</li>
                    <li>4</li>
                </ul>
            </section>
            <section>
                <h2>{translate(lang, 'recent_posts')}</h2>
                <ul>
                    <li>1</li>
                    <li>2</li>
                    <li>3</li>
                    <li>4</li>
                </ul>
            </section>
        </div>
    )
}
