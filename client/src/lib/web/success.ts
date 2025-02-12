import { HttpStatus } from '@domain/base/http-status-code'
import { response } from '@lib/web/response'
import { NextRequest } from 'next/server'

export function NewSuccess(req: NextRequest, body: unknown) {
    return response(
        req,
        HttpStatus.OK,
        JSON.stringify({
            data: body,
            success: true,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}
