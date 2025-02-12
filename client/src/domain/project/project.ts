import { MlString } from '@domain/base/ml-string'

export type Project = {
    _id: string
    name: MlString
    description: MlString
    coverUri: string
    appLink: string
    sourceCodeLink: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}

export type ProjectsMany = {
    projects: Project[]
    count: number
}
