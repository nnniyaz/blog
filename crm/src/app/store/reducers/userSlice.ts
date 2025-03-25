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

type SetSelectedUserAction = {
    type: string
    payload: User | null
}

export const userSlice = createSlice({
    name: 'user',
    initialState: initialState,
    reducers: {
        setUsers: (state, action: SetUsersAction) => {
            state.users = action.payload
        },
        setSelectedUser: (state, action: SetSelectedUserAction) => {
            state.selectedUser = action.payload
        },
    },
})

export default userSlice.reducer
