export function getLocale(acceptLanguage: string | null) {
    if (!acceptLanguage) return 'en'
    const locale = acceptLanguage?.split(',') || ['en-US']
    switch (locale[0]) {
        case 'ru-RU':
        case 'ru':
            return 'ru'
        case 'kz-KZ':
        case 'kz':
            return 'kz'
        default:
            return 'en'
    }
}
