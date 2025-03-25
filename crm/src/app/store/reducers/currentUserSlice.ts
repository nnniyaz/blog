import { User } from '@domain/user/user.ts'
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { RootState } from '@app/store'
import { CurrentUserService } from '@http/currentUser/currentUserService.ts'
import { authSlice } from '@app/store/reducers/authSlice.ts'

type State = {
    isAuth: boolean
    currentUser: User | null
    isLoadingGetCurrentUser: boolean
    error: string | null
}

const initialState: State = {
    isAuth: false,
    currentUser: null,
    isLoadingGetCurrentUser: false,
    error: null,
}

type SetIsAuthAction = {
    type: string
    payload: boolean
}

type SetCurrentUserAction = {
    type: string
    payload: User | null
}

export const getCurrentUser = createAsyncThunk(
    'currentUser/getCurrentUser',
    async (_, { getState, dispatch }) => {
        const state = getState() as RootState
        try {
            const res = await CurrentUserService.getCurrentUser({
                lang: state.system.lang,
                body: null,
            })
            if (res.data.success) {
                dispatch(currentUserSlice.actions.setCurrentUser(res.data.data))
                dispatch(authSlice.actions.setIsAuth(true))
            }
        } catch (error) {
            throw error
        }
    },
)

export const currentUserSlice = createSlice({
    name: 'currentUser',
    initialState: initialState,
    reducers: {
        setIsAuth: (state, action: SetIsAuthAction) => {
            state.isAuth = action.payload
        },
        setCurrentUser: (state, action: SetCurrentUserAction) => {
            state.currentUser = action.payload
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getCurrentUser.pending, (state) => {
            state.isLoadingGetCurrentUser = true
            state.error = null
        })
        builder.addCase(getCurrentUser.fulfilled, (state) => {
            state.isLoadingGetCurrentUser = false
        })
        builder.addCase(getCurrentUser.rejected, (state, action) => {
            state.isLoadingGetCurrentUser = false
            state.error = action.payload as string
        })
    },
})

export default currentUserSlice.reducer
