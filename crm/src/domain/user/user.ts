export type User = {
    id: string
    email: string
    role: UserRole
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}

export enum UserRole {
    ADMIN = 'admin',
    MODERATOR = 'moderator',
}
