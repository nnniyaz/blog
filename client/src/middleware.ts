import { getLocale } from '@lib/utils/locale'
import { NextRequest, NextResponse } from 'next/server'

const locales = ['en', 'ru', 'kz']

export default function middleware(request: NextRequest) {
    const { pathname } = request.nextUrl
    const pathnameHasLocale = locales.some(
        (locale) =>
            pathname.startsWith(`/${locale}/`) || pathname === `/${locale}`,
    )

    if (pathnameHasLocale) {
        const response = NextResponse.next()
        response.headers.set('x-current-path', pathname)
        return response
    }

    const locale = getLocale(request.headers.get('Accept-Language'))
    request.nextUrl.pathname = `/${locale}${pathname}`
    return NextResponse.redirect(request.nextUrl)
}

export const config = {
    matcher: [
        /*
         * Match all request paths except for the ones starting with:
         * - api (API routes)
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico (favicon file)
         */
        '/((?!api|_next/static|_next/image|favicon.ico).*)',
    ],
}
