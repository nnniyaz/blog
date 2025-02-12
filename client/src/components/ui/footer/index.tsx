import { Lang } from '@domain/base/ml-string'
import { Contact } from '@domain/contact/contact'
import { translate } from '@lib/utils/translate'
import Link from 'next/link'

import classes from './footer.module.scss'
import StyledLink from '@components/ui/styled-link/styled-link'

interface FooterProps {
    lang: Lang
    contacts: Contact[]
}

export default function Footer(props: FooterProps) {
    return (
        <footer className={classes.footer}>
            <div>
                <h2 className={'visually_hidden'}>
                    {translate(props.lang, 'contacts')}
                </h2>
                <ul className={classes.footer__contacts}>
                    {props.contacts.map((contact, index) => (
                        <li key={index}>
                            <StyledLink
                                href={contact.link}
                                label={translate(props.lang, contact.label)}
                            />
                        </li>
                    ))}
                </ul>
            </div>
            <div>
                <Link
                    className={classes.footer__copyright}
                    href={'#'}
                >
                    {`${'© 2019+'}`}
                </Link>
            </div>
        </footer>
    )
}
