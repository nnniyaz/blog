import { AxiosResponse } from 'axios'

import { Article } from '@domain/article/article'
import { MlString } from '@domain/base/mlString/mlString'
import { $api, ApiRoutes, Request } from '../index'
import { ErrorResponse, SuccessResponse } from '../response/response'

type ArticleServiceGetAllOut = {
    count: number
    articles: Article[]
}

type ArticleServiceGetByIdIn = {
    id: string
}

type ArticleServiceCreateIn = {
    title: MlString
    content: MlString
}

type ArticleServiceUpdateIn = {
    id: string
    title: MlString
    content: MlString
}

type ArticleServiceDeleteIn = {
    id: string
}

type ArticleServiceRestoreIn = {
    id: string
}

export class ArticleService {
    static async getAll(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<ArticleServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<ArticleServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.ARTICLE_GET_ALL, null, {
            signal: request.controller?.signal,
        })
    }

    static async getById(
        request: Request<ArticleServiceGetByIdIn>,
    ): Promise<AxiosResponse<SuccessResponse<Article> | ErrorResponse>> {
        return $api(request.lang).post<
            SuccessResponse<Article> | ErrorResponse
        >(ApiRoutes.ARTICLE_GET_BY_ID.replace(':id', request.body.id), null, {
            signal: request.controller?.signal,
        })
    }

    static async create(
        request: Request<ArticleServiceCreateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ARTICLE_CREATE,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async update(
        request: Request<ArticleServiceUpdateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ARTICLE_UPDATE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async delete(
        request: Request<ArticleServiceDeleteIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ARTICLE_DELETE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async restore(
        request: Request<ArticleServiceRestoreIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ARTICLE_RESTORE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }
}
