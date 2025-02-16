import { createSlice } from '@reduxjs/toolkit'

import { User } from '@domain/user/user.ts'

type State = {
    users: User[]
    isLoadingGetUsers: boolean
    selectedUser: User | null
    isLoadingGetSelectedUser: boolean
}

const initialState: State = {
    users: [],
    isLoadingGetUsers: false,
    selectedUser: null,
    isLoadingGetSelectedUser: false,
}

type SetUsersAction = {
    type: string
    payload: User[]
}

type SetIsLoadingGetUsersAction = {
    type: string
    payload: boolean
}

type SetSelectedUserAction = {
    type: string
    payload: User | null
}

type SetIsLoadingGetSelectedUserAction = {
    type: string
    payload: boolean
}

export const usersSlice = createSlice({
    name: 'user',
    initialState: initialState,
    reducers: {
        setUsers: (state, action: SetUsersAction) => {
            state.users = action.payload
        },
        setIsLoadingGetUsers: (state, action: SetIsLoadingGetUsersAction) => {
            state.isLoadingGetUsers = action.payload
        },
        setSelectedUser: (state, action: SetSelectedUserAction) => {
            state.selectedUser = action.payload
        },
        setIsLoadingGetSelectedUser: (
            state,
            action: SetIsLoadingGetSelectedUserAction,
        ) => {
            state.isLoadingGetSelectedUser = action.payload
        },
    },
})

export default usersSlice.reducer
