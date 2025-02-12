export function parseIntQueryParam(query: string | null): number {
    return isNaN(parseInt(query || '')) ? 0 : parseInt(query || '')
}

export function parseBooleanQueryParam(query: string | null): boolean {
    return query === 'true'
}
