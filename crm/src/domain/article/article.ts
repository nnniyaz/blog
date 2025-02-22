import { MlString } from '@domain/base/ml-string'

export type Article = {
    id: string
    title: MlString
    content: MlString
    isDeleted: boolean
    createdAt: string
    updatedAt: string
    version: number
}
