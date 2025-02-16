import { AxiosResponse } from 'axios'

import { MlString } from '@domain/base/mlString/mlString.ts'
import { Project } from '@domain/project/project.ts'
import { $api, ApiRoutes, Request } from '@http/index.ts'
import { ErrorResponse, SuccessResponse } from '@http/response/response.ts'

type ProjectServiceGetAllOut = {
    count: number
    projects: Project[]
}

type ProjectServiceGetByIdIn = {
    id: string
}

type ProjectServiceCreateIn = {
    name: MlString
    description: MlString
    coverUri: string
    appLink: string
    sourceCodeLink: string
}

type ProjectServiceUpdateIn = {
    id: string
    name: MlString
    description: MlString
    coverUri: string
    appLink: string
    sourceCodeLink: string
}

type ProjectServiceDeleteIn = {
    id: string
}

type ProjectServiceRestoreIn = {
    id: string
}

export class ProjectService {
    static async getAll(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<ProjectServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<ProjectServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.PROJECT_GET_ALL, null, {
            signal: request.controller?.signal,
        })
    }

    static async getById(
        request: Request<ProjectServiceGetByIdIn>,
    ): Promise<AxiosResponse<SuccessResponse<Project> | ErrorResponse>> {
        return $api(request.lang).post<
            SuccessResponse<Project> | ErrorResponse
        >(ApiRoutes.PROJECT_GET_BY_ID.replace(':id', request.body.id), null, {
            signal: request.controller?.signal,
        })
    }

    static async create(
        request: Request<ProjectServiceCreateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.PROJECT_CREATE,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async update(
        request: Request<ProjectServiceUpdateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.PROJECT_UPDATE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async delete(
        request: Request<ProjectServiceDeleteIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.PROJECT_DELETE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async restore(
        request: Request<ProjectServiceRestoreIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.PROJECT_RESTORE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }
}
