import { MlString } from '@domain/base/mlString/mlString.ts'

export type Author = {
    id: string
    firstName: MlString
    lastName: MlString
    avatarUri: string
    createdAt: string
    updatedAt: string
    version: number
}
