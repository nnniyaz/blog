import Layout from '@components/layout'
import { Lang } from '@domain/base/mlString'
import type { Metadata, Viewport } from 'next'

export const metadata: Metadata = {
    title: 'Сәлем! • Нияз Насыров',
    description: '',
    openGraph: {
        title: 'Сәлем! • Нияз Насыров',
        description: '',
        images: '',
        url: 'https://nassyrov.net',
        type: 'website',
    },
}

export const viewport: Viewport = {
    initialScale: 1,
    width: 'device-width',
}

export default function RootLayout({
    children,
}: Readonly<{ children: React.ReactNode }>) {
    return <Layout lang={Lang.KZ}>{children}</Layout>
}
