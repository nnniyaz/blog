import { MlString } from '@domain/base/mlString/mlString.ts'

export type Contact = {
    id: string
    label: MlString
    link: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}
