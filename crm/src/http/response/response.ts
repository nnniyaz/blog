export type SuccessResponse<T> = {
    success: true
    data: T
}

export type ErrorResponse = {
    success: false
    message: string
}
