import { MlString } from '@domain/base/mlString/mlString.ts'

export type Bio = {
    id: string
    bio: MlString
    active: boolean
    createdAt: string
    updatedAt: string
    version: number
}
