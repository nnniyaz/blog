import { MlString } from '../base/mlString/mlString'

export type Contact = {
    id: string
    label: MlString
    link: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}
