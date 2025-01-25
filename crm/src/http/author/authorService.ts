import {AxiosResponse} from 'axios'

import {Author} from '../../domain/author/author'
import {MlString} from '../../domain/base/mlString/mlString'
import {$api, ApiRoutes, Request} from '../index'
import {ErrorResponse, SuccessResponse} from '../response/response'

type AuthorServiceGetAllOut = {
    count: number
    authors: Author[]
}

type AuthorServiceGetByIdIn = {
    id: string
}

type AuthorServiceCreateIn = {
    firstName: MlString
    lastName: MlString
    avatarUri: string
}

type AuthorServiceUpdateIn = {
    id: string
    firstName: MlString
    lastName: MlString
    avatarUri: string
}

type AuthorServiceDeleteIn = {
    id: string
}

type AuthorServiceRestoreIn = {
    id: string
}

export class AuthorService {
    static async getAll(request: Request<null>): Promise<AxiosResponse<SuccessResponse<AuthorServiceGetAllOut> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<AuthorServiceGetAllOut> | ErrorResponse>(ApiRoutes.AUTHOR_GET_ALL, null, {signal: request.controller?.signal})
    }

    static async getById(request: Request<AuthorServiceGetByIdIn>): Promise<AxiosResponse<SuccessResponse<Author> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<Author> | ErrorResponse>(ApiRoutes.AUTHOR_GET_BY_ID.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async create(request: Request<AuthorServiceCreateIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.AUTHOR_CREATE, null, {signal: request.controller?.signal})
    }

    static async update(request: Request<AuthorServiceUpdateIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.AUTHOR_UPDATE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async delete(request: Request<AuthorServiceDeleteIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.AUTHOR_DELETE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async restore(request: Request<AuthorServiceRestoreIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.AUTHOR_RESTORE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }
}