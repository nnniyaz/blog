import { AxiosResponse } from 'axios'

import { $api, ApiRoutes, Request } from '@http/index.ts'
import { ErrorResponse, SuccessResponse } from '@http/response/response.ts'

type AuthServiceLoginIn = {
    email: string
    password: string
}

export class AuthService {
    static async signIn(
        request: Request<AuthServiceLoginIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.AUTH_LOGIN,
            request.body,
        )
    }

    static async signOut(
        request: Request<null>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.AUTH_LOGOUT,
            null,
        )
    }
}
