import { ContentService } from '@/http/content-service'
import { Lang } from '@domain/base/ml-string'
import { Contact } from '@domain/contact/contact'
import { translate } from '@lib/utils/translate'

import Footer from '../ui/footer'
import Header from '../ui/header'
import './layout.scss'

interface LayoutProps {
    children: React.ReactNode
    lang: Lang
}

export default async function Layout({ children, lang }: LayoutProps) {
    const author = await ContentService.getAuthor()
    const contacts = await ContentService.getContacts()

    return (
        <html lang={lang}>
            <body>
                <div className={'layout'}>
                    <Header
                        lang={lang}
                        fullName={
                            author
                                ? author.firstName[lang] + author.lastName[lang]
                                : translate(lang, 'first_last_name')
                        }
                    />
                    <main className={'content'}>{children}</main>
                    <Footer
                        lang={lang}
                        contacts={
                            contacts
                                ? contacts
                                : ([
                                      {
                                          _id: '1',
                                          label: {
                                              KZ: 'GitHub',
                                              RU: 'GitHub',
                                              EN: 'GitHub',
                                          },
                                          link: '#',
                                          isDeleted: false,
                                          createdAt: '',
                                          updatedAt: '',
                                      },
                                      {
                                          _id: '2',
                                          label: {
                                              KZ: 'LinkedIn',
                                              RU: 'LinkedIn',
                                              EN: 'LinkedIn',
                                          },
                                          link: '#',
                                          isDeleted: false,
                                          createdAt: '',
                                          updatedAt: '',
                                      },
                                  ] as Contact[])
                        }
                    />
                </div>
            </body>
        </html>
    )
}
