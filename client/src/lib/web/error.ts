import { HttpStatus } from '@domain/base/http-status-code'
import { response } from '@lib/web/response'
import { NextRequest } from 'next/server'

export function NewBad(req: NextRequest, error: Error) {
    return response(
        req,
        HttpStatus.BAD_REQUEST,
        JSON.stringify({
            messages: [error.message],
            success: false,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}

export function NewForbidden(req: NextRequest, error: Error) {
    return response(
        req,
        HttpStatus.FORBIDDEN,
        JSON.stringify({
            messages: [error.message],
            success: false,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}

export function NewNotFound(req: NextRequest, error: Error) {
    return response(
        req,
        HttpStatus.NOT_FOUND,
        JSON.stringify({
            messages: [error.message],
            success: false,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}

export function NewInternal(req: NextRequest, error: Error) {
    return response(
        req,
        HttpStatus.INTERNAL_SERVER_ERROR,
        JSON.stringify({
            messages: [error.message],
            success: false,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}

export function NewError(req: NextRequest, error: Error) {
    return response(
        req,
        HttpStatus.INTERNAL_SERVER_ERROR,
        JSON.stringify({
            messages: [error.message],
            success: false,
            traceId: req.headers.get('x-trace-id'),
        }),
    )
}
