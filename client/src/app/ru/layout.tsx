import Layout from '@components/layout'
import { Lang } from '@domain/base/mlString'
import localFont from '@next/font/local'
import type { Metadata } from 'next'

const sfProDisplay = localFont({
    src: [
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYULTRALIGHTITALIC.OTF',
            weight: '100',
            style: 'ultralight',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYTHINITALIC.OTF',
            weight: '200',
            style: 'thin',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYLIGHTITALIC.OTF',
            weight: '300',
            style: 'light',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYREGULAR.OTF',
            weight: '400',
            style: 'normal',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYMEDIUM.OTF',
            weight: '500',
            style: 'medium',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYSEMIBOLDITALIC.OTF',
            weight: '600',
            style: 'semibold',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYBOLD.OTF',
            weight: '700',
            style: 'bold',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYHEAVYITALIC.OTF',
            weight: '800',
            style: 'heavy',
        },
        {
            path: '../../../public/sf-pro-display/SFPRODISPLAYBLACKITALIC.OTF',
            weight: '900',
            style: 'black',
        },
    ],
})

export const metadata: Metadata = {
    title: 'Привет! • Нияз Насыров',
    description: '',
}

export default function RootLayout({
    children,
}: Readonly<{ children: React.ReactNode }>) {
    return (
        <Layout lang={Lang.RU} font={sfProDisplay.className}>
            {children}
        </Layout>
    )
}
