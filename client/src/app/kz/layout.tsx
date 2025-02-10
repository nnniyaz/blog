import Layout from '@components/layout'
import { Lang } from '@domain/base/mlString'
import type { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'Сәлем! • Нияз Насыров',
    description: '',
}

export default function RootLayout({
    children,
}: Readonly<{ children: React.ReactNode }>) {
    return (
        <Layout lang={Lang.KZ}>
            {children}
        </Layout>
    )
}
