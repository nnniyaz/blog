export function matchRoute(pattern: string, url: string): boolean {
    const regex = new RegExp('^' + pattern.replace(':id', '([a-zA-Z0-9-]+)') + '$')
    return regex.test(url)
}
