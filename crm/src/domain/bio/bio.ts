import { MlString } from '@domain/base/ml-string'

export type Bio = {
    id: string
    bio: MlString
    active: boolean
    createdAt: string
    updatedAt: string
    version: number
}
