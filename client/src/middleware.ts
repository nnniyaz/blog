import { getLocale } from '@lib/utils/locale'
import { NextRequest, NextResponse } from 'next/server'
import { v4 as uuidv4 } from 'uuid'

const locales = ['en', 'ru', 'kz']

export default function middleware(req: NextRequest) {
    const { pathname } = req.nextUrl
    let res: NextResponse = NextResponse.next()

    res = addTrackerHeaders(res)

    if (pathname.startsWith('/api')) {
        return res
    }

    return redirectToProperLocale(req, res, pathname)
}

function addTrackerHeaders(res: NextResponse): NextResponse {
    res.headers.set('x-trace-id', uuidv4())
    res.headers.set('x-start-time', Date.now().toString())
    return res
}

function redirectToProperLocale(
    req: NextRequest,
    res: NextResponse,
    pathname: string,
) {
    const pathnameHasLocale = locales.some(
        (locale) =>
            pathname.startsWith(`/${locale}/`) || pathname === `/${locale}`,
    )

    if (pathnameHasLocale) {
        res.headers.set('x-current-path', pathname)
        return res
    }

    const locale = getLocale(req.headers.get('Accept-Language'))
    req.nextUrl.pathname = `/${locale}${pathname}`

    return NextResponse.redirect(req.nextUrl)
}

export const config = {
    matcher: [
        /*
         * Match all req paths except for the ones starting with:
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico (favicon file)
         */
        '/((?!_next/static|_next/image|favicon.ico|robots.txt|sitemap.xml).*)',
    ],
}
