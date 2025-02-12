export type SuccessResponse<T> = {
    data: T
    success: true
    traceId: string
}

export type ErrorResponse = {
    messages: string
    success: false
    traceId: string
}
