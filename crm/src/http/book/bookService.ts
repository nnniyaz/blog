import { AxiosResponse } from 'axios'

import { MlString } from '@domain/base/ml-string'
import { Book } from '@domain/book/book.ts'
import { $api, ApiRoutes, Request } from '@http/index.ts'
import { ErrorResponse, SuccessResponse } from '@http/response/response.ts'

type BookServiceGetAllOut = {
    count: number
    books: Book[]
}

type BookServiceGetByIdIn = {
    id: string
}

type BookServiceCreateIn = {
    title: MlString
    description: MlString
    author: MlString
    coverUri: string
    eBookUri: string
}

type BookServiceUpdateIn = {
    id: string
    title: MlString
    description: MlString
    author: MlString
    coverUri: string
    eBookUri: string
}

type BookServiceDeleteIn = {
    id: string
}

type BookServiceRestoreIn = {
    id: string
}

export class BookService {
    static async getAll(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<BookServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<BookServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.BOOK_GET_ALL, null, { signal: request.controller?.signal })
    }

    static async getById(
        request: Request<BookServiceGetByIdIn>,
    ): Promise<AxiosResponse<SuccessResponse<Book> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<Book> | ErrorResponse>(
            ApiRoutes.BOOK_GET_BY_ID.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async create(
        request: Request<BookServiceCreateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BOOK_CREATE,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async update(
        request: Request<BookServiceUpdateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BOOK_UPDATE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async delete(
        request: Request<BookServiceDeleteIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BOOK_DELETE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async restore(
        request: Request<BookServiceRestoreIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BOOK_RESTORE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }
}
