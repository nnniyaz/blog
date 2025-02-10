import Layout from '@components/layout'
import { Lang } from '@domain/base/mlString'
import type { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'Hello! • Niyaz Nassyrov',
    description: '',
}

export default function RootLayout({
    children,
}: Readonly<{ children: React.ReactNode }>) {
    return (
        <Layout lang={Lang.EN}>
            {children}
        </Layout>
    )
}
