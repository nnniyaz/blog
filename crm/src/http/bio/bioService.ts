import { AxiosResponse } from 'axios'

import { MlString } from '@domain/base/mlString/mlString'
import { Bio } from '@domain/bio/bio'
import { $api, ApiRoutes, Request } from '../index'
import { ErrorResponse, SuccessResponse } from '../response/response'

type BioServiceGetAllOut = {
    count: number
    bios: Bio[]
}

type BioServiceGetByIdIn = {
    id: string
}

type BioServiceCreateIn = {
    bio: MlString
}

type BioServiceUpdateIn = {
    id: string
    bio: MlString
}

type BioServiceDeleteIn = {
    id: string
}

type BioServiceRestoreIn = {
    id: string
}

export class BioService {
    static async getAll(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<BioServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<BioServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.BIO_GET_ALL, null, { signal: request.controller?.signal })
    }

    static async getActive(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<BioServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<BioServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.BIO_GET_ACTIVE, null, {
            signal: request.controller?.signal,
        })
    }

    static async getById(
        request: Request<BioServiceGetByIdIn>,
    ): Promise<AxiosResponse<SuccessResponse<Bio> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<Bio> | ErrorResponse>(
            ApiRoutes.BIO_GET_BY_ID.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async create(
        request: Request<BioServiceCreateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BIO_CREATE,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async update(
        request: Request<BioServiceUpdateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BIO_UPDATE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async delete(
        request: Request<BioServiceDeleteIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BIO_DELETE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async restore(
        request: Request<BioServiceRestoreIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.BIO_RESTORE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }
}
