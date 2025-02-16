import { MlString } from '@domain/base/mlString/mlString.ts'

export type Article = {
    id: string
    title: MlString
    content: MlString
    isDeleted: boolean
    createdAt: string
    updatedAt: string
    version: number
}
