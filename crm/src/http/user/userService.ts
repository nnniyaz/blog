import { AxiosResponse } from 'axios'

import { User, UserRole } from '@domain/user/user.ts'
import { $api, ApiRoutes, Request } from '@http/index.ts'
import { ErrorResponse, SuccessResponse } from '@http/response/response.ts'

type UserServiceGetAllOut = {
    count: number
    users: User[]
}

type UserServiceGetByIdIn = {
    id: string
}

type UserServiceCreateIn = {
    email: string
    password: string
    role: UserRole
}

type UserServiceUpdateEmailIn = {
    id: string
    email: string
}

type UserServiceUpdatePasswordIn = {
    id: string
    password: string
}

type UserServiceUpdateRoleIn = {
    id: string
    role: UserRole
}

type UserServiceDeleteIn = {
    id: string
}

type UserServiceRestoreIn = {
    id: string
}

export class UserService {
    static async getAll(
        request: Request<null>,
    ): Promise<
        AxiosResponse<SuccessResponse<UserServiceGetAllOut> | ErrorResponse>
    > {
        return $api(request.lang).post<
            SuccessResponse<UserServiceGetAllOut> | ErrorResponse
        >(ApiRoutes.USER_GET_ALL, null, { signal: request.controller?.signal })
    }

    static async getById(
        request: Request<UserServiceGetByIdIn>,
    ): Promise<AxiosResponse<SuccessResponse<User> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<User> | ErrorResponse>(
            ApiRoutes.USER_GET_BY_ID.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async create(
        request: Request<UserServiceCreateIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_CREATE,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async updateEmail(
        request: Request<UserServiceUpdateEmailIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_UPDATE_EMAIL.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async updatePassword(
        request: Request<UserServiceUpdatePasswordIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_UPDATE_PASSWORD.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async updateRole(
        request: Request<UserServiceUpdateRoleIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_UPDATE_ROLE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async delete(
        request: Request<UserServiceDeleteIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_DELETE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }

    static async restore(
        request: Request<UserServiceRestoreIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.USER_RESTORE.replace(':id', request.body.id),
            null,
            { signal: request.controller?.signal },
        )
    }
}
