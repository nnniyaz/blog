import { MlString } from '@domain/base/ml-string'

export type Article = {
    _id: string
    title: MlString
    content: MlString
    isDeleted: boolean
    createdAt: string
    updatedAt: string
    version: number
}

export type ArticlesMany = {
    articles: Article[]
    count: number
}
