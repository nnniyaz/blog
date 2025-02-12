export function currentTimeRFC3339() {
    function pad(n: number) {
        return n < 10 ? '0' + n : n
    }

    function timezoneOffset(offset: number) {
        if (offset === 0) {
            return 'Z'
        }
        const sign = offset > 0 ? '-' : '+'
        offset = Math.abs(offset)
        return sign + pad(Math.floor(offset / 60)) + ':' + pad(offset % 60)
    }

    const d = new Date()
    return (
        d.getFullYear() +
        '-' +
        pad(d.getMonth() + 1) +
        '-' +
        pad(d.getDate()) +
        'T' +
        pad(d.getHours()) +
        ':' +
        pad(d.getMinutes()) +
        ':' +
        pad(d.getSeconds()) +
        timezoneOffset(d.getTimezoneOffset())
    )
}
