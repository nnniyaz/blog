import { MlString } from '@domain/base/ml-string'

export type Project = {
    id: string
    name: MlString
    description: MlString
    coverUri: string
    appLink: string
    sourceCodeLink: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}
