import { NewSuccess } from '@lib/response/success'
import { NextRequest } from 'next/server'

export async function GET(req: NextRequest) {
    const traceId = req.headers.get('x-trace-id')
    return NewSuccess(req, { message: traceId })
}
