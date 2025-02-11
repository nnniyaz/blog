import { HttpStatus } from '@domain/base/http-status-code'
import { getLogger } from '@lib/logger/logger'
import { NextRequest } from 'next/server'

export function response(req: NextRequest, status: HttpStatus, body: string) {
    const res = new Response(body, {
        status: status,
        headers: {
            'Content-Type': 'application/json',
        },
    })
    if (status >= 400) {
        getLogger().error(req, res)
    } else if (status >= 300) {
        getLogger().warn(req, res)
    } else {
        getLogger().info(req, res)
    }
    return res
}
